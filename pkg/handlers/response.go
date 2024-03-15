package handlers

import (
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/app/configs"
	"github.com/anurag925/crypto_payment/utils/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpResponse struct {
	Status    string `json:"status"`
	Data      any    `json:"data,omitempty"`
	Message   string `json:"message,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	Internal  error  `json:"internal,omitempty"`
}

func Success(data any) HttpResponse {
	return HttpResponse{
		Status: "success",
		Data:   data,
	}
}

func Error(errorCode, message string, err error) HttpResponse {
	res := HttpResponse{
		Status:    "failure",
		ErrorCode: errorCode,
		Message:   message,
	}
	if app.Config().Env == configs.Development {
		res.Internal = err
	}
	return res
}

func SuccessResponse(c echo.Context, data any) error {
	return c.JSONPretty(http.StatusOK, Success(data), "")
}

func CreatedResponse(c echo.Context, data any) error {
	return c.JSONPretty(http.StatusCreated, Success(data), "")
}

func BadRequestResponse(c echo.Context, message string, err error) error {
	logger.Error(Context(c), "BAD REQUEST HANDLER", "error", err)
	return c.JSONPretty(http.StatusBadRequest, Error("C400", message, err), "")
}

func UnauthorizedResponse(c echo.Context, message string, err error) error {
	logger.Error(Context(c), "UNAUTHORIZED REQUEST HANDLER", "error", err)
	return c.JSONPretty(http.StatusUnauthorized, Error("C401", message, err), "")
}

func InternalServerErrorResponse(c echo.Context, message string, err error) error {
	logger.Error(Context(c), "INTERNAL SERVER ERROR HANDLER", "error", err)
	return c.JSONPretty(http.StatusInternalServerError, Error("C500", message, err), "")
}
