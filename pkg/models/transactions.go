//go:generate go run github.com/dmarkham/enumer -type=TransactionStatus -json -transform=snake -trimprefix=TransactionStatus
//go:generate go run github.com/dmarkham/enumer -type=TransactionType -json -transform=snake -trimprefix=TransactionType
package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type TransactionStatus int8

const (
	TransactionStatusCreated TransactionStatus = iota
	TransactionStatusPending
	TransactionStatusCompleted
	TransactionStatusRefunded
	TransactionStatusRejected
	TransactionStatusCanceled
)

type TransactionType int8

const (
	TransactionTypeCredit TransactionType = iota
	TransactionTypeDebit
)

type Transaction struct {
	ID             int64             `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time         `json:"-"`
	UpdatedAt      time.Time         `json:"-"`
	DeletedAt      gorm.DeletedAt    `gorm:"index" json:"-"`
	Status         TransactionStatus `gorm:"type:smallint;default:0" json:"status"`
	Type           TransactionType   `gorm:"type:smallint;default:0" json:"type"`
	Amount         null.String       `gorm:"type:varchar(255);default:null" json:"amount"`
	Balance        null.String       `gorm:"type:varchar(255);default:null" json:"balance"`
	Pending        null.String       `gorm:"type:varchar(255);default:null" json:"pending"`
	CryptoWallet   CryptoWallet      `json:"-"`
	CryptoWalletID int64             `json:"crypto_wallet_id"`
	Retailer       Retailer          `json:"-"`
	RetailerID     int64             `json:"retailer_id"`
}
