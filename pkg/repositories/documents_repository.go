package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type DocumentRepository interface {
	FindById(context.Context, int64) (models.Document, error)
	FindAll(ctx context.Context) ([]models.Document, error)
	Create(context.Context, *models.Document) error
}
