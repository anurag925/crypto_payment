package app

import (
	"context"
	"sync"

	"github.com/anurag925/crypto_payment/app/configs"
	"github.com/anurag925/crypto_payment/utils/logger"
	"github.com/anurag925/crypto_payment/utils/task"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App interface {
	Config() configs.Config
	Server() HttpServer
	DB() SQL
	Cache() CacheStore
	Worker() AsyncWorker
}

var (
	instance App
	once     sync.Once
)

func New(a App) {
	once.Do(func() { instance = a })
}

func Config() configs.Config {
	return instance.Config()
}

func Server() HttpServer {
	return instance.Server()
}

func DB() SQL {
	return instance.DB()
}

func Cache() CacheStore {
	return instance.Cache()
}

func Worker() AsyncWorker {
	return instance.Worker()
}

type HttpServer interface {
	Init(context.Context, configs.Config, logger.Logger) error
	Instance() *echo.Echo
	Close(context.Context) error
}

type SQL interface {
	Init(context.Context, configs.Config, logger.Logger) error
	Instance() *gorm.DB
	Close(context.Context) error
}

type CacheStore interface {
	Init(context.Context, configs.Config, logger.Logger) error
	Instance() *redis.Client
	Close(context.Context) error
}

type AsyncWorker interface {
	Init(context.Context, configs.Config, logger.Logger) error
	Instance() *task.Client
	Close(context.Context) error
}
