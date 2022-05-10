package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisTimeOut = time.Second * 5

// Config represents all required configurations to create a connection to the redis DB.
type Config struct {
	DB       int
	Password string
	Host     string
	Port     int
	PoolSize int
}

type redisDB struct {
	db *redis.Client
}

// SetRedisTimeOut will set DialTimeout, ReadTimeout, WriteTimeout on Redis Client connection.
func SetRedisTimeOut(timeOut time.Duration) {
	redisTimeOut = timeOut
}

// NewRedisDB creates a connection to the redis DB.
func NewRedisDB(cfg *Config) (*redisDB, error) {
	client := redis.NewClient(&redis.Options{
		DB:           cfg.DB,
		Password:     cfg.Password,
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		PoolSize:     cfg.PoolSize,
		DialTimeout:  redisTimeOut,
		ReadTimeout:  redisTimeOut,
		WriteTimeout: redisTimeOut,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return &redisDB{db: client}, nil
}
