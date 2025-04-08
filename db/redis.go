package db

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

// Инициализация Redis
func InitRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Fatal("Missing REDIS_ADDR")
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Fatal("Redis connection error: ", err)
	} else {
		log.Println("Connected to Redis")
	}
}

// Возвращает клиент Redis
func GetRedisClient() *redis.Client {
	return redisClient
}

// Возвращает контекст
func GetContext() context.Context {
	return ctx
}
