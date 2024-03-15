package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/libs"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type PaymentCreateRequest struct {
	OrderId        string                `json:"order_id"`
	Payment        models.Payment        `json:"payment"`
	BrowserDetails models.BrowserDetails `json:"browser_details"`
	Signature      string                `json:"signature"`
}

type PaymentCreateResponse struct {
	Payment     models.Payment                `json:"payment"`
	ZenResponse libs.ZenPaymentCreateResponse `json:"zen_response"`
}

type PaymentStatus struct {
}

type Transaction struct {
}

type Transactions struct {
	Transactions []Transaction
}

type PaymentService interface {
	Create(ctx context.Context, request PaymentCreateRequest) (PaymentCreateResponse, error)
	Status(ctx context.Context, paymentId int64) (models.Payment, error)
	Transactions() (Transactions, error)
	Callback() error
	PaymentsForAccount(ctx context.Context, email string) ([]models.Payment, error)
}
