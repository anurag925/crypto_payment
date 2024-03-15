package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"

	"github.com/labstack/echo/v4"
)

type otpHandler struct {
	service services.OtpService
}

func NewOtpHandler(service services.OtpService) *otpHandler {
	return &otpHandler{service: service}
}

func (h *otpHandler) SendOtp(c echo.Context) error {
	request := services.GenerateOtp{}
	if err := c.Bind(&request); err != nil {
		return handlers.BadRequestResponse(c, "invalid request data", err)
	}
	logger.Info(handlers.Context(c), "the request is ", "r", request)
	// if err := impl.Validator().Struct(&request); err != nil {
	// 	return handlers.BadRequestResponse(c, "data validation failed", err)
	// }
	if err := h.service.SendOtpMail(handlers.Context(c), request); err != nil {
		return err
	}
	return handlers.SuccessResponse(c, "otp sent successfully")
}

func (h *otpHandler) VerifyOtp(c echo.Context) error {
	request := services.VerifyOtp{}
	if err := c.Bind(&request); err != nil {
		return handlers.BadRequestResponse(c, "invalid request data", err)
	}
	logger.Info(handlers.Context(c), "the request is ", "r", request)
	// if err := impl.Validator().Struct(&request); err != nil {
	// 	return handlers.BadRequestResponse(c, "data validation failed", err)
	// }
	if err := h.service.Verify(handlers.Context(c), request); err != nil {
		return err
	}
	return handlers.SuccessResponse(c, "otp verified successfully")
}
