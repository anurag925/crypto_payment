package libs

import "context"

type CoinbaseLib interface {
	GetUSDTtoEUROExchangeRate(ctx context.Context) (float64, error)
}
