package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"
	"strconv"

	"gopkg.in/guregu/null.v4"
)

type transactionServiceImpl struct {
	retailerRepo    repositories.RetailerRepository
	transactionRepo repositories.TransactionsRepository
	walletRepo      repositories.WalletRepository
}

var _ services.TransactionService = (*transactionServiceImpl)(nil)

func NewTransactionServiceImpl(
	transactionRepo repositories.TransactionsRepository,
	retailerRepo repositories.RetailerRepository,
	walletRepo repositories.WalletRepository,
) *transactionServiceImpl {
	return &transactionServiceImpl{
		transactionRepo: transactionRepo,
		retailerRepo:    retailerRepo,
		walletRepo:      walletRepo,
	}
}

func DefaultTransactionServiceImpl() *transactionServiceImpl {
	return NewTransactionServiceImpl(
		postgresql.DefaultTransactionRepositoryImpl(),
		postgresql.DefaultRetailerRepositoryImpl(),
		postgresql.DefaultWalletRepositoryImpl(),
	)
}

func (s *transactionServiceImpl) AllTransactionsForRetailer(ctx context.Context, r models.Retailer) ([]models.Transaction, error) {
	return s.transactionRepo.AllTransactionsForRetailer(ctx, r)
}

func (s *transactionServiceImpl) CreatePayout(ctx context.Context, r models.Retailer, p services.Payout) error {
	amount, err := strconv.ParseFloat(p.Amount, 64)
	if err != nil {
		logger.Error(ctx, "invalid amount", "error", err)
		return err
	}
	wallet, err := s.retailerRepo.Wallet(ctx, &r)
	if err != nil {
		logger.Error(ctx, "error getting wallet", "error", err)
		return err
	}
	balance, err := strconv.ParseFloat(wallet.Balance.String, 64)
	if err != nil {
		logger.Error(ctx, "invalid data in db for balance", "error", err)
		return err
	}
	pending, err := strconv.ParseFloat(wallet.Pending.String, 64)
	if err != nil {
		logger.Error(ctx, "invalid data in db for pending", "error", err)
		return err
	}
	if balance < amount {
		return errors.New("insufficient balance")
	}
	balance = balance - amount
	pending = pending + amount
	payout := models.Transaction{
		Status:         models.TransactionStatusCreated,
		Type:           models.TransactionTypeDebit,
		Amount:         null.StringFrom(p.Amount),
		Balance:        null.StringFrom(fmt.Sprintf("%v", balance)),
		Pending:        null.StringFrom(fmt.Sprintf("%v", pending)),
		CryptoWalletID: p.CryptoWalletID,
		RetailerID:     r.ID,
	}
	wallet.Balance.SetValid(fmt.Sprintf("%v", balance))
	wallet.Pending.SetValid(fmt.Sprintf("%v", pending))
	if err = s.walletRepo.Save(ctx, &wallet); err != nil {
		logger.Error(ctx, "wallet update failed", "error", err)
		return err
	}
	if err = s.transactionRepo.Create(ctx, &payout); err != nil {
		logger.Error(ctx, "transaction payout create failed", "error", err)
		return err
	}
	return nil
}

func (s *transactionServiceImpl) TransactionsForRetailer(ctx context.Context, r models.Retailer, id int64) (services.TransactionDetail, error) {
	transaction, err := s.transactionRepo.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, "retailer transaction get failed", "error", err)
		return services.TransactionDetail{}, err
	}
	if transaction.RetailerID != r.ID {
		logger.Error(ctx, "transaction not found for retailer")
		return services.TransactionDetail{}, errors.New("transaction not found for retailer")
	}
	wallet, err := s.transactionRepo.CryptoWallet(ctx, transaction)
	if err != nil {
		logger.Error(ctx, "retailer wallet get failed", "error", err)
		return services.TransactionDetail{}, err
	}
	return services.TransactionDetail{Wallet: wallet, Transaction: transaction}, nil
}
