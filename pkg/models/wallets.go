package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Wallet struct {
	ID         int64          `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Balance    null.String    `gorm:"type:varchar(255);default:null" json:"balance"`
	Paid       null.String    `gorm:"type:varchar(255);default:null" json:"paid"`
	Pending    null.String    `gorm:"type:varchar(255);default:null" json:"pending"`
	Retailer   Retailer       `json:"-"`
	RetailerID int64          `json:"retailer_id"`
}
