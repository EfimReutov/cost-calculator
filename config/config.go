package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	EnvMain Env = "main"
	EnvDev  Env = "dev"

	ServerTypeREST = "REST"
	ServerTypeGRPC = "GRPC"
)

// Env represents current environment.
type Env string

// Configuration represents all required configurations.
type Configuration struct {
	Env              Env
	ServiceHost      string
	ServicePort      int
	ServerType       string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     int
	PostgresSSLMode  string
	RedisDB          int
	RedisPassword    string
	RedisHost        string
	RedisPort        int
	RedisPoolSize    int
	SMTPServer       string
	SMTPPort         int
	MailUser         string
	MailPassword     string
}

// LoadCfg read environments and return *Configuration.
func LoadCfg() (*Configuration, error) {
	env := os.Getenv("ENV")
	if env == "" {
		return nil, errors.New("the required $ENV environment variable is missing")
	}
	serviceHost := os.Getenv("SERVICE_HOST")
	if env == "" {
		return nil, errors.New("the required $SERVICE_HOST environment variable is missing")
	}
	servicePort := os.Getenv("SERVICE_PORT")
	if servicePort == "" {
		return nil, errors.New("the required $SERVICE_PORT environment variable is missing")
	}
	servicePortInt, err := strconv.Atoi(servicePort)
	if err != nil {
		return nil, err
	}
	serverType := os.Getenv("SERVER_TYPE")
	if serverType == "" {
		return nil, errors.New("the required $SERVER_TYPE environment variable is missing")
	}
	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresDB == "" {
		return nil, errors.New("$the required $POSTGRES_DB environment variable is missing")
	}
	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		return nil, errors.New("the required $POSTGRES_USER environment variable is missing")
	}
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		return nil, errors.New("the required $POSTGRES_PASSWORD environment variable is missing")
	}
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		return nil, errors.New("the required $POSTGRES_HOST environment variable is missing")
	}
	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort == "" {
		return nil, errors.New("the required $POSTGRES_PORT environment variable is missing")
	}
	postgresPortInt, err := strconv.Atoi(postgresPort)
	if err != nil {
		return nil, err
	}
	postgresSSLMode := os.Getenv("POSTGRES_SSL_MODE")
	if postgresSSLMode == "" {
		return nil, errors.New("the required $POSTGRES_SSL_MODE environment variable is missing")
	}
	redisDB := os.Getenv("REDIS_DB")
	if redisDB == "" {
		return nil, errors.New("the required $REDIS_DB environment variable is missing")
	}
	redisDBInt, err := strconv.Atoi(redisDB)
	if err != nil {
		return nil, err
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		return nil, errors.New("the required $REDIS_PASSWORD environment variable is missing")
	}
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		return nil, errors.New("the required $REDIS_HOST environment variable is missing")
	}
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		return nil, errors.New("the required $REDIS_PORT environment variable is missing")
	}
	redisPortInt, err := strconv.Atoi(redisPort)
	if err != nil {
		return nil, err
	}
	redisPoolSize := os.Getenv("REDIS_POOL_SIZE")
	if redisPoolSize == "" {
		return nil, errors.New("the required $REDIS_POOL_SIZE environment variable is missing")
	}
	redisPoolSizeInt, err := strconv.Atoi(redisPoolSize)
	if err != nil {
		return nil, err
	}
	smtpServer := os.Getenv("SMTP_SERVER")
	if smtpServer == "" {
		return nil, errors.New("the required $SMTP_SERVER environment variable is missing")
	}
	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		return nil, errors.New("the required $SMTP_PORT environment variable is missing")
	}
	smtpPortInt, err := strconv.Atoi(smtpPort)
	if err != nil {
		return nil, err
	}
	mailUser := os.Getenv("MAIL_USER")
	if mailUser == "" {
		return nil, errors.New("the required $MAIL_USER environment variable is missing")
	}
	mailPassword := os.Getenv("MAIL_PASSWORD")
	if mailPassword == "" {
		return nil, errors.New("the required $MAIL_PASSWORD environment variable is missing")
	}
	return &Configuration{
		Env:              Env(env),
		ServiceHost:      serviceHost,
		ServicePort:      servicePortInt,
		ServerType:       serverType,
		PostgresDB:       postgresDB,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresHost:     postgresHost,
		PostgresPort:     postgresPortInt,
		PostgresSSLMode:  postgresSSLMode,
		RedisDB:          redisDBInt,
		RedisHost:        redisHost,
		RedisPort:        redisPortInt,
		RedisPassword:    redisPassword,
		RedisPoolSize:    redisPoolSizeInt,
		SMTPServer:       smtpServer,
		SMTPPort:         smtpPortInt,
		MailUser:         mailUser,
		MailPassword:     mailPassword,
	}, nil
}
