package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Contact struct {
	ID                 int64          `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
	UBO                bool           `gorm:"default:false" json:"ubo"`
	PointOfContact     bool           `gorm:"default:false" json:"point_of_contact"`
	FullName           null.String    `gorm:"type:varchar(255);default:null" json:"full_name"`
	DOB                null.String    `gorm:"type:varchar(255);default:null" json:"dob"`
	Country            null.String    `gorm:"type:varchar(255);default:null" json:"country"`
	CountryOfBirth     null.String    `gorm:"type:varchar(255);default:null" json:"country_of_birth"`
	PercentageHoldings null.Float     `gorm:"type:decimal;default:null" json:"percentage_holdings"`
	IDType             null.String    `gorm:"type:varchar(255);default:null" json:"id_type"`
	IDNo               null.String    `gorm:"type:varchar(255);default:null" json:"id_no"`
	Retailer           Retailer       `json:"-"`
	RetailerID         int64          `json:"retailer_id"`
}
