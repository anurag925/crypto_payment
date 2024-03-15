package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type accountRepositoryImpl struct {
	*RepositoryImpl[models.Account]
}

var _ repositories.AccountRepository = (*accountRepositoryImpl)(nil)

func NewAccountRepositoryImpl(db *gorm.DB) *accountRepositoryImpl {
	return &accountRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Account](db)}
}

func DefaultAccountRepositoryImpl() *accountRepositoryImpl {
	return NewAccountRepositoryImpl(app.DB().Instance())
}

func (r *accountRepositoryImpl) FindByEmail(ctx context.Context, email string) (o models.Account, err error) {
	err = r.db.WithContext(ctx).Where("email = ?", email).First(&o).Error
	return
}

func (r *accountRepositoryImpl) Create(ctx context.Context, o *models.Account) (err error) {
	if o.Password.Valid {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(o.Password.String), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		o.Password.SetValid(string(hashedPassword))
	}
	err = r.db.WithContext(ctx).Create(o).Error
	return
}

func (r *accountRepositoryImpl) Retailer(ctx context.Context, a models.Account) (o models.Retailer, err error) {
	err = r.db.WithContext(ctx).Where("account_id = ?", a.ID).First(&o).Error
	return
}

func (r *accountRepositoryImpl) GetAccountsByCreatedAt(ctx context.Context, start, end time.Time, page, limit int) (a []models.Account, err error) {
	err = r.db.WithContext(ctx).Where("created_at BETWEEN ? AND ?", start, end).Offset((page - 1) * limit).Limit(limit).Find(&a).Error
	return
}

func (r *accountRepositoryImpl) Kyc(ctx context.Context, o *models.Account) (k models.Kyc, err error) {
	err = r.db.Where("account_id = ?", o.ID).Find(&k).Error
	return
}

func (r *accountRepositoryImpl) ResidentialAddress(ctx context.Context, o *models.Account) (a models.Address, err error) {
	err = r.db.Model(o).Where("type = ?", int8(models.AddressTypeResidential)).Association("Address").Find(&a)
	return
}

func (r *accountRepositoryImpl) BusinessAddress(ctx context.Context, o *models.Account) (a models.Address, err error) {
	err = r.db.Model(o).Where("type = ?", int8(models.AddressTypeBusiness)).Association("Address").Find(&a)
	return
}
