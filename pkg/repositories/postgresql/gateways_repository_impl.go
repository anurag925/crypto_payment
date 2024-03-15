package postgresql

import (
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type gatewayRepositoryImpl struct {
	*RepositoryImpl[models.Gateway]
}

var _ repositories.GatewayRepository = (*gatewayRepositoryImpl)(nil)

func NewGatewayRepositoryImpl(db *gorm.DB) *gatewayRepositoryImpl {
	return &gatewayRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Gateway](db)}
}

func DefaultGatewayRepositoryImpl() *gatewayRepositoryImpl {
	return NewGatewayRepositoryImpl(app.DB().Instance())
}
