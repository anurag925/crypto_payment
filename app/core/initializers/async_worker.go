package initializers

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/app/configs"
	"github.com/anurag925/crypto_payment/utils/logger"
	"github.com/anurag925/crypto_payment/utils/task"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

type AsynqClient struct {
	c *task.Client
}

var _ app.AsyncWorker = (*AsynqClient)(nil)

func (w *AsynqClient) Init(ctx context.Context, c configs.Config, l logger.Logger) error {
	l.Info(ctx, "Async worker init ...")
	conn := redisClientOpt(c)
	opts := schedulerOpts(c, l)
	cfg := asynqConfig(c, l)
	w.c = task.New(
		task.WithClient(conn),
		task.WithSchedular(conn, opts),
		task.WithServer(conn, cfg),
	)
	l.Info(ctx, "Async worker init done ...")
	return nil
}
func (w *AsynqClient) Instance() *task.Client {
	return w.c
}
func (w *AsynqClient) Close(context.Context) error {
	return w.c.Stop()
}

func redisClientOpt(c configs.Config) asynq.RedisClientOpt {
	var tlsConfig *tls.Config
	if c.Env != configs.Development {
		tlsConfig = &tls.Config{}
	}
	return asynq.RedisClientOpt{
		Addr:      fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort),
		Password:  c.RedisPassword,
		DB:        c.WorkerRedisDB,
		PoolSize:  c.RedisPoolSize,
		TLSConfig: tlsConfig,
	}
}

func schedulerOpts(c configs.Config, l logger.Logger) *asynq.SchedulerOpts {
	logLevel := asynq.DebugLevel
	if c.Env == configs.Production {
		logLevel = asynq.InfoLevel
	}
	return &asynq.SchedulerOpts{
		Logger:   logger.ConvertLoggerToAsynq(l),
		LogLevel: logLevel,
		PreEnqueueFunc: func(t *asynq.Task, opts []asynq.Option) {
			logger.Info(context.Background(), "Task Start Time", "task_name", t.Type(), "time", time.Now())
		},
		PostEnqueueFunc: func(i *asynq.TaskInfo, err error) {
			logger.Info(context.Background(), "Task End Time", "task_name", i.Type, "time", time.Now())
			// some monitoring raise
		},
	}
}

func asynqConfig(c configs.Config, l logger.Logger) asynq.Config {
	logLevel := asynq.DebugLevel
	if c.Env == configs.Production {
		logLevel = asynq.InfoLevel
	}
	return asynq.Config{
		BaseContext: func() context.Context {
			return context.WithValue(
				context.Background(),
				logger.ContextKeyValues,
				logger.ContextValue{logger.ContextKeyRequestID: uuid.NewString()},
			)
		},
		Logger:   logger.ConvertLoggerToAsynq(l),
		LogLevel: logLevel,
	}
}
