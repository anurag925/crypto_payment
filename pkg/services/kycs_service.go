package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type KycService interface {
	Create(ctx context.Context, k *models.Kyc) error
}
