package impl

import (
	"context"
	"fmt"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"
	"time"
)

type customerServiceImpl struct {
	accountRepo repositories.AccountRepository
	kycRepo     repositories.KycRepository
}

var _ services.CustomerService = (*customerServiceImpl)(nil)

func NewCustomerServiceImpl(accountRepo repositories.AccountRepository, kycRepo repositories.KycRepository) *customerServiceImpl {
	return &customerServiceImpl{accountRepo: accountRepo, kycRepo: kycRepo}
}

func DefaultCustomerServiceImpl() *customerServiceImpl {
	return NewCustomerServiceImpl(postgresql.DefaultAccountRepositoryImpl(), postgresql.DefaultKycRepositoryImpl())
}

func (s *customerServiceImpl) AllCustomersInTimePeriod(ctx context.Context, startDate, endDate time.Time, page int) (services.AdminAllCustomers, error) {
	accounts, err := s.accountRepo.GetAccountsByCreatedAt(ctx, startDate, endDate, page, 1000)
	if err != nil {
		logger.Error(ctx, "accounts fetch failed", "error", err)
		return services.AdminAllCustomers{}, err
	}
	adminCustomers := make([]services.AdminCustomer, 0)
	for _, account := range accounts {
		kyc, err := s.accountRepo.Kyc(ctx, &account)
		if err != nil {
			logger.Error(ctx, "accounts kyc fetch failed", "error", err)
			return services.AdminAllCustomers{}, err
		}
		adminCustomer := services.AdminCustomer{
			ID:        account.ID,
			Name:      fmt.Sprintf("%s %s", account.FirstName.String, account.LastName.String),
			Email:     account.Email,
			KycStage:  kyc.Status.String(),
			CreatedAt: account.CreatedAt.String(),
			Status:    account.Status.String(),
		}
		adminCustomers = append(adminCustomers, adminCustomer)
	}
	return services.AdminAllCustomers{Customers: adminCustomers}, nil
}

func (s *customerServiceImpl) GetCustomerDetailById(ctx context.Context, id int64) (services.CustomerDetail, error) {
	customer, err := s.accountRepo.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, "accounts fetch failed", "error", err)
		return services.CustomerDetail{}, err
	}
	kyc, err := s.accountRepo.Kyc(ctx, &customer)
	if err != nil {
		logger.Error(ctx, "accounts kyc fetch failed", "error", err)
		return services.CustomerDetail{}, err
	}
	documents, err := s.kycRepo.Documents(ctx, &kyc)
	if err != nil {
		logger.Error(ctx, "kyc document fetch failed", "error", err)
		return services.CustomerDetail{}, err
	}
	return services.CustomerDetail{
		Account:   customer,
		Kyc:       kyc,
		Documents: documents,
	}, err
}
