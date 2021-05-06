package init

import (
	"github.com/raylin666/go-gin-api/config"
	"github.com/raylin666/go-gin-api/environment"
)

// 初始化应用配置
type Config struct {
	YmlEnvFileName string
}

// 初始化应用
func (c *Config) InitApplication() {
	// 初始化加载配置文件
	config.InitAutoloadConfig(c.YmlEnvFileName)
	// 初始化环境
	environment.InitEnvironment()
}
