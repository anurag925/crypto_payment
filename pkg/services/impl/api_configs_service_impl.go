package impl

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type apiConfigServiceImpl struct {
	apiConfigRepo repositories.ApiConfigRepository
}

var _ services.ApiConfigService = (*apiConfigServiceImpl)(nil)

func NewApiConfigServiceImpl(apiConfigRepo repositories.ApiConfigRepository) *apiConfigServiceImpl {
	return &apiConfigServiceImpl{apiConfigRepo: apiConfigRepo}
}

func DefaultApiConfigServiceImpl() *apiConfigServiceImpl {
	return NewApiConfigServiceImpl(postgresql.DefaultApiConfigRepositoryImpl())
}

func (s *apiConfigServiceImpl) FindByKey(ctx context.Context, apiKey string) (models.ApiConfig, error) {
	return s.apiConfigRepo.FindByKey(ctx, apiKey)
}

func (s *apiConfigServiceImpl) GenerateChecksum(ctx context.Context, secret string, body string) (string, error) {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(body))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (s *apiConfigServiceImpl) CompareChecksum(ctx context.Context, apiConfig models.ApiConfig, checksum string, body string) (bool, error) {
	generatedChecksum, err := s.GenerateChecksum(ctx, apiConfig.Secret, body)
	if err != nil {
		logger.Error(ctx, "unexpected error", "error", err)
		return false, err
	}
	logger.Debug(ctx, "checksum", "generated", generatedChecksum)
	logger.Debug(ctx, "checksum", "received", checksum)
	if strings.Compare(checksum, generatedChecksum) == 0 {
		return true, nil
	}
	return false, nil
}

func (s *apiConfigServiceImpl) CheckOrderCreateChecksum(ctx context.Context, apiKey string, request services.OrderCreateRequest) (bool, models.Retailer, error) {
	apiConfig, err := s.FindByKey(ctx, apiKey)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, models.Retailer{}, errors.New("invalid api key")
		}
		return false, models.Retailer{}, err
	}
	if request.Internal {
		return true, apiConfig.Retailer, nil
	}
	checksumString := fmt.Sprintf(
		"%s%s%s%s%s",
		request.OrderDetail.CustomerID,
		request.OrderDetail.MerchantTransactionID,
		request.OrderDetail.Email,
		request.OrderDetail.Amount,
		request.OrderDetail.Currency,
	)
	ok, err := s.CompareChecksum(ctx, apiConfig, request.Checksum, checksumString)
	if err != nil {
		return false, models.Retailer{}, err
	}
	if ok {
		return true, apiConfig.Retailer, nil
	}
	return false, models.Retailer{}, nil
}

func (s *apiConfigServiceImpl) ApiConfigsForRetailer(ctx context.Context, o models.Retailer) ([]models.ApiConfig, error) {
	return s.apiConfigRepo.FindByRetailer(ctx, &o)
}

func (s *apiConfigServiceImpl) CreateApiConfig(ctx context.Context, o models.Retailer, apiConfig *models.ApiConfig) error {
	apiConfig.RetailerID = o.ID
	apiConfig.Key = uuid.NewString()
	apiConfig.Secret = uuid.NewString()
	return s.apiConfigRepo.Create(ctx, apiConfig)
}
