package impl

import (
	"context"

	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"
)

type healthCheckService struct {
}

var _ services.HealthCheckService = (*healthCheckService)(nil)

func NewHealthCheckServiceImpl() *healthCheckService {
	return &healthCheckService{}
}

func (s *healthCheckService) PrintConfigs(ctx context.Context) {
	logger.Debug(ctx, "printing configs", "configs", app.Config())
}

// HealthCheck
func (s *healthCheckService) HealthCheck(ctx context.Context) bool {
	if app.Config().UP {
		db, err := app.DB().Instance().DB()
		if err != nil {
			logger.Error(ctx, "error in db instance get", "error", err)
		}
		if err := db.PingContext(ctx); err != nil {
			logger.Error(ctx, "error in db ping", "error", err)
			return false
		}
		return true
	}
	logger.Info(ctx, "app not up", "up", app.Config().UP)
	return false
}
