package environment

import (
	"flag"
	"fmt"
	"github.com/raylin666/go-gin-api/config"
	"github.com/raylin666/go-gin-api/consts"
	"strings"
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
	return e.value == consts.ENVIRONMENT_DEV
}

// 是否测试环境
func (e *environment) IsTest() bool {
	return e.value == consts.ENVIRONMENT_TEST
}

// 是否预发布环境
func (e *environment) IsPre() bool {
	return e.value == consts.ENVIRONMENT_PRE
}

// 是否生产环境
func (e *environment) IsProd() bool {
	return e.value == consts.ENVIRONMENT_PROD
}

// 初始化环境
func InitEnvironment() {
	// go run main.go -env=prod
	env := flag.String("env", "", fmt.Sprintf("请输入运行环境:\n %s:开发环境\n %s:测试环境\n %s:预上线环境\n %s:正式环境\n", consts.ENVIRONMENT_DEV, consts.ENVIRONMENT_TEST, consts.ENVIRONMENT_PRE, consts.ENVIRONMENT_PROD))
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case consts.ENVIRONMENT_DEV:
		active = &environment{value: consts.ENVIRONMENT_DEV}
	case consts.ENVIRONMENT_TEST:
		active = &environment{value: consts.ENVIRONMENT_TEST}
	case consts.ENVIRONMENT_PRE:
		active = &environment{value: consts.ENVIRONMENT_PRE}
	case consts.ENVIRONMENT_PROD:
		active = &environment{value: consts.ENVIRONMENT_PROD}
	default:
		config_env := config.Get().Environment
		if config_env == consts.ENVIRONMENT_DEV {
			fmt.Println("Warning: '-" + consts.ENVIRONMENT_DEV + "' cannot be found, or it is illegal. The default '" + consts.ENVIRONMENT_DEV + "' will be used.")
		}
		active = &environment{value: config_env}
	}
}
