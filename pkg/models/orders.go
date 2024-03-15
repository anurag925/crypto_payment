//go:generate go run github.com/dmarkham/enumer -type=OrderType -json -transform=snake -trimprefix=OrderType
//go:generate go run github.com/dmarkham/enumer -type=OrderStatus -json -transform=snake -trimprefix=OrderStatus
package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type OrderType int8

const (
	OrderTypeBuy OrderType = iota
	OrderTypeRefund
)

type OrderStatus int8

const (
	OrderStatusCreated OrderStatus = iota
	OrderStatusPaymentSuccess
	OrderStatusPaymentFailure
	OrderStatusSuccess
	OrderStatusFailure
)

type Order struct {
	ID                   int64          `gorm:"primaryKey" json:"id"`
	CreatedAt            time.Time      `json:"-"`
	UpdatedAt            time.Time      `json:"-"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`
	GeneratedOrderID     string         `gorm:"type:varchar(255);default:null" json:"generated_order_id"`
	ExternalOrderID      string         `gorm:"type:varchar(255);default:null" json:"external_order_id"`
	Type                 OrderType      `gorm:"type:smallint;default:0" json:"type"`   // buy=0, refund=1
	Status               OrderStatus    `gorm:"type:smallint;default:0" json:"status"` // created=0, payment_success=1, payment_failure=2, success=3, failure=4
	Amount               string         `gorm:"type:varchar(255);default:null" json:"amount"`
	Currency             string         `gorm:"type:varchar(255);default:null" json:"currency"`
	BaseCurrency         null.String    `gorm:"type:varchar(255);default:null" json:"base_currency"`         // Default EURO
	BaseCurrencyAmount   null.String    `gorm:"type:varchar(255);default:null" json:"base_currency_amount"`  // amount * exchange_rate
	ExchangeRate         null.String    `gorm:"type:varchar(255);default:null" json:"exchange_rate"`         // default 1
	Cryptocurrency       null.String    `gorm:"type:varchar(255);default:null" json:"cryptocurrency"`        // Default USDT
	CryptocurrencyAmount null.String    `gorm:"type:varchar(255);default:null" json:"cryptocurrency_amount"` // crypto_exchange_rate * base_currency_amount
	CryptoExchangeRate   null.String    `gorm:"type:varchar(255);default:null" json:"crypto_exchange_rate"`
	Account              Account        `gorm:"foreignKey:AccountID" json:"-"`
	AccountID            int64          `json:"account_id"`
	Retailer             Retailer       `gorm:"foreignKey:RetailerID" json:"-"`
	RetailerID           int64          `json:"retailer_id"`
	RetailerAccount      Account        `gorm:"foreignKey:RetailerAccountID" json:"-"`
	RetailerAccountID    int64          `json:"retailer_account_id"`
	Gateway              Gateway        `json:"-"`
	GatewayID            int64          `json:"gateway_id"`
}
