package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type otpRepositoryImpl struct {
	*RepositoryImpl[models.Otp]
}

var _ repositories.OtpRepository = (*otpRepositoryImpl)(nil)

func NewOtpRepositoryImpl(db *gorm.DB) *otpRepositoryImpl {
	return &otpRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Otp](db)}
}

func DefaultOtpRepositoryImpl() *otpRepositoryImpl {
	return NewOtpRepositoryImpl(app.DB().Instance())
}

func (r *otpRepositoryImpl) LastActive(ctx context.Context, receiver string, t models.OtpType, a models.OtpAction) (o models.Otp, err error) {
	err = r.db.WithContext(ctx).Where("receiver = ? and action = ? and type = ? and verified = false", receiver, a, t).Last(&o).Error
	return
}
