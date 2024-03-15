//go:generate go run github.com/dmarkham/enumer -type=CryptoWalletStatus -json -transform=snake -trimprefix=CryptoWalletStatus
package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type CryptoWalletStatus int8

const (
	CryptoWalletStatusActive CryptoWalletStatus = iota
	CryptoWalletStatusDeactivated
)

type CryptoWallet struct {
	ID               int64              `gorm:"primaryKey" json:"id"`
	CreatedAt        time.Time          `json:"-"`
	UpdatedAt        time.Time          `json:"-"`
	DeletedAt        gorm.DeletedAt     `gorm:"index" json:"-"`
	Status           CryptoWalletStatus `gorm:"type:smallint;default:0" json:"status,omitempty"`
	EncWalletAddress null.String        `gorm:"type:varchar(255);default:null" json:"enc_wallet_address,omitempty"`
	Network          null.String        `gorm:"type:varchar(255);default:null" json:"network,omitempty"`
	Retailer         Retailer           `json:"-"`
	RetailerID       int64              `json:"retailer_id"`
}
