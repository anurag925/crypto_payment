package services

import (
	"context"
	"time"
)

type AdminDetailedTransaction struct {
	ID                 int64     `json:"id"`
	GeneratedID        string    `json:"generated_id"`
	RetailerName       string    `json:"retailer_name"`
	RetailerID         int64     `json:"retailer_id"`
	CustomerName       string    `json:"customer_name"`
	CustomerID         int64     `json:"customer_id"`
	RetailerOrderID    string    `json:"retailer_order_id"`
	FiatAmount         string    `json:"fiat_amount"`
	CryptoAmount       string    `json:"crypto_amount"`
	DateAndTime        time.Time `json:"date_and_time"`
	CryptoExchangeRate string    `json:"crypto_exchange_rate"`
	PaymentMethod      string    `json:"payment_method"`
	PGFees             string    `json:"pg_fees"`
	RNFees             string    `json:"rn_fees"`
	Status             string    `json:"status"`
	DestinationAddress string    `json:"destination_address"`
}

type AdminTransaction struct {
	ID            int64     `json:"id"`
	GeneratedID   string    `json:"generated_id"`
	DateAndTime   time.Time `json:"date_and_time"`
	Retailer      string    `json:"retailer"`
	Customer      string    `json:"customer"`
	PaymentMethod string    `json:"payment_method"`
	FiatAmount    string    `json:"fiat_amount"`
	CryptoAmount  string    `json:"crypto_amount"`
	PGFees        string    `json:"pg_fees"`
	RNFees        string    `json:"rn_fees"`
	Status        string    `json:"status"`
}

type AdminAllTransactions struct {
	Transactions []AdminTransaction `json:"transactions"`
}

type AdminTransactionService interface {
	GetAllAdminTransactions(ctx context.Context, start, end time.Time, page int) (AdminAllTransactions, error)
	GetAdminTransactionByGeneratedId(ctx context.Context, generatedID string) (AdminDetailedTransaction, error)
}
