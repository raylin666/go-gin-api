package main

import (
	"fmt"
	"gin-api/internal/config"
	"gin-api/internal/env"
	"gin-api/internal/routers"
	"log"
	"net/http"
	"time"
)

func init()  {
	config.InitConfig()
	env.InitEnv()
}

func main()  {
	router := routers.InitRouter()

	host := config.Get().Http.Host
	port := config.Get().Http.Port

	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
		Handler: router,
		ReadTimeout: 60 * time.Second,
		WriteTimeout: 60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf(fmt.Sprintf("[info] start http server listening %s:%d", host, port))

	err := server.ListenAndServe()
	if err != nil {
		panic("HTTP 服务启动失败")
	}
}
