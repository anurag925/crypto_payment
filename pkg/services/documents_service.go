package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type DocumentService interface {
	Create(ctx context.Context, d *models.Document) error
	Upload(ctx context.Context, path, access, contentType string) (string, error)
}
