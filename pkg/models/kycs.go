//go:generate go run github.com/dmarkham/enumer -type=KycStatus -json -transform=snake -trimprefix=KycStatus
package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type KycStatus int8

const (
	KycStatusPending KycStatus = iota
	KycStatusVerified
	KycStatusFailed
)

type Kyc struct {
	ID         int64          `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Status     KycStatus      `gorm:"type:smallint;default:0" json:"status"`
	VerifiedAt null.Time      `gorm:"type:timestamptz;default:null" json:"verified_at"`
	FailedAt   null.Time      `gorm:"type:timestamptz;default:null" json:"failed_at"`
	Account    Account        `json:"-"`
	AccountID  int64          `json:"account_id"`
}
