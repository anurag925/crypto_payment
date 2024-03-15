package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/libs"
	"github.com/anurag925/crypto_payment/pkg/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type paymentsHandler struct {
	service services.InternalPaymentService
}

func NewPaymentsHandler(service services.InternalPaymentService) *paymentsHandler {
	return &paymentsHandler{service: service}
}

func (h *paymentsHandler) Callback(c echo.Context) error {
	var request libs.CallbackRequest
	if err := c.Bind(&request); err != nil {
		return err
	}
	if err := h.service.Callback(handlers.Context(c), request); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, "payment creation successful")
}
