package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/services"

	"github.com/labstack/echo/v4"
)

type cryptoWalletHandler struct {
	service services.CryptoWalletService
}

func NewCryptoWalletHandler(s services.CryptoWalletService) *cryptoWalletHandler {
	return &cryptoWalletHandler{s}
}

func (h *cryptoWalletHandler) GetCryptoWalletsForRetailer(c echo.Context) error {
	cryptoWallets, err := h.service.CryptoWalletsForRetailer(handlers.Context(c), handlers.Retailer(c))
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, cryptoWallets)
}

func (h *cryptoWalletHandler) CreateCryptoWalletForRetailer(c echo.Context) error {
	cryptoWallet := models.CryptoWallet{}
	if err := c.Bind(&cryptoWallet); err != nil {
		return handlers.BadRequestResponse(c, "crypto wallet bad body", err)
	}
	err := h.service.CreateCryptoWallet(handlers.Context(c), handlers.Retailer(c), &cryptoWallet)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, cryptoWallet)
}

func (h *cryptoWalletHandler) UpdateCryptoWalletForRetailer(c echo.Context) error {
	cryptoWallet := models.CryptoWallet{}
	if err := c.Bind(&cryptoWallet); err != nil {
		return handlers.BadRequestResponse(c, "crypto wallet bad body", err)
	}
	err := h.service.CreateCryptoWallet(handlers.Context(c), handlers.Retailer(c), &cryptoWallet)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, cryptoWallet)
}
