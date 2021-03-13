package cache

import (
	"context"
	"fmt"
	"gin-api/internal/config"
	"github.com/go-redis/redis/v8"
	"strings"
)

var (
	rds map[string]*redis.Client
	ctx = context.Background()
)

func InitRedis() {
	rds = make(map[string]*redis.Client)

	c := config.Get().Redis

	for key, value := range c {
		conn := redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", value.Addr, value.Port),
			Password:     value.Password,
			DB:           value.Db,
			MaxRetries:   value.MaxRetries,   // 最大的重试次数
			PoolSize:     value.PoolSize,     // 连接池最大连接数，默认为 CPU 数 * 10
			PoolTimeout:  value.PoolTimeout,  // 连接池超时时间
			MinIdleConns: value.MinIdleConns, // 最小空闲连接数
			IdleTimeout:  value.IdleTimeout,  // 空闲连接超时时间
			DialTimeout:  value.DialTimeout,  // 建立连接超时时间
			ReadTimeout:  value.ReadTimeout,  // 读超时时间
			WriteTimeout: value.WriteTimeout, // 写超时时间
		})

		_, err := conn.Ping(ctx).Result()
		if err == nil {
			rds[strings.ToLower(key)] = conn
		}
	}
}

func GetRedis(connection string) *redis.Client {
	return rds[connection]
}

func Close(connection string) error {
	return rds[connection].Close()
}

func CloseAll() error {
	for _, client := range rds {
		if err := client.Close(); err != nil {
			return err
		}
	}

	return nil
}