package libs

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type PaymentCreateRequest struct {
	Account        models.Account
	Order          models.Order
	Payment        models.Payment
	BrowserDetails models.BrowserDetails
	Signature      string
}

type PaymentLib interface {
	Create(ctx context.Context, r PaymentCreateRequest) (ZenPaymentCreateResponse, error)
	Callback(ctx context.Context, request CallbackRequest) error
}
