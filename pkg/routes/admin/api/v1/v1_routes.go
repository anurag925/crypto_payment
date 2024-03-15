package v1

import (
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/handlers"
	v1 "github.com/anurag925/crypto_payment/pkg/handlers/admin/api/v1"
	"github.com/anurag925/crypto_payment/pkg/services/impl"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	AccountID int64  `json:"account_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}

func Routes(api *echo.Group) {
	api_v1 := api.Group("/v1", echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte(app.Config().JwtSecret),
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(jwtCustomClaims)
			},
			ErrorHandler: func(c echo.Context, err error) error {
				return handlers.UnauthorizedResponse(c, "unauthorized account", err)
			},
		},
	), adminValidation, setAccount)

	transactionHandler := v1.NewTransactionHandler(impl.DefaultAdminTransactionServiceImpl())
	transactionRoutes := api_v1.Group("/transactions")
	transactionRoutes.GET("", transactionHandler.GetAllTransactions)
	transactionRoutes.GET("/:generated_id", transactionHandler.GetTransactionByGeneratedID)

	retailerHandler := v1.NewRetailerHandler(impl.DefaultRetailerServiceImpl())
	retailerRoutes := api_v1.Group("/retailers")
	retailerRoutes.GET("", retailerHandler.GetAllRetailers)
	retailerRoutes.POST("", retailerHandler.PostRetailers)
	retailerRoutes.POST("/:retailers_id/configurations", retailerHandler.PostRetailersConfigurations)
	retailerRoutes.GET("/:id", retailerHandler.GetRetailerDetails)

	customerHandler := v1.NewCustomerHandler(impl.DefaultCustomerServiceImpl())
	customerRoutes := api_v1.Group("/customers")
	customerRoutes.GET("", customerHandler.GetCustomersInTimePeriod)
	customerRoutes.GET("/:id", customerHandler.GetCustomerDetailById)

	accountHandler := v1.NewAccountHandler()
	accountRoutes := api_v1.Group("/accounts")
	accountRoutes.GET("/dashboard", accountHandler.Dashboard)

}
