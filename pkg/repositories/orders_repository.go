package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type OrderRepository interface {
	FindById(context.Context, int64) (models.Order, error)
	PreloadFindById(ctx context.Context, id int64) (t models.Order, err error)
	FindByExternalOrderID(context.Context, string) (models.Order, error)
	FindByGeneratedOrderID(context.Context, string) (models.Order, error)
	FindAll(ctx context.Context) ([]models.Order, error)
	Create(context.Context, *models.Order) error
	Payments(ctx context.Context, o *models.Order) (p []models.Payment, err error)
	Account(ctx context.Context, o *models.Order) (a models.Account, err error)
}
