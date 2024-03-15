package repositories

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type TxnFeeConfigRepository interface {
	FindById(context.Context, int64) (models.TxnFeeConfig, error)
	FindAll(ctx context.Context) ([]models.TxnFeeConfig, error)
	Create(context.Context, *models.TxnFeeConfig) error
	FindByRetailerID(ctx context.Context, id int64) (t models.TxnFeeConfig, err error)
	FindByRetailerIDAndPaymentMode(ctx context.Context, id int64, mode models.PaymentMode) (t models.TxnFeeConfig, err error)
}
