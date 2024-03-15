package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Document struct {
	ID        int64          `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Type      null.String    `gorm:"type:varchar(255);default:null" json:"type"`
	SubType   null.String    `gorm:"type:varchar(255);default:null" json:"sub_type"`
	URL       null.String    `gorm:"type:varchar(255);default:null" json:"url"`
	Data      any            `gorm:"type:jsonb" json:"data"`
	Verified  bool           `gorm:"default:false" json:"verified"`
	KycID     int64          `json:"kyc_id"`
	Kyc       Kyc            `json:"-"`
}
