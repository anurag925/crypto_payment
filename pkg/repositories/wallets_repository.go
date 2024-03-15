package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type WalletRepository interface {
	FindById(context.Context, int64) (models.Wallet, error)
	FindAll(ctx context.Context) ([]models.Wallet, error)
	Create(context.Context, *models.Wallet) error
	Save(context.Context, *models.Wallet) error
}
