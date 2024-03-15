package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type orderRepositoryImpl struct {
	*RepositoryImpl[models.Order]
}

var _ repositories.OrderRepository = (*orderRepositoryImpl)(nil)

func NewOrderRepositoryImpl(db *gorm.DB) *orderRepositoryImpl {
	return &orderRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Order](db)}
}

func DefaultOrderRepositoryImpl() *orderRepositoryImpl {
	return NewOrderRepositoryImpl(app.DB().Instance())
}

func (r *orderRepositoryImpl) FindByExternalOrderID(ctx context.Context, externalOrderID string) (o models.Order, err error) {
	err = r.db.WithContext(ctx).Where("external_order_id = ?", externalOrderID).First(&o).Error
	return
}

func (r *orderRepositoryImpl) Payments(ctx context.Context, o *models.Order) (p []models.Payment, err error) {
	err = r.db.WithContext(ctx).Where("order_id = ?", o.ID).Find(&p).Error
	return
}

func (r *orderRepositoryImpl) FindByGeneratedOrderID(ctx context.Context, generatedOrderID string) (o models.Order, err error) {
	err = r.db.WithContext(ctx).Where("generated_order_id = ?", generatedOrderID).First(&o).Error
	return
}

func (r *orderRepositoryImpl) Account(ctx context.Context, o *models.Order) (a models.Account, err error) {
	err = r.db.WithContext(ctx).Where("id = ?", o.AccountID).First(&a).Error
	return
}
