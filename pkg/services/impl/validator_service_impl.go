package impl

import (
	"context"
	"github.com/anurag925/crypto_payment/utils/logger"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	v     *validator.Validate
	vOnce sync.Once
) // ValidatorService validator service

func Validator() *validator.Validate {
	vOnce.Do(func() {
		v = validator.New()
	})
	if v == nil {
		logger.Error(context.Background(), "validator is nil")
	}
	return v
}
