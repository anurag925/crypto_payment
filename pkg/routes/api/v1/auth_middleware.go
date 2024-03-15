package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func retailerValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*jwtCustomClaims)
		role := claims.Role
		if role != "retailer" {
			return handlers.UnauthorizedResponse(c, "account not a retailer", nil)
		}
		return next(c)
	}
}

func setRetailer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*jwtCustomClaims)
		accountID := claims.AccountID
		retailer, err := postgresql.DefaultRetailerRepositoryImpl().FindByAccountID(handlers.Context(c), accountID)
		if err != nil {
			return handlers.UnauthorizedResponse(c, "unable to find retailer", nil)
		}
		c.Set("retailer", retailer)
		return next(c)
	}
}
