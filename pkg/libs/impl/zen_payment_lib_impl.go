package impl

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/libs"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/utils/logger"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/parnurzeal/gorequest"
	"golang.org/x/exp/slices"
)

type zenPaymentLibImpl struct {
}

var _ libs.PaymentLib = (*zenPaymentLibImpl)(nil)

func NewZenPaymentLib() *zenPaymentLibImpl {
	return &zenPaymentLibImpl{}
}

func (l *zenPaymentLibImpl) Create(ctx context.Context, r libs.PaymentCreateRequest) (libs.ZenPaymentCreateResponse, error) {
	zenPaymentConfig := paymentMethods[paymentModeToZenType[r.Payment.Mode]]
	if err := l.validatePayment(ctx, r, zenPaymentConfig); err != nil {
		return libs.ZenPaymentCreateResponse{}, err
	}
	zenPaymentRequest := l.zenGeneralPaymentRequest(ctx, r, zenPaymentConfig)
	if err := l.setPaymentSpecificDetails(ctx, r, zenPaymentConfig, &zenPaymentRequest); err != nil {
		return libs.ZenPaymentCreateResponse{}, err
	}
	l.setPaymentAuthorizationForDifferentCurrency(ctx, r.Order, &zenPaymentRequest)
	res, err := l.createTransactionAtZen(ctx, zenPaymentRequest)
	if err != nil {
		return libs.ZenPaymentCreateResponse{}, err
	}
	return res, nil
}

func (l *zenPaymentLibImpl) validatePayment(ctx context.Context, r libs.PaymentCreateRequest, c PaymentMethodConfigs) error {

	if !slices.Contains(c.Currencies, r.Order.Currency) {
		return ErrCurrencyNotSupported
	}
	floatAmount, err := strconv.ParseFloat(r.Order.Amount, 64)
	if err != nil {
		return err
	}
	if floatAmount > c.MaxTxn || floatAmount < c.MinTxn {
		return err
	}
	return nil
}

func (l *zenPaymentLibImpl) zenGeneralPaymentRequest(ctx context.Context, r libs.PaymentCreateRequest, c PaymentMethodConfigs) libs.ZenPaymentCreateRequest {
	return libs.ZenPaymentCreateRequest{
		MerchantTransactionID: r.Payment.GeneratedID,
		PaymentChannel:        c.ChannelCode,
		Amount:                r.Order.Amount,
		Currency:              r.Order.Currency,
		CustomIpnURL:          app.Config().ZenIpnUrl,
		Comment:               "Purchase",
		FraudFields:           libs.ZenFingerPrint{FingerPrintID: r.Signature},
		Items: []libs.ZenItem{{
			Code:            "O1",
			Category:        "Customer Purchase",
			Name:            "Customer Item",
			Price:           r.Order.Amount,
			Quantity:        1,
			LineAmountTotal: r.Order.Amount,
		}},
		Customer: libs.ZenCustomer{
			FirstName: r.Account.FirstName.String,
			LastName:  r.Account.LastName.String,
			Email:     r.Account.Email,
			Phone:     r.Account.MobileNumber.String,
			IP:        r.Payment.IPAddress,
		},
	}
}

func (l *zenPaymentLibImpl) setPaymentSpecificDetails(ctx context.Context, r libs.PaymentCreateRequest, c PaymentMethodConfigs, zenPaymentRequest *libs.ZenPaymentCreateRequest) error {
	switch c.GatewayType {
	case "onetime":
		cardData := r.Payment.Data.(map[string]any)
		zenPaymentRequest.PaymentSpecificData = libs.OneTimePaymentSpecificData{
			Type:       "onetime",
			Descriptor: "card pay",
			Card: libs.OnetimeCardDetails{
				Number:     cardData["number"].(string),
				ExpiryDate: cardData["expiry_date"].(string),
				Cvv:        cardData["cvv"].(string),
			},
			Skip3Ds:         true,
			BrowserDetails:  r.BrowserDetails,
			ReturnVerifyURL: "",
		}
	case "general":
		zenPaymentRequest.PaymentSpecificData = libs.GeneralPaymentSpecificData{
			Type:      "general",
			ReturnUrl: "",
		}
	case "external_payment_token":
		tokenData := r.Payment.Data.(map[string]any)
		zenPaymentRequest.PaymentSpecificData = libs.TokenPaymentSpecificData{
			Type:            "external_payment_token",
			Descriptor:      "token pay",
			BrowserDetails:  r.BrowserDetails,
			Token:           tokenData["token"].(string),
			ReturnVerifyURL: "",
		}
	case "trustly":
		zenPaymentRequest.PaymentSpecificData = libs.GeneralPaymentSpecificData{
			Type:      "general",
			ReturnUrl: "",
		}
	default:
		return ErrInvalidGateway
	}
	return nil
}

func (l *zenPaymentLibImpl) setPaymentAuthorizationForDifferentCurrency(ctx context.Context, o models.Order, zenPaymentRequest *libs.ZenPaymentCreateRequest) {
	if o.Currency != "EUR" {
		zenPaymentRequest.Authorization = &libs.ZenSeparateCurrencyAuthorization{
			Amount:   o.Amount,
			Currency: o.Currency,
		}
	}
}

func (l *zenPaymentLibImpl) createTransactionAtZen(ctx context.Context, request libs.ZenPaymentCreateRequest) (libs.ZenPaymentCreateResponse, error) {
	logger.Info(ctx, "the request is ", "request", request)
	zenResponse := libs.ZenPaymentCreateResponse{}
	response, body, errs := gorequest.New().Post(l.transactionApi()).
		AppendHeader("Authorization", "Bearer e00f47713eb44a4db21ff837d5ae79e4").SendStruct(&request).
		Retry(zenPaymentApiRetryCount, zenPaymentApiRetryWaitTime, zenPaymentApiRetryStatus...).EndBytes()
	if len(errs) > 0 {
		return zenResponse, errs[0]
	}
	logger.Info(ctx, "the response is ", "response", string(body))
	if response.StatusCode != http.StatusCreated {
		logger.Error(ctx, "the status code for zen payments is ", "code", response.StatusCode)
		return zenResponse, ErrPaymentCreateAtZenFailed
	}
	if err := json.Unmarshal(body, &zenResponse); err != nil {
		logger.Error(ctx, "json unmarshal for zen response failed", "error", err)
		return zenResponse, err
	}
	logger.Info(ctx, "the response form zen is ", "zen", zenResponse)
	return zenResponse, nil
}

func (l *zenPaymentLibImpl) Callback(ctx context.Context, request libs.CallbackRequest) error {
	hashString := fmt.Sprintf(
		"%s%s%s%s%s",
		request.MerchantTransactionID,
		request.Currency,
		request.Amount,
		request.Status,
		app.Config().ZenMerchantSecret,
	)
	logger.Info(ctx, "the hash string is", " hashString", hashString)
	if !l.checkHash(ctx, hashString, request.Hash) {
		return errors.New("invalid hash not matching")
	}

	return nil
}

func (l *zenPaymentLibImpl) checkHash(ctx context.Context, givenString, hashString string) bool {
	h := sha256.New()
	h.Write([]byte(givenString))
	generatedHash := strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
	logger.Debug(ctx, "the generated hash is", " generatedHash", generatedHash, "hashString", hashString)
	return strings.Compare(generatedHash, hashString) == 0
}

func (l *zenPaymentLibImpl) host() string {
	return os.Getenv("ZEN_HOST")
}

func (l *zenPaymentLibImpl) transactionApi() string {
	return l.host() + os.Getenv("ZEN_TRANSACTIONS_API")
}
