package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type CryptoWalletService interface {
	FindById(ctx context.Context, id int64) (models.CryptoWallet, error)
	CryptoWalletsForRetailer(ctx context.Context, o models.Retailer) ([]models.CryptoWallet, error)
	CreateCryptoWallet(ctx context.Context, o models.Retailer, cryptoWallet *models.CryptoWallet) error
}
