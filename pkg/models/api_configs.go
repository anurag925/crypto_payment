package models

import (
	"time"

	"gorm.io/gorm"
)

type ApiConfig struct {
	ID          int64          `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Key         string         `gorm:"type:varchar(255);default:null;index" json:"key"`
	Secret      string         `gorm:"type:varchar(255);default:null" json:"secret"`
	CallbackURL string         `gorm:"type:varchar(255);default:null" json:"callback_url"`
	Retailer    Retailer       `json:"-"`
	RetailerID  int64          `json:"retailer_id"`
}
