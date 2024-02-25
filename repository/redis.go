package repository

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewCache() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successo redis")

	return client
}
