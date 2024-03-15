package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func adminValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*jwtCustomClaims)
		role := claims.Role
		if role != "admin" {
			return handlers.UnauthorizedResponse(c, "account not a admin", nil)
		}
		return next(c)
	}
}

func setAccount(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*jwtCustomClaims)
		accountID := claims.AccountID
		account, err := postgresql.DefaultAccountRepositoryImpl().FindById(handlers.Context(c), accountID)
		if err != nil {
			return handlers.UnauthorizedResponse(c, "unable to find admin account", nil)
		}
		c.Set("account", account)
		return next(c)
	}
}
