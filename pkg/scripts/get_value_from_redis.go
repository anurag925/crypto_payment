package scripts

import (
	"context"
	"fmt"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/services"
)

func GetValueFromRedis() {
	b, err := app.Cache().Instance().HGet(context.Background(), services.EuroExchangeRate, "USDT").Bytes()
	if err != nil {
		fmt.Printf("the error is %v", err)
		return
	}
	fmt.Printf("the value is %s", b)
}
