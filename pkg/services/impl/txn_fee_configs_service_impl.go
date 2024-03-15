package impl

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
)

type txnFeeConfigServiceImpl struct {
	txnFeeConfigRepo repositories.TxnFeeConfigRepository
}

func NewTxnFeeConfigServiceImpl(txnFeeConfigRepo repositories.TxnFeeConfigRepository) *txnFeeConfigServiceImpl {
	return &txnFeeConfigServiceImpl{txnFeeConfigRepo: txnFeeConfigRepo}
}

func DefaultTxnFeeConfigServiceImpl() *txnFeeConfigServiceImpl {
	return NewTxnFeeConfigServiceImpl(postgresql.DefaultTxnFeeConfigRepositoryImpl())
}

func (s *txnFeeConfigServiceImpl) Create(ctx context.Context, t *models.TxnFeeConfig) error {
	return s.txnFeeConfigRepo.Create(ctx, t)
}
