//go:generate go run github.com/dmarkham/enumer -type=AddressType -json -transform=snake -trimprefix=AddressType
//go:generate go run github.com/dmarkham/enumer -type=AddressStatus -json -transform=snake -trimprefix=AddressStatus
package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type AddressType int8

const (
	AddressTypeResidential AddressType = iota
	AddressTypeBusiness
)

type AddressStatus int8

const (
	AddressStatusActive AddressStatus = iota
	AddressStatusRemoved
)

type Address struct {
	ID           int64          `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Type         AddressType    `gorm:"type:smallint;default:0" json:"type"`   // residential=0, business=1
	Status       AddressStatus  `gorm:"type:smallint;default:0" json:"status"` // active=0, removed=1
	AddressLine1 null.String    `gorm:"type:varchar(255);default:null" json:"address_line_1"`
	AddressLine2 null.String    `gorm:"type:varchar(255);default:null" json:"address_line_2"`
	City         null.String    `gorm:"type:varchar(255);default:null" json:"city"`
	State        null.String    `gorm:"type:varchar(255);default:null" json:"state"`
	Country      null.String    `gorm:"type:varchar(255);default:null" json:"country"`
	PostalCode   null.String    `gorm:"type:varchar(255);default:null" json:"postal_code"`
	Account      Account        `gorm:"foreignKey:AccountID" json:"-"`
	AccountID    int64          `json:"-"`
}
