package db

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	defaultMaxIdleConns = 10
	defaultMaxOpenConns = 10
)

var db *sqlx.DB

type Config struct {
	Driver          string
	URL             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifeTime time.Duration
}

func (c *Config) maxIdleConns() int {
	if c.MaxIdleConns == 0 {
		return defaultMaxIdleConns
	}
	return c.MaxIdleConns
}
func (c *Config) maxOpenConns() int {
	if c.MaxOpenConns == 0 {
		return defaultMaxOpenConns
	}
	return c.MaxOpenConns
}

func Init(config *Config) error {
	d, err := NewDB(config)
	if err != nil {
		return err
	}
	db = d
	return nil
}

func NewDB(config *Config) (*sqlx.DB, error) {
	d, err := sqlx.Open(config.Driver, config.URL)
	if err != nil {
		return nil, err
	}

	if err = d.Ping(); err != nil {
		return nil, err
	}

	d.SetMaxIdleConns(config.maxIdleConns())
	d.SetMaxOpenConns(config.maxOpenConns())
	d.SetConnMaxLifetime(config.ConnMaxLifeTime)

	return d, err
}

func Close() error {
	return db.Close()
}

func Get() *sqlx.DB {
	return db
}

func WithTimeout(ctx context.Context, timeout time.Duration, op func(ctx context.Context) error) (err error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return op(ctxWithTimeout)
}
