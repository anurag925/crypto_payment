package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type KycRepository interface {
	FindById(context.Context, int64) (models.Kyc, error)
	FindAll(ctx context.Context) ([]models.Kyc, error)
	Create(context.Context, *models.Kyc) error
	Documents(ctx context.Context, o *models.Kyc) (d []models.Document, err error)
}
