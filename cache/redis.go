package cache

import (
	"github.com/go-redis/redis"
	"log"
)

type Config struct {
	DB       int
	Addr     string
	Password string
}

func NewRedis(c *Config)(client *redis.Client) {
	client = redis.NewClient(&redis.Options {
		Addr:     c.Addr,
		Password: c.Password, // no password set
		DB:       c.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("redis connect ping failed, err:%v", err)
		panic(err)
	}else {
		log.Printf("redis connect ping response:%v", pong)
	}
	return
}