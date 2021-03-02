package main

import (
	"gin-api/internal/config"
	"gin-api/pkg/logger"
)

func init()  {
	config.InitConfig()
}

func main()  {
	logger.NewWriteInstance("app").WithFields(logger.Fields{
		"msg": "I'm logger",
	}.Fields()).Info("hhh")
}
