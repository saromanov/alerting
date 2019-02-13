package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/saromanov/alerting/alerting"
	"github.com/saromanov/alerting/storage"
	"github.com/saromanov/alerting/structs"
)

const defaultAddress = "localhost:6379"

// db provides handling of the redis
type db struct {
	client *redis.Client
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

	return &db{
		client: client,
	}, nil
}

// Set provides inserting of the message struct to redis
func (d *db) Set(s *structs.Message) error {
	return nil
}

// Get provides getting of the message by id
func (d *db) Get(id string) (*structs.Message, error) {
	return nil, nil
}

// View provides searching of teh messages
func (d *db) View() ([]*structs.Message, error) {
	return nil, nil
}
