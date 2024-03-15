package tasks

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
)

func RegisterTasks(ctx context.Context) error {
	app.Worker().Instance().AddTask("update_euro_to_usdt_price", updateEuroToUsdtPrice{})
	app.Worker().Instance().AddTask("euro_exchange_rates", euroExchangeRates{})
	return nil
}

func ScheduleTasks(ctx context.Context) (err error) {
	err = app.Worker().Instance().New("euro_exchange_rates").Periodic("*/1 * * * *").Save()
	return err
}
