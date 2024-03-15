package routes

import (
	"context"
	"encoding/json"
	"os"

	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/app/configs"
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/routes/admin"
	"github.com/anurag925/crypto_payment/pkg/routes/api"
	"github.com/anurag925/crypto_payment/utils/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	serverInstance := app.Server()
	baseRoutes(serverInstance)
	printRoutes(serverInstance)
}

// baseRoutes
func baseRoutes(s app.HttpServer) {
	server := s.Instance()
	healthCheckController := handlers.NewHealthCheckController()
	server.GET("/health_check", healthCheckController.Ping)
	server.GET("/hello", healthCheckController.Hello)

	api.Routes(s.Instance().Group(""))
	admin.Routes(s.Instance().Group(""))

	// add request id to the request and other request related data
	server.Use(middleware.RequestIDWithConfig(
		middleware.RequestIDConfig{
			RequestIDHandler: func(ctx echo.Context, s string) {
				ctx.Set("context", context.WithValue(
					ctx.Request().Context(), logger.ContextKeyValues,
					logger.ContextValue{
						logger.ContextKeyRequestID: s,
						logger.ContextKeyAccountID: ctx.Get("account_id"),
					}),
				)
			},
		},
	))
}

// printRoutes writes the routes to a file for debugging
func printRoutes(s app.HttpServer) error {
	if app.Config().Env != configs.Development {
		return nil
	}
	data, err := json.MarshalIndent(s.Instance().Routes(), "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile("routes.json", data, 0644); err != nil {
		return err
	}
	return nil
}
