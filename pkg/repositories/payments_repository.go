package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"time"
)

type PaymentRepository interface {
	FindById(context.Context, int64) (models.Payment, error)
	FindByGeneratedID(context.Context, string) (models.Payment, error)
	FindAll(ctx context.Context) ([]models.Payment, error)
	Create(context.Context, *models.Payment) error
	GetPaymentByCreatedAt(ctx context.Context, start, end time.Time, page, limit int) (p []models.Payment, err error)
	PaymentsForAccount(ctx context.Context, a models.Account) (p []models.Payment, err error)
}
