package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/libs"
)

type InternalPaymentService interface {
	Create(ctx context.Context) error
	Status(ctx context.Context) error
	Transactions(ctx context.Context) error
	Callback(ctx context.Context, request libs.CallbackRequest) error
}
