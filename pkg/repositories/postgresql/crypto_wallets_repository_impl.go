package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type cryptoWalletRepositoryImpl struct {
	*RepositoryImpl[models.CryptoWallet]
}

var _ repositories.CryptoWalletRepository = (*cryptoWalletRepositoryImpl)(nil)

func NewCryptoWalletRepositoryImpl(db *gorm.DB) *cryptoWalletRepositoryImpl {
	return &cryptoWalletRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.CryptoWallet](db)}
}

func DefaultCryptoWalletRepositoryImpl() *cryptoWalletRepositoryImpl {
	return NewCryptoWalletRepositoryImpl(app.DB().Instance())
}

func (r *cryptoWalletRepositoryImpl) FindByRetailerAndActive(ctx context.Context, o *models.Retailer) (a []models.CryptoWallet, err error) {
	err = r.db.Where("retailer_id = ? and status = ?", o.ID, int8(models.CryptoWalletStatusActive)).Find(&a).Error
	return
}
