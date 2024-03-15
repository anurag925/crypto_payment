package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"time"
)

type RetailerRepository interface {
	FindById(context.Context, int64) (models.Retailer, error)
	FindAll(ctx context.Context) ([]models.Retailer, error)
	FindByAccount(ctx context.Context, a *models.Account) (o models.Retailer, err error)
	Account(ctx context.Context, o *models.Retailer) (a models.Account, err error)
	Wallet(ctx context.Context, o *models.Retailer) (w models.Wallet, err error)
	TxnFeeConfigs(ctx context.Context, o *models.Retailer) (t []models.TxnFeeConfig, err error)
	TxnFeeConfigsByPaymentMode(ctx context.Context, o *models.Retailer, mode models.PaymentMode) (t models.TxnFeeConfig, err error)
	CryptoWallets(ctx context.Context, o *models.Retailer) (w []models.CryptoWallet, err error)
	Contact(ctx context.Context, o *models.Retailer) (c models.Contact, err error)
	Shareholders(ctx context.Context, o *models.Retailer) (c []models.Contact, err error)
	Kyc(ctx context.Context, o *models.Retailer) (k models.Kyc, err error)
	Documents(ctx context.Context, o *models.Retailer) (d []models.Document, err error)
	Address(ctx context.Context, o *models.Retailer) (a models.Address, err error)
	GetOrdersByCreatedAt(ctx context.Context, m *models.Retailer, start, end time.Time, page, limit int) (o []models.Order, err error)
	Create(context.Context, *models.Retailer) error
	Save(context.Context, *models.Retailer) error
}
