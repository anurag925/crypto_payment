package impl

import (
	"context"
	"fmt"
	"github.com/anurag925/crypto_payment/pkg/libs"
	"github.com/anurag925/crypto_payment/pkg/libs/impl"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"
	"strconv"

	"github.com/google/uuid"
)

type paymentServiceImpl struct {
	accountRepo repositories.AccountRepository
	paymentRepo repositories.PaymentRepository
	orderRepo   repositories.OrderRepository
	txnFeeRepo  repositories.TxnFeeConfigRepository
}

var _ services.PaymentService = (*paymentServiceImpl)(nil)

func NewPaymentServiceImpl(
	accountRepo repositories.AccountRepository,
	paymentRepo repositories.PaymentRepository,
	orderRepo repositories.OrderRepository,
	txnFeeRepo repositories.TxnFeeConfigRepository,
) *paymentServiceImpl {
	return &paymentServiceImpl{
		accountRepo: accountRepo,
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
		txnFeeRepo:  txnFeeRepo,
	}
}

func DefaultPaymentServiceImpl() *paymentServiceImpl {
	return NewPaymentServiceImpl(
		postgresql.DefaultAccountRepositoryImpl(),
		postgresql.DefaultPaymentRepositoryImpl(),
		postgresql.DefaultOrderRepositoryImpl(),
		postgresql.DefaultTxnFeeConfigRepositoryImpl(),
	)
}

func (s *paymentServiceImpl) Create(ctx context.Context, request services.PaymentCreateRequest) (services.PaymentCreateResponse, error) {
	response := services.PaymentCreateResponse{}
	order, err := s.orderRepo.FindByGeneratedOrderID(ctx, request.OrderId)
	if err != nil {
		return response, err
	}
	request.Payment.OrderID = order.ID
	request.Payment.GeneratedID = uuid.New().String()
	request.Payment.Amount = order.Amount
	settlementAmount, err := s.calculateSettlementAmount(ctx, order, &request.Payment)
	if err != nil {
		return response, err
	}
	request.Payment.SettlementAmount = settlementAmount
	if err := s.paymentRepo.Create(ctx, &request.Payment); err != nil {
		return response, err
	}
	account, err := s.orderRepo.Account(ctx, &order)
	if err != nil {
		return response, err
	}
	// call zen api for creating payment
	res, err := impl.NewZenPaymentLib().Create(ctx, libs.PaymentCreateRequest{
		Account:        account,
		Order:          order,
		Payment:        request.Payment,
		BrowserDetails: request.BrowserDetails,
		Signature:      request.Signature,
	})
	if err != nil {
		return response, err
	}
	return services.PaymentCreateResponse{Payment: request.Payment, ZenResponse: res}, nil
}

func (s *paymentServiceImpl) Status(ctx context.Context, paymentId int64) (models.Payment, error) {
	// payment status logic
	return s.paymentRepo.FindById(ctx, paymentId)
}

func (s *paymentServiceImpl) Transactions() (services.Transactions, error) {
	// fetch transactions
	return services.Transactions{}, nil
}

func (s *paymentServiceImpl) Callback() error {
	// handle callbacks
	return nil
}

func (s *paymentServiceImpl) PaymentsForAccount(ctx context.Context, email string) ([]models.Payment, error) {
	logger.Info(ctx, "fetching payments for account", "email", email, "accountRepo", s.accountRepo, "paymentRepo", s.paymentRepo, "orderRepo", s.orderRepo, "txnFeeRepo", s.txnFeeRepo)
	account, err := s.accountRepo.FindByEmail(ctx, email)
	if err != nil {
		logger.Error(ctx, "error while fetching account", "err", err)
		return nil, err
	}
	return s.paymentRepo.PaymentsForAccount(ctx, account)
}

func (s *paymentServiceImpl) calculateSettlementAmount(ctx context.Context, o models.Order, p *models.Payment) (string, error) {
	txnFees, err := s.txnFeeRepo.FindByRetailerIDAndPaymentMode(ctx, o.RetailerID, p.Mode)
	if err != nil {
		return "", err
	}
	floatAmount, err := strconv.ParseFloat(o.CryptocurrencyAmount.String, 64)
	if err != nil {
		return "", err
	}
	logger.Info(ctx, "crypto amount to be paid is", "floatAmount", floatAmount)
	totalTxnFees := convertTo2Decimal(txnFees.FixedFees.Float64 + ((txnFees.PercentFees.Float64 / 100) * floatAmount))
	totalPgTxnFees := convertTo2Decimal(txnFees.PgFixedFees.Float64 + ((txnFees.PgPercentageFees.Float64 / 100) * floatAmount))
	p.TxnFees.SetValid(fmt.Sprintf("%v", totalTxnFees))
	p.PGFees.SetValid(fmt.Sprintf("%v", totalPgTxnFees))
	logger.Info(ctx, "payment fees are", "totalTxnFees", totalTxnFees, "totalPgFees", totalPgTxnFees)
	settlementAmount := floatAmount - (totalTxnFees + totalPgTxnFees)
	p.SettlementAmount = fmt.Sprintf("%v", settlementAmount)
	return fmt.Sprintf("%v", settlementAmount), nil
}
