package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type ApiConfigRepository interface {
	FindById(ctx context.Context, id int64) (models.ApiConfig, error)
	FindByKey(ctx context.Context, key string) (models.ApiConfig, error)
	FindByRetailer(ctx context.Context, o *models.Retailer) (a []models.ApiConfig, err error)
	FindAll(ctx context.Context) ([]models.ApiConfig, error)
	Create(context.Context, *models.ApiConfig) error
}
