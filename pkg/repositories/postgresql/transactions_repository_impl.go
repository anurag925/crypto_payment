package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type transactionRepositoryImpl struct {
	*RepositoryImpl[models.Transaction]
}

var _ repositories.TransactionsRepository = (*transactionRepositoryImpl)(nil)

func NewTransactionRepositoryImpl(db *gorm.DB) *transactionRepositoryImpl {
	return &transactionRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Transaction](db)}
}

func DefaultTransactionRepositoryImpl() *transactionRepositoryImpl {
	return NewTransactionRepositoryImpl(app.DB().Instance())
}

func (s *transactionRepositoryImpl) AllTransactionsForRetailer(ctx context.Context, r models.Retailer) (t []models.Transaction, err error) {
	err = s.db.WithContext(ctx).Where("retailer_id = ?", r.ID).Find(&t).Error
	return
}

func (s *transactionRepositoryImpl) CryptoWallet(ctx context.Context, t models.Transaction) (w models.CryptoWallet, err error) {
	err = s.db.WithContext(ctx).Model(&t).Association("CryptoWallet").Find(&w)
	return
}
