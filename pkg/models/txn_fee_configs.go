package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type TxnFeeConfig struct {
	ID               int64          `gorm:"primaryKey" json:"id"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Type             PaymentMode    `gorm:"type:smallint;default:0" json:"type"`
	PercentFees      null.Float     `gorm:"type:decimal;default:null" json:"percent_fees"`
	FixedFees        null.Float     `gorm:"type:decimal;default:null" json:"fixed_fees"`
	PgPercentageFees null.Float     `gorm:"type:decimal;default:null" json:"pg_percentage_fees"`
	PgFixedFees      null.Float     `gorm:"type:decimal;default:null" json:"pg_fixed_fees"`
	Retailer         Retailer       `json:"-"`
	RetailerID       int64          `json:"retailer_id"`
}
