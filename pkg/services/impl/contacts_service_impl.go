package impl

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
)

type contactServiceImpl struct {
	contactRepo repositories.ContactRepository
}

func NewContactServiceImpl(contactRepo repositories.ContactRepository) *contactServiceImpl {
	return &contactServiceImpl{contactRepo: contactRepo}
}

func DefaultContactServiceImpl() *contactServiceImpl {
	return NewContactServiceImpl(postgresql.DefaultContactRepositoryImpl())
}

func (s *contactServiceImpl) Create(ctx context.Context, k *models.Contact) error {
	return s.contactRepo.Create(ctx, k)
}
