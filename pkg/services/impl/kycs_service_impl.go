package impl

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
)

type kycServiceImpl struct {
	kycRepo repositories.KycRepository
}

func NewKycServiceImpl(kycRepo repositories.KycRepository) *kycServiceImpl {
	return &kycServiceImpl{kycRepo: kycRepo}
}

func DefaultKycServiceImpl() *kycServiceImpl {
	return NewKycServiceImpl(postgresql.DefaultKycRepositoryImpl())
}

func (s *kycServiceImpl) Create(ctx context.Context, k *models.Kyc) error {
	return s.kycRepo.Create(ctx, k)
}
