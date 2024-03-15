package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"time"
)

type AccountRepository interface {
	FindById(context.Context, int64) (models.Account, error)
	FindByEmail(context.Context, string) (models.Account, error)
	FindAll(ctx context.Context) ([]models.Account, error)
	Create(context.Context, *models.Account) error
	Retailer(ctx context.Context, a models.Account) (o models.Retailer, err error)
	GetAccountsByCreatedAt(ctx context.Context, start, end time.Time, page, limit int) (a []models.Account, err error)
	Kyc(ctx context.Context, o *models.Account) (k models.Kyc, err error)
	ResidentialAddress(ctx context.Context, o *models.Account) (a models.Address, err error)
	BusinessAddress(ctx context.Context, o *models.Account) (a models.Address, err error)
}
