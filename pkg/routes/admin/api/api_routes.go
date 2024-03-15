package api

import (
	v1 "github.com/anurag925/crypto_payment/pkg/routes/admin/api/v1"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	api := g.Group("/api")
	v1.Routes(api)
}
