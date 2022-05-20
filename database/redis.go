package database

import (
	"github.com/go-redis/redis/v7"
)

func StartRedis() *redis.Client {
	dsn := "localhost:16379"
	// dsn := os.Getenv("REDIS_DSN")
	client := redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}

var REDISCLIENT *redis.Client = StartRedis()
