package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/pkg/services/impl"
	"strconv"

	"github.com/labstack/echo/v4"
)

type accountHandler struct {
	service services.AccountService
}

func NewAccountHandler(s services.AccountService) *accountHandler {
	return &accountHandler{s}
}

func (h *accountHandler) PostAccount(c echo.Context) error {
	account := models.Account{}
	if err := c.Bind(&account); err != nil {
		return handlers.BadRequestResponse(c, "invalid request data", err)
	}
	if err := impl.Validator().Struct(&account); err != nil {
		return handlers.BadRequestResponse(c, "data validation failed", err)
	}
	err := h.service.Create(handlers.Context(c), &account)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, account)
}

func (h *accountHandler) GetAccount(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid id provided", err)
	}
	res, err := h.service.GetAccount(handlers.Context(c), id)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, res)
}

func (h *accountHandler) Login(c echo.Context) error {
	req := services.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return handlers.BadRequestResponse(c, "invalid request data", err)
	}
	if err := impl.Validator().Struct(&req); err != nil {
		return handlers.BadRequestResponse(c, "data validation failed", err)
	}
	res, err := h.service.Login(handlers.Context(c), req)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, res)
}
