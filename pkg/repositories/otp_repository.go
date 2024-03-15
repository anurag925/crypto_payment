package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type OtpRepository interface {
	FindById(context.Context, int64) (models.Otp, error)
	FindAll(ctx context.Context) ([]models.Otp, error)
	Create(context.Context, *models.Otp) error
	Save(context.Context, *models.Otp) error
	LastActive(ctx context.Context, r string, t models.OtpType, a models.OtpAction) (o models.Otp, err error)
}
