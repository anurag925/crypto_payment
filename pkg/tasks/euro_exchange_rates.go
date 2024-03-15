package tasks

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/libs/impl"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"

	"github.com/hibiken/asynq"
)

type euroExchangeRates struct {
}

func (u euroExchangeRates) ProcessTask(ctx context.Context, t *asynq.Task) error {
	rate, err := impl.NewCoinbaseLib().AllExchangeRates(ctx, "EUR")
	if err != nil {
		return err
	}
	// logger.Info(ctx, "euro exchange rate are", "rate", rate)
	for k, v := range rate {
		if err := app.Cache().Instance().HSet(ctx, services.EuroExchangeRate, k, v).Err(); err != nil {
			logger.Error(ctx, "error save euro exchange rate", "err", err)
		}
	}
	return nil
}
