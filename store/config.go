package store

import "fmt"

type Config struct {
	dbname   string
	user     string
	password string
	host     string
	port     int
	sslMode  string
}

func (c *Config) connStr() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=%s",
		c.dbname, c.user, c.password, c.host, c.port, c.sslMode)
}

func NewConfig(dbname string, user string, password string, host string, port int, sslMode bool) *Config {
	ssl := "disable"
	if sslMode {
		ssl = "enable"
	}
	return &Config{
		dbname:   dbname,
		user:     user,
		password: password,
		host:     host,
		port:     port,
		sslMode:  ssl,
	}
}
