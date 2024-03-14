package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func connect() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", DB: 0,
		Protocol: 3,
	})

	status, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	log.Println(status)
}
