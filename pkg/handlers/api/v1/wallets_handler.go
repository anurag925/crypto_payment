package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/services"

	"github.com/labstack/echo/v4"
)

type walletHandler struct {
	service services.WalletService
}

func NewWalletHandler(s services.WalletService) *walletHandler {
	return &walletHandler{s}
}

func (h *walletHandler) GetWalletForRetailer(c echo.Context) error {
	res, err := h.service.WalletForRetailer(handlers.Context(c), handlers.Retailer(c))
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, res)
}
