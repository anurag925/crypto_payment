package tasks

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/libs/impl"
	"github.com/anurag925/crypto_payment/pkg/services"
	"time"

	"github.com/hibiken/asynq"
)

type updateEuroToUsdtPrice struct {
}

func (u updateEuroToUsdtPrice) ProcessTask(ctx context.Context, t *asynq.Task) error {
	rate, err := impl.NewCoinbaseLib().GetUSDTtoEUROExchangeRate(ctx, "EUR")
	if err != nil {
		return err
	}
	app.Cache().Instance().Set(ctx, services.EuroToUsdtExchangeRate, rate, 1*time.Hour)
	return nil
}
