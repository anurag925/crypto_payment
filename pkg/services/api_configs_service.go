package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type ApiConfigService interface {
	FindByKey(ctx context.Context, apiKey string) (models.ApiConfig, error)
	ApiConfigsForRetailer(ctx context.Context, o models.Retailer) ([]models.ApiConfig, error)
	CreateApiConfig(ctx context.Context, o models.Retailer, apiConfig *models.ApiConfig) error
	GenerateChecksum(ctx context.Context, secret string, body string) (string, error)
	CompareChecksum(ctx context.Context, apiConfig models.ApiConfig, checksum string, body string) (bool, error)
	CheckOrderCreateChecksum(ctx context.Context, apiKey string, request OrderCreateRequest) (bool, models.Retailer, error)
}
