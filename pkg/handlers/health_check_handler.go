package handlers

import (
	"context"
	"net/http"

	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/pkg/services/impl"
	"github.com/anurag925/crypto_payment/utils/logger"

	"github.com/labstack/echo/v4"
)

type healthCheckController struct {
	s services.HealthCheckService
}

func NewHealthCheckController() *healthCheckController {
	return &healthCheckController{impl.NewHealthCheckServiceImpl()}
}

// Ping api to check health of the service
func (h *healthCheckController) Ping(c echo.Context) error {
	if h.s.HealthCheck(Context(c)) {
		return c.JSON(http.StatusOK, "Pong")
	}
	return c.JSON(http.StatusOK, "")
}

// Ping api to check health of the service
func (h *healthCheckController) Hello(c echo.Context) error {
	logger.Info(Context(c), "hola amigos adios")
	return c.JSON(http.StatusOK, "world")
}

func Context(c echo.Context) context.Context {
	return c.Get("context").(context.Context)
}
