package v1

import (
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/handlers"
	v1 "github.com/anurag925/crypto_payment/pkg/handlers/api/v1"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
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
	v1_authorized := api.Group("/v1", echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte(app.Config().JwtSecret),
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(jwtCustomClaims)
			},
			ErrorHandler: func(c echo.Context, err error) error {
				return handlers.UnauthorizedResponse(c, "unauthorized account", err)
			},
		},
	), retailerValidation)

	v1_unauthorized := api.Group("/v1")

	accountHandler := v1.NewAccountHandler(impl.NewAccountService(postgresql.DefaultAccountRepositoryImpl()))
	accountRoutes := v1_authorized.Group("/accounts")
	v1_unauthorized.Group("/accounts").POST("/register", accountHandler.PostAccount)
	v1_unauthorized.Group("/accounts").POST("/login", accountHandler.Login)
	accountRoutes.GET("/:id", accountHandler.GetAccount)

	otpHandler := v1.NewOtpHandler(impl.NewOtpServiceImpl(postgresql.DefaultOtpRepositoryImpl()))
	otpRoutes := v1_unauthorized.Group("/otp")
	otpRoutes.POST("/send", otpHandler.SendOtp)
	otpRoutes.POST("/verify", otpHandler.VerifyOtp)

	orderHandler := v1.NewOrderHandler(impl.DefaultOrderServiceImpl())
	orderRoutes := v1_unauthorized.Group("/orders")
	orderRoutes.POST("", orderHandler.PostOrder)
	orderRoutes.GET("/settlement_amount", orderHandler.GetSettlementAmount)

	paymentHandler := v1.NewPaymentsHandler(impl.DefaultPaymentServiceImpl())
	externalPaymentRoutes := v1_unauthorized.Group("/payments")
	externalPaymentRoutes.POST("", paymentHandler.PostPayment)
	externalPaymentRoutes.GET("/:payment_id/status", paymentHandler.GetStatus)
	externalPaymentRoutes.GET("/list", paymentHandler.GetPaymentsForCustomer)

	documentHandler := v1.NewDocumentHandler(impl.DefaultDocumentServiceImpl())
	documentRoutes := v1_authorized.Group("/documents")
	documentRoutes.POST("", documentHandler.CreateDocument)
	documentRoutes.GET("/upload", documentHandler.UploadDocument)

	// start of retailer routes
	retailerRoutes := v1_authorized.Group("/retailers/:id", setRetailer)

	walletHandler := v1.NewWalletHandler(impl.DefaultWalletServiceImpl())
	retailerWalletRoutes := retailerRoutes.Group("/wallets")
	retailerWalletRoutes.GET("", walletHandler.GetWalletForRetailer)

	transactionHandler := v1.NewTransactionHandler(impl.DefaultTransactionServiceImpl())
	retailerTransactionRoutes := retailerRoutes.Group("/transactions")
	retailerTransactionRoutes.GET("", transactionHandler.GetAllTransactionsForRetailer)
	retailerTransactionRoutes.GET("/:transaction_id", transactionHandler.GetRetailerTransactionDetails)
	retailerTransactionRoutes.POST("/payout", transactionHandler.CreateRetailerPayout)

	retailerOrderRoutes := retailerRoutes.Group("/orders")
	retailerOrderRoutes.GET("", orderHandler.GetOrdersForRetailer)
	retailerOrderRoutes.GET("/:order_id", orderHandler.GetOrderDetailForRetailer)

	apiConfigHandler := v1.NewApiConfigHandler(impl.DefaultApiConfigServiceImpl())
	retailerApiConfigRoutes := retailerRoutes.Group("/api-configs")
	retailerApiConfigRoutes.GET("", apiConfigHandler.GetApiConfigsForRetailer)
	retailerApiConfigRoutes.POST("", apiConfigHandler.CreateApiConfigForRetailer)

	cryptoWalletHander := v1.NewCryptoWalletHandler(impl.DefaultCryptoWalletServiceImpl())
	retailerCryptoWalletRoutes := retailerRoutes.Group("/crypto-wallets")
	retailerCryptoWalletRoutes.GET("", cryptoWalletHander.GetCryptoWalletsForRetailer)
	retailerCryptoWalletRoutes.POST("", cryptoWalletHander.CreateCryptoWalletForRetailer)
	retailerCryptoWalletRoutes.PUT("", cryptoWalletHander.UpdateCryptoWalletForRetailer)

	retailerDocumentRoutes := retailerRoutes.Group("/documents")
	retailerDocumentRoutes.POST("", documentHandler.CreateDocument)
}
