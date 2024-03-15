package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(s services.TransactionService) *transactionHandler {
	return &transactionHandler{s}
}

func (h *transactionHandler) GetAllTransactionsForRetailer(c echo.Context) error {
	res, err := h.service.AllTransactionsForRetailer(handlers.Context(c), handlers.Retailer(c))
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, res)
}

func (h *transactionHandler) CreateRetailerPayout(c echo.Context) error {
	var req services.Payout
	if err := c.Bind(&req); err != nil {
		return handlers.BadRequestResponse(c, "invalid payout request body", err)
	}
	err := h.service.CreatePayout(handlers.Context(c), handlers.Retailer(c), req)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, "payout successfully created")
}

func (h *transactionHandler) GetRetailerTransactionDetails(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid transaction id", err)
	}
	res, err := h.service.TransactionsForRetailer(handlers.Context(c), handlers.Retailer(c), id)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, res)
}
