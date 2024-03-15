package initializers

import (
	"context"
	"fmt"

	"github.com/anurag925/crypto_payment/app/configs"
	"github.com/anurag925/crypto_payment/utils/logger"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type PostgreSQL struct {
	db *gorm.DB
}

// InitPostgresDB initializes postgres db
func (p *PostgreSQL) Init(ctx context.Context, c configs.Config, l logger.Logger) error {
	l.Info(ctx, "DB connection init ...")
	// Replace the connection details with your own
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Kolkata",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort, c.DBSSLMode,
	)

	// Connect to the database
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: zapgorm2.New(l.Instance().(*zap.Logger)),
	})
	if err != nil {
		return err
	}
	if c.Env == configs.Development {
		conn.Debug()
	}
	p.db = conn
	l.Info(ctx, "DB connection completed ...")
	return nil
}

func (p *PostgreSQL) Instance() *gorm.DB {
	return p.db
}

func (p *PostgreSQL) Close(ctx context.Context) error {
	db, err := p.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
