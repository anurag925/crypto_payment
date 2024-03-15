package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/services"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	service services.AdminTransactionService
}

func NewTransactionHandler(s services.AdminTransactionService) *transactionHandler {
	return &transactionHandler{s}
}

func (h *transactionHandler) GetAllTransactions(c echo.Context) error {
	startDate, err := time.Parse("02/01/2006", c.QueryParam("start_date"))
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid start date", err)
	}
	endDate, err := time.Parse("02/01/2006", c.QueryParam("end_date"))
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid end date", err)
	}
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid page number", err)
	}
	res, err := h.service.GetAllAdminTransactions(handlers.Context(c), startDate, endDate, int(page))
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, res)
}

func (h *transactionHandler) GetTransactionByGeneratedID(c echo.Context) error {
	id := c.Param("generated_id")
	res, err := h.service.GetAdminTransactionByGeneratedId(handlers.Context(c), id)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, res)
}
