package impl

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
)

type addressServiceImpl struct {
	addressRepo repositories.AddressRepository
}

func NewAddressServiceImpl(addressRepo repositories.AddressRepository) *addressServiceImpl {
	return &addressServiceImpl{addressRepo: addressRepo}
}

func DefaultAddressServiceImpl() *addressServiceImpl {
	return NewAddressServiceImpl(postgresql.DefaultAddressRepositoryImpl())
}

func (s *addressServiceImpl) Create(ctx context.Context, a *models.Address) error {
	return s.addressRepo.Create(ctx, a)
}
