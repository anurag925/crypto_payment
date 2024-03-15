package main

import (
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/app/core"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
)

func main() {
	app.New(core.GetBackendApp())
	postgresql.Migrate(app.DB().Instance())
}
