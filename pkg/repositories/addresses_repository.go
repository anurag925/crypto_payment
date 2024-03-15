package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type AddressRepository interface {
	FindById(context.Context, int64) (models.Address, error)
	FindAll(ctx context.Context) ([]models.Address, error)
	Create(context.Context, *models.Address) error
}
