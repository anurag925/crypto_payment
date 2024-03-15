//go:generate go run github.com/dmarkham/enumer -type=PaymentStatus -json -transform=snake -trimprefix=PaymentStatus
//go:generate go run github.com/dmarkham/enumer -type=PaymentMode -json -transform=snake -trimprefix=PaymentMode
//go:generate go run github.com/dmarkham/enumer -type=PaymentPlatform -json -transform=snake -trimprefix=PaymentPlatform
package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type PaymentStatus int8

const (
	PaymentStatusCreated PaymentStatus = iota
	PaymentStatusPending
	PaymentStatusCompleted
	PaymentStatusRejected
	PaymentStatusExpired
	PaymentStatusCanceled
	PaymentStatusRefunded
)

type PaymentMode int8

const (
	PaymentModeCard PaymentMode = iota
	PaymentModeGooglePay
	PaymentModeApplePay
	PaymentModeSofort
	PaymentModeTrustly
	PaymentModeBancontact
	PaymentModeIdeal
	PaymentModeBlik
)

type PaymentPlatform int8

const (
	PaymentPlatformWeb PaymentPlatform = iota
	PaymentPlatformAndroid
	PaymentPlatformIos
)

type Payment struct {
	ID               int64           `gorm:"primaryKey" json:"id"`
	CreatedAt        time.Time       `json:"-"`
	UpdatedAt        time.Time       `json:"-"`
	DeletedAt        gorm.DeletedAt  `gorm:"index" json:"-"`
	GeneratedID      string          `gorm:"type:varchar(255);default:null;index" json:"generated_id"`
	Amount           string          `gorm:"type:varchar(255);default:null" json:"amount"`
	Status           PaymentStatus   `gorm:"type:smallint;default:0" json:"status"`
	Mode             PaymentMode     `gorm:"type:smallint;default:0" json:"mode"`
	Platform         PaymentPlatform `gorm:"type:smallint;default:0" json:"platform"`
	Data             any             `gorm:"type:jsonb" json:"data"`
	Notes            null.String     `gorm:"type:varchar(255);default:null" json:"notes"`
	SettlementTime   null.Time       `gorm:"default:null" json:"settlement_time"`
	RefundTime       null.Time       `gorm:"default:null" json:"refund_time"`
	PgTransactionNo  null.String     `gorm:"type:varchar(255);default:null" json:"pg_transaction_no"`
	TxnReferenceNo   null.String     `gorm:"type:varchar(255);default:null" json:"txn_reference_no"`
	TxnFees          null.String     `gorm:"type:varchar(255);default:null" json:"txn_fees"`
	PGFees           null.String     `gorm:"type:varchar(255);default:null" json:"pg_fees"`
	Order            Order           `json:"-"`
	OrderID          int64           `json:"order_id"`
	RedirectUrl      string          `json:"redirect_url"`
	SettlementAmount string          `json:"settlement_amount"`
	IPAddress        string          `json:"ip_address"`
}

type CardData struct {
	Number     string `json:"number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        string `json:"cvv"`
}

type TokenData struct {
	Token string `json:"token"`
}

type BrowserDetails struct {
	AcceptHeader string `json:"acceptHeader"`
	ColorDepth   string `json:"colorDepth"`
	JavaEnabled  bool   `json:"javaEnabled"`
	Lang         string `json:"lang"`
	ScreenHeight string `json:"screenHeight"`
	ScreenWidth  string `json:"screenWidth"`
	Timezone     string `json:"timezone"`
	WindowSize   string `json:"windowSize"`
	UserAgent    string `json:"userAgent"`
}
