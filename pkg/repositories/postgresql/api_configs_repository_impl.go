package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type apiConfigRepositoryImpl struct {
	*RepositoryImpl[models.ApiConfig]
}

var _ repositories.ApiConfigRepository = (*apiConfigRepositoryImpl)(nil)

func NewApiConfigRepositoryImpl(db *gorm.DB) *apiConfigRepositoryImpl {
	return &apiConfigRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.ApiConfig](db)}
}

func DefaultApiConfigRepositoryImpl() *apiConfigRepositoryImpl {
	return NewApiConfigRepositoryImpl(app.DB().Instance())
}

func (r *apiConfigRepositoryImpl) FindByKey(ctx context.Context, key string) (a models.ApiConfig, err error) {
	err = r.db.Preload("Retailer").First(&a, "key = ?", key).Error
	return
}

func (r *apiConfigRepositoryImpl) FindByRetailer(ctx context.Context, o *models.Retailer) (a []models.ApiConfig, err error) {
	err = r.db.Where("retailer_id = ?", o.ID).Find(&a).Error
	return
}
