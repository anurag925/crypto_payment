package postgresql

import (
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"

	"gorm.io/gorm"
)

type documentRepositoryImpl struct {
	*RepositoryImpl[models.Document]
}

var _ repositories.DocumentRepository = (*documentRepositoryImpl)(nil)

func NewDocumentRepositoryImpl(db *gorm.DB) *documentRepositoryImpl {
	return &documentRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Document](db)}
}

func DefaultDocumentRepositoryImpl() *documentRepositoryImpl {
	return NewDocumentRepositoryImpl(app.DB().Instance())
}
