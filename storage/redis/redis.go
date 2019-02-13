package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/saromanov/alerting/storage"
)

// db provides handling of the redis
type db struct {
}

// Setup provides initialization of Redis
func Setup() (storage.Storage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("unable to setup redis: %v", err)
	}

	return &db{}, nil
}
