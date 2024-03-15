package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/services"

	"github.com/labstack/echo/v4"
)

type apiConfigHandler struct {
	service services.ApiConfigService
}

func NewApiConfigHandler(s services.ApiConfigService) *apiConfigHandler {
	return &apiConfigHandler{s}
}

func (h *apiConfigHandler) GetApiConfigsForRetailer(c echo.Context) error {
	apiConfigs, err := h.service.ApiConfigsForRetailer(handlers.Context(c), handlers.Retailer(c))
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, apiConfigs)
}

func (h *apiConfigHandler) CreateApiConfigForRetailer(c echo.Context) error {
	apiConfig := models.ApiConfig{}
	if err := c.Bind(&apiConfig); err != nil {
		return handlers.BadRequestResponse(c, "wrong request body", err)
	}
	err := h.service.CreateApiConfig(handlers.Context(c), handlers.Retailer(c), &apiConfig)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, apiConfig)
}
