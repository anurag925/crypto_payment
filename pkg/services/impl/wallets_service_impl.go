package impl

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
)

type walletServiceImpl struct {
	walletRepo   repositories.WalletRepository
	retailerRepo repositories.RetailerRepository
}

var _ services.WalletService = (*walletServiceImpl)(nil)

func NewWalletServiceImpl(walletRepo repositories.WalletRepository, retailerRepo repositories.RetailerRepository) *walletServiceImpl {
	return &walletServiceImpl{walletRepo: walletRepo, retailerRepo: retailerRepo}
}

func DefaultWalletServiceImpl() *walletServiceImpl {
	return NewWalletServiceImpl(postgresql.DefaultWalletRepositoryImpl(), postgresql.DefaultRetailerRepositoryImpl())
}

func (s *walletServiceImpl) WalletForRetailer(ctx context.Context, r models.Retailer) (w models.Wallet, err error) {
	return s.retailerRepo.Wallet(ctx, &r)
}
