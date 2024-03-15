package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type AdminInfoRetailer struct {
	ID               int64  `json:"id"`
	RetailerName     string `json:"retailer_name"`
	Status           string `json:"status"`
	OTV              string `json:"otv"`
	MTV              string `json:"mtv"`
	PendingPayout    string `json:"pending_payout"`
	AvailableBalance string `json:"available_balance"`
}

type AdminInfoRetailers struct {
	Retailers []AdminInfoRetailer `json:"retailers"`
}

type RetailersDetail struct {
	Account       models.Account        `json:"account,omitempty"`
	Retailer      models.Retailer       `json:"retailer,omitempty"`
	Contact       models.Contact        `json:"contact,omitempty"`
	Address       models.Address        `json:"address,omitempty"`
	Kyc           models.Kyc            `json:"kyc,omitempty"`
	Shareholders  []models.Contact      `json:"shareholders,omitempty"`
	Documents     []models.Document     `json:"documents,omitempty"`
	CryptoWallets []models.CryptoWallet `json:"crypto_wallets,omitempty"`
	TxnFeeConfigs []models.TxnFeeConfig `json:"txn_fee_configs,omitempty"`
}

type RetailersCreateData struct {
	Account       models.Account        `json:"account"`
	Retailer      models.Retailer       `json:"retailer"`
	Address       models.Address        `json:"address"`
	Contacts      []models.Contact      `json:"contacts"`
	Documents     []models.Document     `json:"documents"`
	CryptoWallets []models.CryptoWallet `json:"crypto_wallets"`
	TxnFeeConfigs []models.TxnFeeConfig `json:"txn_fee_configs"`
}
type RetailerService interface {
	FindById(ctx context.Context, id int64) (models.Retailer, error)
	Create(ctx context.Context, r *models.Retailer) error
	Update(ctx context.Context, r *models.Retailer) error
	CreateRetailersAdmin(ctx context.Context, request RetailersCreateData) (RetailersDetail, error)
	CreateRetailersConfigurationsAdmin(ctx context.Context, request RetailersCreateData) (RetailersDetail, error)
	AllRetailers(ctx context.Context) (AdminInfoRetailers, error)
	RetailersDetail(ctx context.Context, retailerId int64) (RetailersDetail, error)
}
