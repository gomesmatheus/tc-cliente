package database

import "github.com/go-redis/redis/v8"

func NewRedisDb() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cliente-redis:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}
