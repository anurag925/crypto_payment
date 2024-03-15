package impl

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"
	"time"
)

type adminTransactionServiceImpl struct {
	accountRepo repositories.AccountRepository
	paymentRepo repositories.PaymentRepository
}

var _ services.AdminTransactionService = (*adminTransactionServiceImpl)(nil)

func NewAdminTransactionServiceImpl(paymentRepo repositories.PaymentRepository, accountRepo repositories.AccountRepository) *adminTransactionServiceImpl {
	return &adminTransactionServiceImpl{paymentRepo: paymentRepo, accountRepo: accountRepo}
}

func DefaultAdminTransactionServiceImpl() *adminTransactionServiceImpl {
	return NewAdminTransactionServiceImpl(postgresql.DefaultPaymentRepositoryImpl(), postgresql.DefaultAccountRepositoryImpl())
}

func (s *adminTransactionServiceImpl) GetAllAdminTransactions(ctx context.Context, start, end time.Time, page int) (services.AdminAllTransactions, error) {
	logger.Debug(ctx, "Getting all transactions")
	payments, err := s.paymentRepo.GetPaymentByCreatedAt(ctx, start, end, page, 1000)
	if err != nil {
		logger.Error(ctx, "error retrieving all transactions", "error", err)
		return services.AdminAllTransactions{}, err
	}
	logger.Info(ctx, "all transactions retrieved", "count", len(payments))
	adminTxns := make([]services.AdminTransaction, 0)
	for _, payment := range payments {
		adminTxn := services.AdminTransaction{
			ID:            payment.ID,
			GeneratedID:   payment.GeneratedID,
			DateAndTime:   payment.CreatedAt,
			Retailer:      payment.Order.RetailerAccount.Email,
			Customer:      payment.Order.Account.Email,
			PaymentMethod: payment.Mode.String(),
			FiatAmount:    payment.SettlementAmount,
			CryptoAmount:  payment.Order.CryptocurrencyAmount.String,
			PGFees:        payment.PGFees.String,
			RNFees:        payment.TxnFees.String,
			Status:        payment.Status.String(),
		}
		adminTxns = append(adminTxns, adminTxn)
	}
	logger.Info(ctx, "all adminTxns retrieved", "count", len(adminTxns))
	return services.AdminAllTransactions{
		Transactions: adminTxns,
	}, nil
}

func (s *adminTransactionServiceImpl) GetAdminTransactionByGeneratedId(ctx context.Context, generatedID string) (services.AdminDetailedTransaction, error) {
	logger.Debug(ctx, "Getting transaction by generated id")
	payment, err := s.paymentRepo.FindByGeneratedID(ctx, generatedID)
	if err != nil {
		return services.AdminDetailedTransaction{}, err
	}
	order := payment.Order
	customer_account := order.Account
	retailer_account := order.RetailerAccount
	retailer, err := s.accountRepo.Retailer(ctx, retailer_account)
	if err != nil {
		return services.AdminDetailedTransaction{}, err
	}
	return services.AdminDetailedTransaction{
		ID:                 payment.ID,
		GeneratedID:        payment.GeneratedID,
		RetailerName:       retailer.RetailName.String,
		RetailerID:         order.RetailerID,
		CustomerName:       customer_account.FirstName.String,
		CustomerID:         order.AccountID,
		RetailerOrderID:    order.ExternalOrderID,
		FiatAmount:         order.BaseCurrencyAmount.String,
		CryptoAmount:       order.CryptocurrencyAmount.String,
		DateAndTime:        payment.CreatedAt,
		CryptoExchangeRate: order.CryptoExchangeRate.String,
		PaymentMethod:      payment.Mode.String(),
		PGFees:             payment.PGFees.String,
		RNFees:             payment.TxnFees.String,
		Status:             payment.Status.String(),
		DestinationAddress: "hadhasbhabdshbsajsakndxknwq2skwnsklx",
	}, nil
}
