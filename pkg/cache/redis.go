package cache

import (
	"fmt"
	"gin-api/internal/config"
	"github.com/go-redis/redis"
	"strings"
)

var rds map[string]*redis.Client

func InitRedis() {
	rds = make(map[string]*redis.Client)

	c := config.Get().Redis

	for key, value := range c {
		conn := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", value.Addr, value.Port),
			Password: value.Password,
			DB: value.Db,
		})

		_, err := conn.Ping().Result()
		if err == nil {
			rds[strings.ToLower(key)] = conn
		}
	}
}

func GetRedis(connection string) *redis.Client {
	return rds[connection]
}
