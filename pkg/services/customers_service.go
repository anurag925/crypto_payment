package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"time"
)

type AdminCustomer struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	KycStage  string `json:"kyc_stage"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

type AdminAllCustomers struct {
	Customers []AdminCustomer `json:"customers"`
}

type CustomerDetail struct {
	Account   models.Account    `json:"account"`
	Kyc       models.Kyc        `json:"kyc"`
	Documents []models.Document `json:"documents"`
}

type CustomerService interface {
	GetCustomerDetailById(ctx context.Context, id int64) (CustomerDetail, error)
	AllCustomersInTimePeriod(ctx context.Context, startDate, endDate time.Time, page int) (AdminAllCustomers, error)
}
