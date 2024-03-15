package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"time"

	"gorm.io/gorm"
)

type paymentRepositoryImpl struct {
	*RepositoryImpl[models.Payment]
}

var _ repositories.PaymentRepository = (*paymentRepositoryImpl)(nil)

func NewPaymentRepositoryImpl(db *gorm.DB) *paymentRepositoryImpl {
	return &paymentRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Payment](db)}
}

func DefaultPaymentRepositoryImpl() *paymentRepositoryImpl {
	return NewPaymentRepositoryImpl(app.DB().Instance())
}

func (r *paymentRepositoryImpl) FindByGeneratedID(ctx context.Context, generatedID string) (o models.Payment, err error) {
	err = r.db.WithContext(ctx).Preload("Order").
		Preload("Order.Account").Preload("Order.RetailerAccount").
		Where("generated_id = ?", generatedID).First(&o).Error
	return
}

func (r *paymentRepositoryImpl) GetPaymentByCreatedAt(ctx context.Context, start, end time.Time, page, limit int) (p []models.Payment, err error) {
	err = r.db.WithContext(ctx).Preload("Order").
		Preload("Order.Account").
		Preload("Order.RetailerAccount").
		Where("created_at BETWEEN ? AND ?", start, end).
		Offset((page - 1) * limit).Limit(limit).Find(&p).Error
	return
}

func (r *paymentRepositoryImpl) PaymentsForAccount(ctx context.Context, a models.Account) (p []models.Payment, err error) {
	err = r.db.WithContext(ctx).InnerJoins("Order").Where("account_id = ?", a.ID).Find(&p).Error
	return
}
