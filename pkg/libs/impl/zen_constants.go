package impl

import (
	"errors"
	"github.com/anurag925/crypto_payment/pkg/models"
	"net/http"
	"time"
)

var (
	zenPaymentApiRetryStatus   = []int{http.StatusBadGateway, http.StatusGatewayTimeout}
	zenPaymentApiRetryWaitTime = 2 * time.Second
	zenPaymentApiRetryCount    = 3
)

var (
	ErrCurrencyNotSupported     = errors.New("currency not supported for payment method")
	ErrInvalidAmount            = errors.New("bad amount given for payment method")
	ErrInvalidGateway           = errors.New("invalid gateway type")
	ErrPaymentCreateAtZenFailed = errors.New("zen payment request failed")
)

type PaymentMethodConfigs struct {
	Name        string
	ChannelCode string
	Currencies  []string
	MinTxn      float64
	MaxTxn      float64
	GatewayType string
}

var zenPaymentStatus = map[string]models.PaymentStatus{
	"AUTHORIZED": models.PaymentStatusPending,
	"PENDING":    models.PaymentStatusPending,
	"REJECTED":   models.PaymentStatusRejected,
	"CANCELED":   models.PaymentStatusCanceled,
	"ACCEPTED":   models.PaymentStatusCompleted,
}

var zenPaymentType = map[string]models.OrderType{
	"TRT_PURCHASE": models.OrderTypeBuy,
	"TRT_REFUND":   models.OrderTypeRefund,
}

var paymentMethods = map[string]PaymentMethodConfigs{
	"apple_pay": {
		"ApplePay",
		"PCL_APPLEPAY",
		[]string{"EUR"},
		0.01,
		10000.0,
		"external_payment_token",
	},
	"google_pay": {
		"GooglePay",
		"PCL_GOOGLEPAY",
		[]string{"EUR"},
		0.01,
		100000.0,
		"external_payment_token",
	},
	"mastercard": {
		"Mastercard",
		"PCL_CARD",
		[]string{"AED", "AUD", "BGN", "CAD", "CHF", "CNY", "CZK", "DKK", "EUR", "GBP", "HKD", "HUF", "ILS", "JPY", "KES", "MXN", "NOK", "NZD", "PLN", "QAR", "RON", "SAR", "SEK", "SGD", "THB", "TRY", "UGX", "USD", "ZAR"},
		0.01,
		100000.0,
		"onetime",
	},
	"visa": {
		"Visa",
		"PCL_CARD",
		[]string{"USD", "AUD", "GBP", "BGN", "CAD", "CZK", "DKK", "EUR", "HKD", "HUF", "ILS", "JPY", "KES", "MXN", "NZD", "NOK", "PLN", "QAR", "CNY", "RON", "SAR", "SGD", "ZAR", "SEK", "CHF", "THB", "TRY", "UGX", "AED"},
		0.01,
		100000.0,
		"onetime",
	},
	"zen_card": {
		"ZenCard",
		"PCL_CARD",
		[]string{"AED", "AUD", "BGN", "CAD", "CHF", "CNY", "CZK", "DKK", "EUR", "GBP", "HKD", "HUF", "ILS", "JPY", "KES", "MXN", "NOK", "NZD", "PLN", "QAR", "RON", "SAR", "SEK", "SGD", "THB", "TRY", "UGX", "USD", "ZAR"},
		0.01,
		100000.0,
		"onetime",
	},
	"pay_by_zen": {
		"PayByZen",
		"PCL_PBZ",
		[]string{"EUR", "GBP", "PLN", "USD"},
		0.01,
		1000000.0,
		"general",
	},
	"trustly": {
		"Trustly",
		"PCL_TRUSTLY",
		[]string{"EUR", "PLN", "DKK", "SEK", "NOK", "CZK"},
		0.01,
		100000.0,
		"trustly",
	},
	"bancontact": {
		"Bancontact",
		"PCL_BANCONTACT",
		[]string{"EUR"},
		0.01,
		10000.0,
		"general",
	},
	"sofort": {
		"Sofort",
		"PCL_SOFORT",
		[]string{"EUR"},
		0.01,
		10000.0,
		"general",
	},
	"ideal": {
		"Ideal",
		"PCL_IDEAL",
		[]string{"EUR"},
		0.01,
		10000.0,
		"general",
	},
	"blik": {
		"BLIK",
		"PCl_BLIK_REDIRECT",
		[]string{"PLN"},
		0.01,
		100000.0,
		"general",
	},
}

var paymentModeToZenType = map[models.PaymentMode]string{
	models.PaymentModeApplePay:   "apple_pay",
	models.PaymentModeGooglePay:  "google_pay",
	models.PaymentModeBancontact: "bancontact",
	models.PaymentModeBlik:       "blik",
	models.PaymentModeCard:       "mastercard",
	models.PaymentModeSofort:     "sofort",
	models.PaymentModeIdeal:      "ideal",
	models.PaymentModeTrustly:    "trustly",
}
