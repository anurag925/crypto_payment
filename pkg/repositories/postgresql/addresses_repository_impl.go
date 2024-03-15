package postgresql

import (
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type addressRepositoryImpl struct {
	*RepositoryImpl[models.Address]
}

var _ repositories.AddressRepository = (*addressRepositoryImpl)(nil)

func NewAddressRepositoryImpl(db *gorm.DB) *addressRepositoryImpl {
	return &addressRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Address](db)}
}

func DefaultAddressRepositoryImpl() *addressRepositoryImpl {
	return NewAddressRepositoryImpl(app.DB().Instance())
}
