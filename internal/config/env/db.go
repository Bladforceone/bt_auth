package env

import (
	"bt_auth/internal/config"
	"errors"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

type dbConfig struct {
	dsn string
}

func NewDBConfig() (config.DBConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("dsn not found")
	}

	return &dbConfig{
		dsn: dsn,
	}, nil
}

func (d *dbConfig) DSN() string {
	return d.dsn
}
