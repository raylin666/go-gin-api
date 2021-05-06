package config

import (
	"github.com/raylin666/go-gin-api/config/autoload"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configs = new(Config)

type Config struct {
	Environment string                       `yaml:"Environment"`
	App         autoload.App                 `yaml:"App"`
	Http        autoload.Http                `yaml:"Http"`
	Database    map[string]autoload.Database `yaml:"Database"`
	Redis       map[string]autoload.Redis	 `yaml:"Redis"`
	Jwt 		autoload.Jwt			     `yaml:"Jwt"`
	Logs 		autoload.Logs				 `yaml:"Logs"`
}

// 初始化加载配置文件, ymlEnvFileName 文件可复制该扩展包根目录的 .env.yml.default 文件进行修改, 它包含了所有配置项。
func InitAutoloadConfig(ymlEnvFileName string)  {
	cYaml, err := ioutil.ReadFile(ymlEnvFileName)
	if err != nil {
		panic(err)
	}

	_ = yaml.Unmarshal(cYaml, &configs)
}

// 获取配置项
func Get() *Config {
	return configs
}
