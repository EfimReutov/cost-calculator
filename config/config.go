package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	EnvMain Env = "main"
	EnvDev  Env = "dev"
)

type Env string

type Configuration struct {
	Env              Env
	ServiceHost      string
	ServicePort      int
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     int
	PostgresSSLMode  string
	SMTPServer       string
	SMTPPort         int
	MailUser         string
	MailPassword     string
}

func LoadCfg() (*Configuration, error) {
	env := os.Getenv("ENV")
	if env == "" {
		return nil, errors.New("$ENV is empty")
	}
	serviceHost := os.Getenv("SERVICE_HOST")
	if env == "" {
		return nil, errors.New("SERVICE_HOST is empty")
	}
	servicePort := os.Getenv("SERVICE_PORT")
	if servicePort == "" {
		return nil, errors.New("SERVICE_PORT is empty")
	}
	servicePortInt, err := strconv.Atoi(servicePort)
	if err != nil {
		return nil, err
	}
	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresDB == "" {
		return nil, errors.New("$POSTGRES_DB is empty0")
	}
	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		return nil, errors.New("POSTGRES_USER is empty")
	}
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		return nil, errors.New("POSTGRES_PASSWORD is empty")
	}
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		return nil, errors.New("POSTGRES_HOST is empty")
	}
	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort == "" {
		return nil, errors.New("POSTGRES_PORT is empty")
	}
	postgresPortInt, err := strconv.Atoi(postgresPort)
	if err != nil {
		return nil, err
	}
	postgresSSLMode := os.Getenv("POSTGRES_SSL_MODE")
	if postgresSSLMode == "" {
		return nil, errors.New("POSTGRES_SSL_MODE is empty")
	}
	smtpServer := os.Getenv("SMTP_SERVER")
	if smtpServer == "" {
		return nil, errors.New("SMTP_SERVER is empty")
	}
	smtpPort := os.Getenv("SMTP_PORT")
	smtpPortInt, err := strconv.Atoi(smtpPort)
	if err != nil {
		return nil, err
	}
	if smtpPort == "" {
		return nil, errors.New("SMTP_PORT is empty")
	}
	mailUser := os.Getenv("MAIL_USER")
	if mailUser == "" {
		return nil, errors.New("MAIL_USER is empty")
	}
	mailPassword := os.Getenv("MAIL_PASSWORD")
	if mailPassword == "" {
		return nil, errors.New("MAIL_PASSWORD is empty")
	}
	return &Configuration{
		Env:              Env(env),
		ServiceHost:      serviceHost,
		ServicePort:      servicePortInt,
		PostgresDB:       postgresDB,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresHost:     postgresHost,
		PostgresPort:     postgresPortInt,
		PostgresSSLMode:  postgresSSLMode,
		SMTPServer:       smtpServer,
		SMTPPort:         smtpPortInt,
		MailUser:         mailUser,
		MailPassword:     mailPassword,
	}, nil
}
