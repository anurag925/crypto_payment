package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type CryptoWalletRepository interface {
	FindById(context.Context, int64) (models.CryptoWallet, error)
	FindAll(ctx context.Context) ([]models.CryptoWallet, error)
	Create(context.Context, *models.CryptoWallet) error
	Save(ctx context.Context, o *models.CryptoWallet) (err error)
	FindByRetailerAndActive(ctx context.Context, o *models.Retailer) (a []models.CryptoWallet, err error)
}
