package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type WalletService interface {
	WalletForRetailer(ctx context.Context, r models.Retailer) (w models.Wallet, err error)
}
