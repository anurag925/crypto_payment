package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/utils/logger"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Migrate(db *gorm.DB) {
	if err := db.Debug().AutoMigrate(
		&models.Account{},
		&models.Retailer{},
		&models.ApiConfig{},
		&models.Address{},
		&models.Order{},
		&models.Gateway{},
		&models.Payment{},
		&models.Kyc{},
		&models.Document{},
		&models.TxnFeeConfig{},
		&models.Contact{},
		&models.CryptoWallet{},
		&models.Wallet{},
		&models.Transaction{},
		&models.Otp{},
	); err != nil {
		logger.Error(context.Background(), "migration error", "error", err)
	}
}

type RepositoryImpl[T any] struct {
	db *gorm.DB
}

func NewRepositoryImpl[T any](db *gorm.DB) *RepositoryImpl[T] {
	return &RepositoryImpl[T]{db: db}
}

func (r *RepositoryImpl[T]) FindAll(ctx context.Context) (t []T, err error) {
	err = r.db.WithContext(ctx).Find(&t).Error
	return
}

func (r *RepositoryImpl[T]) FindById(ctx context.Context, id int64) (t T, err error) {
	err = r.db.WithContext(ctx).First(&t, id).Error
	return
}

func (r *RepositoryImpl[T]) PreloadFindById(ctx context.Context, id int64) (t T, err error) {
	err = r.db.WithContext(ctx).Preload(clause.Associations).First(&t, id).Error
	return
}

func (r *RepositoryImpl[T]) Create(ctx context.Context, o *T) (err error) {
	err = r.db.WithContext(ctx).Create(o).Error
	return
}

func (r *RepositoryImpl[T]) Save(ctx context.Context, o *T) (err error) {
	err = r.db.WithContext(ctx).Save(o).Error
	return
}
