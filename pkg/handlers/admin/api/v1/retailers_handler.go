package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type retailerHandler struct {
	service services.RetailerService
}

func NewRetailerHandler(s services.RetailerService) *retailerHandler {
	return &retailerHandler{s}
}

func (h *retailerHandler) GetAllRetailers(c echo.Context) error {
	res, err := h.service.AllRetailers(handlers.Context(c))
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, res)
}

func (h *retailerHandler) GetRetailerDetails(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid id provided", err)
	}
	res, err := h.service.RetailersDetail(handlers.Context(c), id)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, res)
}

func (h *retailerHandler) PostRetailers(c echo.Context) error {
	var request services.RetailersCreateData
	if err := c.Bind(&request); err != nil {
		return handlers.BadRequestResponse(c, "invalid data for retailer creation", err)
	}
	response, err := h.service.CreateRetailersAdmin(handlers.Context(c), request)
	if err != nil {
		return err
	}
	return handlers.CreatedResponse(c, response)
}

func (h *retailerHandler) PostRetailersConfigurations(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("retailers_id"), 10, 64)
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid retailers id provided", err)
	}
	var request services.RetailersCreateData
	if err := c.Bind(&request); err != nil {
		return handlers.BadRequestResponse(c, "invalid data for retailer creation", err)
	}
	request.Retailer.ID = id
	response, err := h.service.CreateRetailersConfigurationsAdmin(handlers.Context(c), request)
	if err != nil {
		return err
	}
	return handlers.CreatedResponse(c, response)
}
