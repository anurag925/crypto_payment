package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"time"
)

type OrderCreateRequest struct {
	ApiKey      string      `json:"api_key" validate:"required"`
	OrderDetail OrderDetail `json:"order_detail" validate:"required"`
	Checksum    string      `json:"checksum"`
	Internal    bool        `json:"internal"`
}

type OrderDetail struct {
	CustomerID            string `json:"customer_id" validate:"required"`
	MerchantTransactionID string `json:"merchant_transaction_id" validate:"required"`
	FirstName             string `json:"first_name" validate:"required"`
	LastName              string `json:"last_name" validate:"required"`
	Email                 string `json:"email" validate:"required"`
	CountryCode           string `json:"country_code" validate:"required"`
	MobileNumber          string `json:"mobile_number" validate:"required"`
	AddressLine1          string `json:"address_line_1" validate:"required"`
	AddressLine2          string `json:"address_line_2"`
	City                  string `json:"city" validate:"required"`
	State                 string `json:"state"`
	Country               string `json:"country" validate:"required"`
	Amount                string `json:"amount" validate:"required"`
	Currency              string `json:"currency" validate:"required"`
	// PaymentMethod         string `json:"payment_method" validate:"required"`
	// RedirectURL           string `json:"redirect_url" validate:"required"`
}

type RetailerOrderDetail struct {
	Order           models.Order     `json:"order"`
	CustomerAccount models.Account   `json:"customer_account"`
	Retailer        models.Retailer  `json:"retailer"`
	RetailerAccount models.Account   `json:"retailer_account"`
	Payments        []models.Payment `json:"payments"`
}

type SettlementRequest struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type SettlementResponse struct {
	SettlementAmount float64            `json:"settlement_amount"`
	Currency         string             `json:"currency"`
	Fees             float64            `json:"fees"`
	PGFees           float64            `json:"pg_fees"`
	Mode             models.PaymentMode `json:"mode"`
}

type OrderService interface {
	CreateOrder(ctx context.Context, a *models.Order) error
	CreateBuyOrder(ctx context.Context, req OrderCreateRequest) (models.Order, error)
	OrdersForRetailer(ctx context.Context, r models.Retailer, start, end time.Time, page int) ([]models.Order, error)
	OrderDetailForRetailer(ctx context.Context, r models.Retailer, orderID int64) (RetailerOrderDetail, error)
	SettlementAmount(ctx context.Context, r SettlementRequest) (SettlementResponse, error)
}
