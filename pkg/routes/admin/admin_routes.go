package admin

import (
	"github.com/anurag925/crypto_payment/pkg/routes/admin/api"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	admin := g.Group("/admin")
	api.Routes(admin)
}
