//go:generate go run github.com/dmarkham/enumer -type=OtpType -json -transform=snake -trimprefix=OtpType
//go:generate go run github.com/dmarkham/enumer -type=OtpAction -json -transform=snake -trimprefix=OtpAction
package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type OtpType int8

const (
	OtpTypeEmail OtpType = iota
	OtpTypeMobile
)

type OtpAction int8

const (
	OtpActionTransactionOtp OtpAction = iota
)

type Otp struct {
	ID            int64          `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Receiver      string         `gorm:"type:varchar(255)" json:"receiver"`
	Type          OtpType        `gorm:"type:smallint;default:0" json:"type"`
	Value         int            `json:"value"`
	Action        OtpAction      `gorm:"type:smallint;default:0" json:"action"`
	Verified      bool           `gorm:"type:boolean;default:false" json:"status"`
	VerifiedAt    null.Time      `gorm:"type:timestamptz;default:null" json:"verified_at"`
	VerifyingID   null.Int       `json:"verifying_id"`
	VerifyingType null.String    `gorm:"type:varchar(255);default:null" json:"verifying_type"`
	RetryCount    int            `json:"retry_count"`
}
