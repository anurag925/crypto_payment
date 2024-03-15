package postgresql

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"time"

	"gorm.io/gorm"
)

type retailerRepositoryImpl struct {
	*RepositoryImpl[models.Retailer]
}

var _ repositories.RetailerRepository = (*retailerRepositoryImpl)(nil)

func NewRetailerRepositoryImpl(db *gorm.DB) *retailerRepositoryImpl {
	return &retailerRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Retailer](db)}
}

func DefaultRetailerRepositoryImpl() *retailerRepositoryImpl {
	return NewRetailerRepositoryImpl(app.DB().Instance())
}

func (r *retailerRepositoryImpl) FindByAccountID(ctx context.Context, id int64) (o models.Retailer, err error) {
	err = r.db.WithContext(ctx).Where("account_id = ?", id).First(&o).Error
	return
}

func (r *retailerRepositoryImpl) Account(ctx context.Context, o *models.Retailer) (a models.Account, err error) {
	err = r.db.WithContext(ctx).Where("id = ?", o.AccountID).First(&a).Error
	return
}

func (r *retailerRepositoryImpl) Wallet(ctx context.Context, o *models.Retailer) (w models.Wallet, err error) {
	err = r.db.WithContext(ctx).Where("retailer_id = ?", o.ID).First(&w).Error
	return
}

func (r *retailerRepositoryImpl) TxnFeeConfigs(ctx context.Context, o *models.Retailer) (t []models.TxnFeeConfig, err error) {
	err = r.db.WithContext(ctx).Where("retailer_id = ?", o.ID).Find(&t).Error
	return
}

func (r *retailerRepositoryImpl) TxnFeeConfigsByPaymentMode(ctx context.Context, o *models.Retailer, mode models.PaymentMode) (t models.TxnFeeConfig, err error) {
	err = r.db.WithContext(ctx).Where("retailer_id = ? AND payment_mode = ?", o.ID, mode).First(&t).Error
	return
}

func (r *retailerRepositoryImpl) CryptoWallets(ctx context.Context, o *models.Retailer) (w []models.CryptoWallet, err error) {
	err = r.db.WithContext(ctx).Where("retailer_id = ?", o.ID).Find(&w).Error
	return
}

func (r *retailerRepositoryImpl) Contact(ctx context.Context, o *models.Retailer) (c models.Contact, err error) {
	err = r.db.WithContext(ctx).Where("point_of_contact = true").Where("retailer_id = ?", o.ID).Find(&c).Error
	return
}

func (r *retailerRepositoryImpl) Shareholders(ctx context.Context, o *models.Retailer) (c []models.Contact, err error) {
	err = r.db.WithContext(ctx).Where("ubo = true").Where("retailer_id = ?", o.ID).Find(&c).Error
	return
}

func (r *retailerRepositoryImpl) Documents(ctx context.Context, o *models.Retailer) (d []models.Document, err error) {
	kyc := models.Kyc{}
	err = r.db.WithContext(ctx).Where("account_id = ?", o.AccountID).Find(&kyc).Error
	if err != nil {
		return
	}
	err = r.db.WithContext(ctx).Where("kyc_id = ?", kyc.ID).Find(&d).Error
	return
}

func (r *retailerRepositoryImpl) Kyc(ctx context.Context, o *models.Retailer) (k models.Kyc, err error) {
	err = r.db.WithContext(ctx).Where("account_id = ?", o.AccountID).Find(&k).Error
	return
}

func (r *retailerRepositoryImpl) Address(ctx context.Context, o *models.Retailer) (a models.Address, err error) {
	err = r.db.WithContext(ctx).Where("type = ?", int8(models.AddressTypeBusiness)).
		Where("account_id = ?", o.AccountID).Find(&a).Error
	return
}

func (r *retailerRepositoryImpl) GetOrdersByCreatedAt(ctx context.Context, m *models.Retailer, start, end time.Time, page, limit int) (o []models.Order, err error) {
	err = r.db.WithContext(ctx).
		Where("retailer_account_id = ?", m.AccountID).
		Where("created_at BETWEEN ? AND ?", start, end).
		Offset((page - 1) * limit).
		Limit(limit).Find(&o).Error
	return
}

func (r *retailerRepositoryImpl) FindByAccount(ctx context.Context, a *models.Account) (o models.Retailer, err error) {
	err = r.db.WithContext(ctx).Where("account_id = ?", a.ID).First(&o).Error
	return
}
