package store

import (
	"cost-calculator/config"
	"fmt"
)

type postgresConfig struct {
	dbname   string
	user     string
	password string
	host     string
	port     int
	sslMode  string
}

func (c *postgresConfig) connStr() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=%s",
		c.dbname, c.user, c.password, c.host, c.port, c.sslMode)
}

func newConfig(cfg *config.Configuration) *postgresConfig {
	return &postgresConfig{
		dbname:   cfg.PostgresDB,
		user:     cfg.PostgresUser,
		password: cfg.PostgresPassword,
		host:     cfg.PostgresHost,
		port:     cfg.PostgresPort,
		sslMode:  cfg.PostgresSSLMode,
	}
}
