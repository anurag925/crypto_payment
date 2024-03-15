package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/app/core"
	"github.com/anurag925/crypto_payment/pkg/routes"
	"github.com/anurag925/crypto_payment/pkg/tasks"
	"github.com/anurag925/crypto_payment/utils/logger"
)

func main() {
	app.New(core.GetBackendApp())
	logger.Info(context.Background(), "App init done ...")
	routes.Init()
	logger.Info(context.Background(), "Router init done ...")

	tasks.RegisterTasks(context.Background())
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

	// Start server
	go func() {
		logger.Info(context.Background(), "Http Server starting ...")
		if err := app.Server().Instance().Start(fmt.Sprintf(":%d", app.Config().Port)); err != nil && err != http.ErrServerClosed {
			logger.Fatal(context.Background(), "shutting down the server because", "error", err)
		}
		logger.Info(context.Background(), "Http Server shutting down ...")
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info(context.Background(), "Shutting down the application ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	logger.Info(context.Background(), "Cancelling context ...")
	defer cancel()
	if err := Shutdown(ctx); err != nil {
		logger.Fatal(context.Background(), "error shutting down server", "error", err)
	}
	logger.Info(context.Background(), "Bye ...")
}

func Shutdown(ctx context.Context) error {
	if err := app.DB().Close(ctx); err != nil {
		return err
	}
	if err := app.Server().Close(ctx); err != nil {
		return err
	}
	if err := app.Cache().Close(ctx); err != nil {
		return err
	}
	if err := app.Worker().Close(ctx); err != nil {
		return err
	}
	return nil
}
