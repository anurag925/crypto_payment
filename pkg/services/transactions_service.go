package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
)

type Payout struct {
	Amount         string `json:"amount"`
	CryptoWalletID int64  `json:"crypto_wallet_id"`
}

type TransactionDetail struct {
	Transaction models.Transaction  `json:"transaction"`
	Wallet      models.CryptoWallet `json:"wallet"`
}

type TransactionService interface {
	AllTransactionsForRetailer(ctx context.Context, r models.Retailer) ([]models.Transaction, error)
	CreatePayout(ctx context.Context, r models.Retailer, p Payout) error
	TransactionsForRetailer(ctx context.Context, r models.Retailer, id int64) (TransactionDetail, error)
}
