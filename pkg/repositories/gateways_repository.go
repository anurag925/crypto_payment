package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type GatewayRepository interface {
	FindById(context.Context, int64) (models.Gateway, error)
	FindAll(ctx context.Context) ([]models.Gateway, error)
	Create(context.Context, *models.Gateway) error
}
