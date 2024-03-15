package initializers

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/anurag925/crypto_payment/app/configs"
	"github.com/anurag925/crypto_payment/utils/logger"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	cache *redis.Client
}

func (s *RedisCache) Init(ctx context.Context, c configs.Config, l logger.Logger) error {
	l.Info(ctx, "Redis connection init ...")
	s.cache = redis.NewClient(redisOptions(c))
	if err := s.cache.Ping(ctx).Err(); err != nil {
		return err
	}
	l.Info(ctx, "Redis connection done ...")
	return nil
}

func (s *RedisCache) Instance() *redis.Client {
	return s.cache
}

func (s *RedisCache) Close(ctx context.Context) error {
	return s.cache.Close()
}

func redisOptions(c configs.Config) *redis.Options {
	var tlsConfig *tls.Config
	if c.Env != configs.Development {
		tlsConfig = &tls.Config{}
	}
	return &redis.Options{
		Addr:       fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort),
		Password:   c.RedisPassword,
		DB:         c.RedisDB,
		MaxRetries: c.RedisMaxRetry,
		PoolSize:   c.RedisPoolSize,
		TLSConfig:  tlsConfig,
	}
}
