package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/saromanov/alerting/alerting"
	"github.com/saromanov/alerting/storage"
)

const defaultAddress = "localhost:6379"

// db provides handling of the redis
type db struct {
}

// Setup provides initialization of Redis
func Setup(c *alerting.Config) (storage.Storage, error) {
	address := defaultAddress
	if c.RedisAddress != "" {
		address = c.RedisAddress
	}
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: c.RedisPassword,
		DB:       c.RedisDB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("unable to setup redis: %v", err)
	}

	return &db{}, nil
}
