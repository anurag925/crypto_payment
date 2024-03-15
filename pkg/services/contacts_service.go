package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type ContactService interface {
	Create(ctx context.Context, k *models.Contact) error
}
