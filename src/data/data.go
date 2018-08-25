package data

import (
	. "configuration"
	"github.com/go-redis/redis"
	"log"
)

var RedisClient *redis.Client = nil

func InitRedisClient()  {
	config := Config.Redis

	if RedisClient == nil {
		RedisClient = redis.NewClient(&redis.Options{
			Addr: 		        config.Host + ":" + config.Port,
			Password: 	        config.Password,
			DB: 		        config.DB,
			MaxRetries:         10,
			OnConnect:          onRedisConnect,
		})
	}

	RedisClient.Ping().Result()
}

func onRedisConnect(conn *redis.Conn) error {
	log.Println("Connection Redis succeed, DB Size: ", conn.DBSize())
	return nil
}