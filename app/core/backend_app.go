package core

import (
	"context"
	"fmt"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/app/configs"
	"github.com/anurag925/crypto_payment/app/core/initializers"
	"github.com/anurag925/crypto_payment/utils/logger"
)

type backendApp struct {
	c  configs.Config
	s  *initializers.EchoServer
	db *initializers.PostgreSQL
	ca *initializers.RedisCache
	w  *initializers.AsynqClient
}

var _ app.App = (*backendApp)(nil)

func GetBackendApp() *backendApp {
	instance := &backendApp{
		c:  configs.Config{},
		s:  &initializers.EchoServer{},
		db: &initializers.PostgreSQL{},
		ca: &initializers.RedisCache{},
		w:  &initializers.AsynqClient{},
	}
	if err := instance.c.Load(); err != nil {
		panic(fmt.Sprintf("config load failed err: %+v", err))
	}
	ctx := context.Background()
	logger := logger.Init(logger.NewZapLogger(string(instance.c.Env)))
	logger.Info(ctx, "Logger and Config Initialized ...")
	logger.Debug(ctx, "Config Initialized ...", "configs", instance.c)
	logger.Info(ctx, "Starting Application ...")
	if err := instance.s.Init(ctx, instance.c, logger); err != nil {
		logger.Fatal(ctx, "server init failed err", "error", err)
	}
	if err := instance.db.Init(ctx, instance.c, logger); err != nil {
		logger.Fatal(ctx, "db init failed err", "error", err)
	}
	if err := instance.ca.Init(ctx, instance.c, logger); err != nil {
		logger.Fatal(ctx, "cache init failed err", "error", err)
	}
	if err := instance.w.Init(ctx, instance.c, logger); err != nil {
		logger.Fatal(ctx, "worker init failed err", "error", err)
	}
	logger.Info(ctx, "Basic App init done ...")
	return instance
}

func (a *backendApp) Config() configs.Config {
	return a.c
}

func (a *backendApp) Server() app.HttpServer {
	return a.s
}

func (a *backendApp) DB() app.SQL {
	return a.db
}

func (a *backendApp) Cache() app.CacheStore {
	return a.ca
}

func (a *backendApp) Worker() app.AsyncWorker {
	return a.w
}
