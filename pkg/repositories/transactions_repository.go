package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type TransactionsRepository interface {
	FindById(context.Context, int64) (models.Transaction, error)
	FindAll(ctx context.Context) ([]models.Transaction, error)
	Create(context.Context, *models.Transaction) error
	Save(context.Context, *models.Transaction) error
	AllTransactionsForRetailer(ctx context.Context, r models.Retailer) ([]models.Transaction, error)
	CryptoWallet(ctx context.Context, t models.Transaction) (w models.CryptoWallet, err error)
}
