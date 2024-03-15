package interfaces

import (
	v1 "github.com/anurag925/crypto_payment/pkg/routes/api/v1"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	api := g.Group("/interface")
	v1.Routes(api)
}
