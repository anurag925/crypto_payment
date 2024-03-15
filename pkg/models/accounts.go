//go:generate go run github.com/dmarkham/enumer -type=AccountType -json -transform=snake -trimprefix=AccountType
//go:generate go run github.com/dmarkham/enumer -type=AccountStatus -json -transform=snake -trimprefix=AccountStatus
package models

import (
	"encoding/json"
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type AccountType int8

const (
	AccountTypeCustomer AccountType = iota
	AccountTypeRetailer
	AccountTypeAdmin
)

type AccountStatus int8

const (
	AccountStatusCreated AccountStatus = iota
	AccountStatusOnboarding
	AccountStatusActive
	AccountStatusBlocked
)

type Account struct {
	ID           int64          `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Type         AccountType    `gorm:"type:smallint;default:0" json:"type"`
	Status       AccountStatus  `gorm:"type:smallint;default:0" json:"status"`
	FirstName    null.String    `gorm:"type:varchar(255);default:null" json:"first_name"`
	LastName     null.String    `gorm:"type:varchar(255);default:null" json:"last_name"`
	Email        string         `gorm:"not null;type:varchar(255);default:null;unique" json:"email"`
	CountryCode  null.String    `gorm:"type:varchar(255);default:null" json:"country_code"`
	MobileNumber null.String    `gorm:"type:varchar(255);default:null" json:"mobile_number"`
	Password     null.String    `gorm:"type:varchar(255);default:null" json:"password"`
	TwoFA        bool           `gorm:"column:twofa;default:false" json:"twofa"`
	Country      null.String    `gorm:"type:varchar(255);default:null" json:"country"`
}

func (a Account) MarshalJSON() ([]byte, error) {
	type PasswordMaskedAccount Account
	passwordMaskedAccount := PasswordMaskedAccount(a)
	if passwordMaskedAccount.Password.Valid {
		passwordMaskedAccount.Password.SetValid("xxxxxxxxxxxxxxxxx")
	}
	return json.Marshal(passwordMaskedAccount)
}
