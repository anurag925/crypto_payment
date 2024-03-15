package postgresql

import (
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type walletRepositoryImpl struct {
	*RepositoryImpl[models.Wallet]
}

var _ repositories.WalletRepository = (*walletRepositoryImpl)(nil)

func NewWalletRepositoryImpl(db *gorm.DB) *walletRepositoryImpl {
	return &walletRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Wallet](db)}
}

func DefaultWalletRepositoryImpl() *walletRepositoryImpl {
	return NewWalletRepositoryImpl(app.DB().Instance())
}
