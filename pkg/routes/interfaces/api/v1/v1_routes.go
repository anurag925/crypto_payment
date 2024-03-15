package v1

import (
	v1 "github.com/anurag925/crypto_payment/pkg/handlers/interfaces/api/v1"
	"github.com/anurag925/crypto_payment/pkg/services/impl"

	"github.com/labstack/echo/v4"
)

func Routes(api *echo.Group) {
	v1Unauthorized := api.Group("v1")

	paymentHandler := v1.NewPaymentsHandler(impl.DefaultInternalPaymentServiceImpl())
	externalPaymentRoutes := v1Unauthorized.Group("/payments")
	externalPaymentRoutes.POST("/callback", paymentHandler.Callback)

}
