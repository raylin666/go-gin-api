package cache

import (
	"github.com/raylin666/go-utils/cache"
	"go-gin-api/config"
)

func InitRedis() {
	var (
		c  map[string]*cache.RedisConfig
		cr = config.Get().Redis
	)

	c = make(map[string]*cache.RedisConfig, len(cr))

	for key, value := range cr {
		rc := &cache.RedisConfig{
			Addr:         value.Addr,
			Port:         value.Port,
			Password:     value.Password,
			Db:           value.Db,
			MaxRetries:   value.MaxRetries,
			PoolSize:     value.PoolSize,
			PoolTimeout:  value.PoolTimeout,
			MinIdleConns: value.MinIdleConns,
			IdleTimeout:  value.IdleTimeout,
			DialTimeout:  value.DialTimeout,
			ReadTimeout:  value.ReadTimeout,
			WriteTimeout: value.WriteTimeout,
		}
		c[key] = rc
	}

	cache.InitRedis(c)
}

// 获取链接
func GetRedis(connection string) *cache.Redis {
	return cache.GetRedis(connection)
}

// 关闭链接
func CloseRedis(connection string) error {
	return cache.CloseRedis(connection)
}

// 关闭所有链接
func CloseAllRedis() error {
	return cache.CloseAllRedis()
}
