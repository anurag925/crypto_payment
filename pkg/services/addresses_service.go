package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type AddressService interface {
	Create(ctx context.Context, a *models.Address) error
}
