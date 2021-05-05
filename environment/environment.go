package environment

import (
	"flag"
	"fmt"
	"github.com/raylin666/go-gin-api/config"
	"strings"
)

const (
	ENVIRONMENT_DEV  = "dev"
	ENVIRONMENT_TEST = "test"
	ENVIRONMENT_PRE  = "pre"
	ENVIRONMENT_PROD = "prod"
)

var (
	active Environment
)

var _ Environment = (*environment)(nil)

type environment struct {
	value string
}

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsTest() bool
	IsPre() bool
	IsProd() bool
}

// 获取当前环境
func GetEnvironment() Environment {
	return active
}

// 获取当前环境值
func (e *environment) Value() string {
	return e.value
}

// 是否开发环境
func (e *environment) IsDev() bool {
	return e.value == ENVIRONMENT_DEV
}

// 是否测试环境
func (e *environment) IsTest() bool {
	return e.value == ENVIRONMENT_TEST
}

// 是否预发布环境
func (e *environment) IsPre() bool {
	return e.value == ENVIRONMENT_PRE
}

// 是否生产环境
func (e *environment) IsProd() bool {
	return e.value == ENVIRONMENT_PROD
}

// 初始化环境
func InitEnvironment() {
	// go run main.go -env=prod
	env := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n test:测试环境\n pre:预上线环境\n prod:正式环境\n")
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case ENVIRONMENT_DEV:
		active = &environment{value: ENVIRONMENT_DEV}
	case ENVIRONMENT_TEST:
		active = &environment{value: ENVIRONMENT_TEST}
	case ENVIRONMENT_PRE:
		active = &environment{value: ENVIRONMENT_PRE}
	case ENVIRONMENT_PROD:
		active = &environment{value: ENVIRONMENT_PROD}
	default:
		config_env := config.Get().Environment
		if config_env == ENVIRONMENT_DEV {
			fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default '" + ENVIRONMENT_DEV + "' will be used.")
		}
		active = &environment{value: config_env}
	}
}
