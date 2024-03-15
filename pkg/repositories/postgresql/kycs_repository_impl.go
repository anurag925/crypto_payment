package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type kycRepositoryImpl struct {
	*RepositoryImpl[models.Kyc]
}

var _ repositories.KycRepository = (*kycRepositoryImpl)(nil)

func NewKycRepositoryImpl(db *gorm.DB) *kycRepositoryImpl {
	return &kycRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Kyc](db)}
}

func DefaultKycRepositoryImpl() *kycRepositoryImpl {
	return NewKycRepositoryImpl(app.DB().Instance())
}

func (r *kycRepositoryImpl) Documents(ctx context.Context, o *models.Kyc) (d []models.Document, err error) {
	err = r.db.Where("kyc_id = ?", o.ID).Find(&d).Error
	return
}
