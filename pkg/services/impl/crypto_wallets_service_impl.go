package impl

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
)

type cryptoWalletServiceImpl struct {
	cryptoWalletRepo repositories.CryptoWalletRepository
}

func NewCryptoWalletService(repo repositories.CryptoWalletRepository) *cryptoWalletServiceImpl {
	return &cryptoWalletServiceImpl{cryptoWalletRepo: repo}
}

func DefaultCryptoWalletServiceImpl() *cryptoWalletServiceImpl {
	return NewCryptoWalletService(postgresql.DefaultCryptoWalletRepositoryImpl())
}

func (s *cryptoWalletServiceImpl) FindById(ctx context.Context, id int64) (models.CryptoWallet, error) {
	return s.cryptoWalletRepo.FindById(ctx, id)
}
func (s *cryptoWalletServiceImpl) CryptoWalletsForRetailer(ctx context.Context, o models.Retailer) ([]models.CryptoWallet, error) {
	return s.cryptoWalletRepo.FindByRetailerAndActive(ctx, &o)
}
func (s *cryptoWalletServiceImpl) CreateCryptoWallet(ctx context.Context, o models.Retailer, cryptoWallet *models.CryptoWallet) error {
	cryptoWallet.RetailerID = o.ID
	return s.cryptoWalletRepo.Save(ctx, cryptoWallet)
}
