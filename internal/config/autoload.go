package config

import (
	"gin-api/internal/config/autoload"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var get ConfigYaml

type ConfigYaml struct {
	Environment string                       `yaml:"Environment"`
	App         autoload.App                 `yaml:"App"`
	Http        autoload.Http                `yaml:"Http"`
	Database    map[string]autoload.Database `yaml:"Database"`
	Redis       map[string]autoload.Redis	 `yaml:"Redis"`
	Jwt 		autoload.Jwt			     `yaml:"Jwt"`
}

func InitConfig() {
	cYaml, err := ioutil.ReadFile(".env.yml")
	if err != nil {
		panic(err)
	}

	_ = yaml.Unmarshal(cYaml, &get)
}

func Get() ConfigYaml {
	return get
}
