package handlers

import (
	"github.com/anurag925/crypto_payment/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Retailer(c echo.Context) models.Retailer {
	return c.Get("retailer").(models.Retailer)
}

func Account(c echo.Context) models.Account {
	return c.Get("account").(models.Account)
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
