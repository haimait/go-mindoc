package utils

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	NewClient()
}

func NewClient() *redis.Client {
	if RedisClient != nil {
		return RedisClient
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: beego.AppConfig.String("RedisHost"),
		DB:   0, // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("redis connection failed: %v", err)
	}

	return RedisClient
}
