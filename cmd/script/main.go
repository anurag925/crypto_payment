package main

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/app/core"
	"github.com/anurag925/crypto_payment/pkg/scripts"
	"github.com/anurag925/crypto_payment/utils/logger"
)

func main() {
	app.New(core.GetBackendApp())
	logger.Info(context.Background(), "App init done ...")
	scripts.GetValueFromRedis()
}
