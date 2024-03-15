package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type txnFeeConfigRepositoryImpl struct {
	*RepositoryImpl[models.TxnFeeConfig]
}

var _ repositories.TxnFeeConfigRepository = (*txnFeeConfigRepositoryImpl)(nil)

func NewTxnFeeConfigRepositoryImpl(db *gorm.DB) *txnFeeConfigRepositoryImpl {
	return &txnFeeConfigRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.TxnFeeConfig](db)}
}

func DefaultTxnFeeConfigRepositoryImpl() *txnFeeConfigRepositoryImpl {
	return NewTxnFeeConfigRepositoryImpl(app.DB().Instance())
}

func (r *txnFeeConfigRepositoryImpl) FindByRetailerID(ctx context.Context, id int64) (t models.TxnFeeConfig, err error) {
	err = r.db.WithContext(ctx).Where("retailer_id = ?", id).First(&t).Error
	return
}

func (r *txnFeeConfigRepositoryImpl) FindByRetailerIDAndPaymentMode(ctx context.Context, id int64, mode models.PaymentMode) (t models.TxnFeeConfig, err error) {
	err = r.db.WithContext(ctx).Where("retailer_id = ? and type = ?", id, mode).First(&t).Error
	return
}
