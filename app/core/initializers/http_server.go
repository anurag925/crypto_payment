package initializers

import (
	"context"

	"github.com/anurag925/crypto_payment/app/configs"
	"github.com/anurag925/crypto_payment/utils/logger"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

type EchoServer struct {
	e *echo.Echo
}

func (s *EchoServer) Init(ctx context.Context, c configs.Config, l logger.Logger) error {
	l.Info(ctx, "Starting init ...")
	server := echo.New()
	if c.Env == configs.Production {
		server.Logger.SetLevel(log.INFO)
		server.HideBanner = true
		server.HidePort = true
	} else {
		server.Debug = true
		server.Logger.SetLevel(log.DEBUG)
	}
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://dev.pay.crypto_payment.io, *"},
	}))
	server.Use(echozap.ZapLogger(l.Instance().(*zap.Logger)))
	server.Use(middleware.BodyDump(func(ctx echo.Context, b1, b2 []byte) {
		logger.Info(ctx.Get("context").(context.Context), "request", "method", ctx.Request().Method, "uri", ctx.Request().RequestURI, "body", b1)
		logger.Info(ctx.Get("context").(context.Context), "response", "method", ctx.Request().Method, "uri", ctx.Request().RequestURI, "body", b2)
	}))
	s.e = server
	l.Info(ctx, "Starting server ...")
	return nil
}

func (s *EchoServer) Instance() *echo.Echo {
	return s.e
}

func (s *EchoServer) Start() {}

func (s *EchoServer) Close(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
