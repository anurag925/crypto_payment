package impl

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/libs"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
)

type internalPaymentServiceImpl struct {
	paymentRepo repositories.PaymentRepository
}

var _ services.InternalPaymentService = (*internalPaymentServiceImpl)(nil)

func NewInternalPaymentServiceImpl(paymentRepo repositories.PaymentRepository) *internalPaymentServiceImpl {
	return &internalPaymentServiceImpl{paymentRepo: paymentRepo}
}

func DefaultInternalPaymentServiceImpl() *internalPaymentServiceImpl {
	return NewInternalPaymentServiceImpl(postgresql.DefaultPaymentRepositoryImpl())
}

func (s *internalPaymentServiceImpl) Create(ctx context.Context) error { return nil }
func (s *internalPaymentServiceImpl) Status(ctx context.Context) error { return nil }
func (s *internalPaymentServiceImpl) Transactions(ctx context.Context) error {
	return nil
}
func (s *internalPaymentServiceImpl) Callback(ctx context.Context, request libs.CallbackRequest) error {
	return nil
}
