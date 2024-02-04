package db_store

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type DBClient struct {
	redisClient *redis.Client
}

func NewDBClient() *DBClient {
	dbClient := &DBClient{
		redisClient: redis.NewClient(
			&redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			},
		),
	}
	return dbClient
}

func (d *DBClient) Set(short_url, long_url string) error {
	// short_url is key, long_url is value
	ctx := context.Background()

	// Check if key exists
	_, err := d.redisClient.Get(ctx, short_url).Result()
	if err == nil {
		return fmt.Errorf("Value %s already exists", short_url)
	}

	err = d.redisClient.Set(ctx, short_url, long_url, 0).Err()
	return err
}

func (d *DBClient) Get(short_url string) (string, error) {
	// short_url is key, long_url is value
	ctx := context.Background()

	val, err := d.redisClient.Get(ctx, short_url).Result()
	if err != nil {
		return "", err
	}
	return val, err
}
