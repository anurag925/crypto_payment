package v1

import (
	"net/http"
	"strconv"

	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/services"

	"github.com/labstack/echo/v4"
)

type paymentsHandler struct {
	service services.PaymentService
}

func NewPaymentsHandler(service services.PaymentService) *paymentsHandler {
	return &paymentsHandler{service: service}
}

func (h *paymentsHandler) PostPayment(c echo.Context) error {
	var request services.PaymentCreateRequest
	if err := c.Bind(&request); err != nil {
		return err
	}
	ip_address := c.RealIP()
	request.Payment.IPAddress = ip_address
	res, err := h.service.Create(handlers.Context(c), request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *paymentsHandler) GetStatus(c echo.Context) error {
	paymentId, err := strconv.ParseInt(c.Param("payment_id"), 10, 64)
	if err != nil {
		return handlers.BadRequestResponse(c, "bad payment id", err)
	}
	if res, err := h.service.Status(handlers.Context(c), paymentId); err != nil {
		// error sent at internal server error status 500
		return err
	} else {
		return c.JSON(http.StatusOK, res)
	}
}

func (c *paymentsHandler) GetAllTransactions(ctx echo.Context) error {
	if res, err := c.service.Transactions(); err != nil {
		// error sent at internal server error status 500
		return err
	} else {
		return ctx.JSON(http.StatusOK, res)
	}
}

func (c *paymentsHandler) GetTransactionsBy(ctx echo.Context) error {
	if res, err := c.service.Transactions(); err != nil {
		// error sent at internal server error status 500
		return err
	} else {
		return ctx.JSON(http.StatusOK, res)
	}
}

func (h *paymentsHandler) GetPaymentsForCustomer(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return handlers.BadRequestResponse(c, "invalid email id", nil)
	}
	payments, err := h.service.PaymentsForAccount(handlers.Context(c), email)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, payments)
}
