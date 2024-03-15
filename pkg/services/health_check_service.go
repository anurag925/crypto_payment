package services

import "context"

type HealthCheckService interface {
	PrintConfigs(context.Context)
	HealthCheck(context.Context) bool
}
