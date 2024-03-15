package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type orderServiceImpl struct {
	orderRepo         repositories.OrderRepository
	retailerRepo      repositories.RetailerRepository
	apiConfigService  services.ApiConfigService
	accountService    services.AccountService
	addressService    services.AddressService
	txnFeesConfigRepo repositories.TxnFeeConfigRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) *orderServiceImpl {
	return &orderServiceImpl{orderRepo: orderRepo}
}

func DefaultOrderServiceImpl() *orderServiceImpl {
	return &orderServiceImpl{
		postgresql.DefaultOrderRepositoryImpl(),
		postgresql.DefaultRetailerRepositoryImpl(),
		DefaultApiConfigServiceImpl(),
		DefaultAccountServiceImpl(),
		DefaultAddressServiceImpl(),
		postgresql.DefaultTxnFeeConfigRepositoryImpl(),
	}
}

func (s *orderServiceImpl) CreateOrder(ctx context.Context, a *models.Order) error {
	return s.orderRepo.Create(ctx, a)
}

func (s *orderServiceImpl) CreateBuyOrder(ctx context.Context, req services.OrderCreateRequest) (models.Order, error) {
	ok, retailer, err := s.apiConfigService.CheckOrderCreateChecksum(ctx, req.ApiKey, req)
	if err != nil {
		return models.Order{}, err
	}
	if !ok {
		return models.Order{}, errors.New("checksum mismatch")
	}
	orderDetail := req.OrderDetail
	// check if external order id already exists
	order, err := s.orderRepo.FindByExternalOrderID(ctx, orderDetail.MerchantTransactionID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(ctx, "unexpected error", "error", err)
		return models.Order{}, err
	}
	if order.ID != 0 {
		return order, errors.New("order already exists")
	}

	logger.Info(ctx, "the order details are", "orderDetail", orderDetail)
	// find account for payment
	account, err := s.accountService.GetAccountByEmail(ctx, orderDetail.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(ctx, "unexpected error", "error", err)
		return models.Order{}, err
	}
	logger.Info(ctx, "the account is", "account", account)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Info(ctx, "creating account", "email", orderDetail.Email)
		// create account
		account = models.Account{
			Type:         models.AccountTypeCustomer,
			Status:       models.AccountStatusActive,
			FirstName:    null.StringFrom(orderDetail.FirstName),
			LastName:     null.StringFrom(orderDetail.LastName),
			Email:        orderDetail.Email,
			CountryCode:  null.StringFrom(orderDetail.CountryCode),
			MobileNumber: null.StringFrom(orderDetail.MobileNumber),
			Country:      null.StringFrom(orderDetail.Country),
		}
		if err := s.accountService.Create(ctx, &account); err != nil {
			logger.Error(ctx, "unexpected error", "error", err)
			return models.Order{}, err
		}
		logger.Info(ctx, "created account", "account_id", account.ID)
		address := models.Address{
			Type:         models.AddressTypeResidential,
			Status:       models.AddressStatusActive,
			AddressLine1: null.StringFrom(orderDetail.AddressLine1),
			AddressLine2: null.StringFrom(orderDetail.AddressLine2),
			City:         null.StringFrom(orderDetail.City),
			State:        null.StringFrom(orderDetail.State),
			Country:      null.StringFrom(orderDetail.Country),
			AccountID:    account.ID,
		}
		if err := s.addressService.Create(ctx, &address); err != nil {
			logger.Error(ctx, "unexpected error", "error", err)
			return models.Order{}, err
		}
		logger.Info(ctx, "created address", "address_id", address.ID)
	}
	amount, err := strconv.ParseFloat(orderDetail.Amount, 64)
	if err != nil {
		logger.Error(ctx, "unexpected error", "error", err)
		return models.Order{}, err
	}
	exchangeRate, err := s.baseCurrencyExchangeRate(ctx, orderDetail.Currency)
	if err != nil {
		logger.Error(ctx, "unexpected error", "error", err)
		return models.Order{}, err
	}
	cryptoExchangeRate, err := s.cryptoExchangeRate(ctx)
	if err != nil {
		logger.Error(ctx, "unexpected error", "error", err)
		return models.Order{}, err
	}
	baseCurrencyAmount := convertTo2Decimal(amount * exchangeRate)
	cryptoCurrencyAmount := convertTo2Decimal(baseCurrencyAmount * cryptoExchangeRate)
	logger.Info(ctx, "the exchange rate is",
		"exchange_rate", exchangeRate,
		"base_currency_amount", baseCurrencyAmount,
		"crypto_exchange_rate", cryptoExchangeRate,
		"crypto_currency_amount", cryptoCurrencyAmount)
	gatewayID, err := s.gatewayID(ctx)
	if err != nil {
		return models.Order{}, err
	}

	// create order
	order = models.Order{
		RetailerID:           retailer.ID,
		RetailerAccountID:    retailer.AccountID,
		AccountID:            account.ID,
		GatewayID:            gatewayID,
		GeneratedOrderID:     uuid.New().String(),
		ExternalOrderID:      orderDetail.MerchantTransactionID,
		Amount:               orderDetail.Amount,
		Currency:             orderDetail.Currency,
		ExchangeRate:         null.StringFrom(fmt.Sprintf("%v", exchangeRate)),
		BaseCurrency:         null.StringFrom("EUR"),
		BaseCurrencyAmount:   null.StringFrom(fmt.Sprintf("%v", baseCurrencyAmount)),
		Cryptocurrency:       null.StringFrom("USDT"),
		CryptoExchangeRate:   null.StringFrom(fmt.Sprintf("%v", cryptoExchangeRate)),
		CryptocurrencyAmount: null.StringFrom(fmt.Sprintf("%v", cryptoCurrencyAmount)),
		Type:                 models.OrderTypeBuy,
		Status:               models.OrderStatusCreated,
	}
	if err := s.CreateOrder(ctx, &order); err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (s *orderServiceImpl) OrdersForRetailer(ctx context.Context, r models.Retailer, start, end time.Time, page int) ([]models.Order, error) {
	return s.retailerRepo.GetOrdersByCreatedAt(ctx, &r, start, end, page, 1000)
}

func (s *orderServiceImpl) OrderDetailForRetailer(ctx context.Context, r models.Retailer, orderID int64) (services.RetailerOrderDetail, error) {
	order, err := s.orderRepo.PreloadFindById(ctx, orderID)
	if err != nil {
		return services.RetailerOrderDetail{}, err
	}
	payments, err := s.orderRepo.Payments(ctx, &order)
	if err != nil {
		return services.RetailerOrderDetail{}, err
	}
	retailer, err := s.retailerRepo.FindByAccount(ctx, &order.RetailerAccount)
	if err != nil {
		return services.RetailerOrderDetail{}, err
	}
	return services.RetailerOrderDetail{
		Order:           order,
		CustomerAccount: order.Account,
		RetailerAccount: order.RetailerAccount,
		Retailer:        retailer,
		Payments:        payments,
	}, nil
}

func (s *orderServiceImpl) SettlementAmount(ctx context.Context, r services.SettlementRequest) (services.SettlementResponse, error) {
	exchangeRate, err := s.baseCurrencyExchangeRate(ctx, r.Currency)
	if err != nil {
		logger.Error(ctx, "base currency exchange rate ", "error", err)
		return services.SettlementResponse{}, err
	}
	cryptoExchangeRate, err := s.cryptoExchangeRate(ctx)
	if err != nil {
		logger.Error(ctx, "crypto currency exchange rate ", "error", err)
		return services.SettlementResponse{}, err
	}
	baseCurrencyAmount := convertTo2Decimal(r.Amount * exchangeRate)
	cryptoCurrencyAmount := convertTo2Decimal(baseCurrencyAmount * cryptoExchangeRate)
	defaultPaymentMode, err := models.PaymentModeString(app.Config().DefaultPaymentMode)
	if err != nil {
		logger.Error(ctx, "payment mode from string ", "error", err)
		return services.SettlementResponse{}, err
	}
	settlement, err := s.settlementAmount(ctx, defaultPaymentMode, app.Config().DefaultRetailerID, cryptoCurrencyAmount)
	if err != nil {
		logger.Error(ctx, "error in creating settlement ", "error", err)
		return services.SettlementResponse{}, err
	}
	settlement.Currency = r.Currency
	return settlement, nil
}

func (s *orderServiceImpl) baseCurrencyExchangeRate(ctx context.Context, currency string) (float64, error) {
	rate, err := app.Cache().Instance().HGet(ctx, services.EuroExchangeRate, currency).Float64()
	if err != nil {
		return 0.0, err
	}
	logger.Info(ctx, "the rate is", "currency", currency, "rate", rate)
	return convertTo2Decimal(1 / rate), nil
}

func (s *orderServiceImpl) cryptoExchangeRate(ctx context.Context) (float64, error) {
	rate, err := app.Cache().Instance().HGet(ctx, services.EuroExchangeRate, "USDT").Float64()
	if err != nil {
		return 0.0, err
	}
	logger.Info(ctx, "the crypto rate is", "rate", rate)
	return convertTo2Decimal(rate), nil
}

func (s *orderServiceImpl) gatewayID(ctx context.Context) (int64, error) {
	return strconv.ParseInt(os.Getenv("PG_ID"), 10, 64)
}

func (s *orderServiceImpl) settlementAmount(ctx context.Context, mode models.PaymentMode, retailerId int64, amount float64) (services.SettlementResponse, error) {
	txnFees, err := s.txnFeesConfigRepo.FindByRetailerIDAndPaymentMode(ctx, retailerId, mode)
	if err != nil {
		return services.SettlementResponse{}, err
	}
	logger.Info(ctx, "crypto amount to be paid is", "floatAmount", amount)
	totalTxnFees := convertTo2Decimal(txnFees.FixedFees.Float64 + ((txnFees.PercentFees.Float64 / 100) * amount))
	totalPgTxnFees := convertTo2Decimal(txnFees.PgFixedFees.Float64 + ((txnFees.PgPercentageFees.Float64 / 100) * amount))
	settlementAmount := amount - (totalTxnFees + totalPgTxnFees)
	return services.SettlementResponse{
		SettlementAmount: settlementAmount,
		Fees:             totalTxnFees,
		PGFees:           totalPgTxnFees,
		Mode:             mode,
	}, nil
}
