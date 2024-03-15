//go:generate go run github.com/dmarkham/enumer -type=GatewayStatus -json -transform=snake -trimprefix=GatewayStatus
package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type GatewayStatus int8

const (
	GatewayStatusCreated GatewayStatus = iota
	GatewayStatusActive
	GatewayStatusInactive
	GatewayStatusDeprecated
)

type Gateway struct {
	ID             int64          `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Status         GatewayStatus  `gorm:"type:smallint;default:0" json:"status"`
	Secret         null.String    `gorm:"size:255;default:null" json:"secret"`
	CreateURL      null.String    `gorm:"size:255;default:null" json:"create_url"`
	CallbackURL    null.String    `gorm:"size:255;default:null" json:"callback_url"`
	DowntimeURL    null.String    `gorm:"size:255;default:null" json:"downtime_url"`
	StatusCheckURL null.String    `gorm:"size:255;default:null" json:"status_check_url"`
}
