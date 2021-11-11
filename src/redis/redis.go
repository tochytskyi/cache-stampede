package redis

import (
	"log"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "treatfield-redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		log.Fatalln("Redis error", err)
	}

	log.Println(pong)
}

func GetInstance() *redis.Client {
	if client == nil {
		Init()
	}
	return client
}
