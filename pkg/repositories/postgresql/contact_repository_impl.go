package postgresql

import (
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type contactRepositoryImpl struct {
	*RepositoryImpl[models.Contact]
}

var _ repositories.ContactRepository = (*contactRepositoryImpl)(nil)

func NewContactRepositoryImpl(db *gorm.DB) *contactRepositoryImpl {
	return &contactRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Contact](db)}
}

func DefaultContactRepositoryImpl() *contactRepositoryImpl {
	return NewContactRepositoryImpl(app.DB().Instance())
}
