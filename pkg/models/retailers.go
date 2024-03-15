//go:generate go run github.com/dmarkham/enumer -type=RetailerCategory -json -transform=snake -trimprefix=RetailerCategory
package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type RetailerCategory int8

const (
	RetailerCategoryUnknown RetailerCategory = iota
)

type Retailer struct {
	ID              int64            `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt       time.Time        `json:"-"`
	UpdatedAt       time.Time        `json:"-"`
	DeletedAt       gorm.DeletedAt   `gorm:"index" json:"-"`
	Category        RetailerCategory `gorm:"type:smallint;default:0" json:"category"`
	RetailName      null.String      `gorm:"type:varchar(255);default:null" json:"retail_name"`
	Website         null.String      `gorm:"type:varchar(255);default:null" json:"website"`
	PaymentGateway  null.String      `gorm:"type:varchar(255);default:null" json:"payment_gateway"`
	DailyLimit      null.String      `gorm:"type:varchar(255);default:null" json:"daily_limit"`
	WeeklyLimit     null.String      `gorm:"type:varchar(255);default:null" json:"weekly_limit"`
	MonthlyLimit    null.String      `gorm:"type:varchar(255);default:null" json:"monthly_limit"`
	SingleTxnLimit  null.String      `gorm:"type:varchar(255);default:null" json:"single_txn_limit"`
	OverallTxnValue null.String      `gorm:"type:varchar(255);default:null" json:"overall_txn_value"`
	MonthlyTxnValue null.String      `gorm:"type:varchar(255);default:null" json:"monthly_txn_value"`
	Account         Account          `json:"-"`
	AccountID       int64            `json:"account_id"`
}
