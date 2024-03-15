package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type ContactRepository interface {
	FindById(context.Context, int64) (models.Contact, error)
	FindAll(ctx context.Context) ([]models.Contact, error)
	Create(context.Context, *models.Contact) error
}
