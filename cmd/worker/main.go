package main

import (
	"context"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/app/core"
	"github.com/anurag925/crypto_payment/pkg/tasks"
	"github.com/anurag925/crypto_payment/utils/logger"
)

func main() {
	app.New(core.GetBackendApp())
	logger.Info(context.Background(), "App init done ...")

	if err := tasks.RegisterTasks(context.Background()); err != nil {
		logger.Info(context.Background(), "Task Registration failed ...", "error", err)
	}
	logger.Info(context.Background(), "Task Registration init done ...")

	go func() {
		logger.Info(context.Background(), "Async schedular server starting ...")
		if err := app.Worker().Instance().StartScheduler(); err != nil {
			logger.Fatal(context.Background(), "shutting down the async schedular because", "error", err)
		}
		logger.Info(context.Background(), "Async schedular server shutting down ...")
	}()

	if err := tasks.ScheduleTasks(context.Background()); err != nil {
		logger.Info(context.Background(), "Task Scheduling failed ...", "error", err)
	}
	logger.Info(context.Background(), "Task Scheduling done ...")

	if err := app.Worker().Instance().Start(); err != nil {
		logger.Fatal(context.Background(), "unable to start worker", "error", err)
	}
}
