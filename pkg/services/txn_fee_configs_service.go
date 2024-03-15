package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type TxnFeeConfigService interface {
	Create(ctx context.Context, t *models.TxnFeeConfig) error
}
