package persistence

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/sauercrowd/urlshortener-backend/pkg/flags"
)

type Redis struct {
	client *redis.Client
}

func Create(f *flags.Flags) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", f.RedisHost),
		Password: f.RedisPassword,
		DB:       f.RedisDB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("Could not ping redis: %v", err)
	}
	return &Redis{client: client}, nil
}
