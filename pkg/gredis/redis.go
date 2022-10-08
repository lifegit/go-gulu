package gredis

import (
	"github.com/go-redis/redis"
)

func CreateRedis(redisAddr, redisPassword string, idx int) (client *redis.Client, err error) {
	//initializing redis client
	client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // no password set
		DB:       idx,           // use default DB
	})
	_, err = client.Ping().Result() // ok: pong == "PONG"

	return
}
