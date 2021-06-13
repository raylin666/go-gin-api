package main

import (
	"go-gin-api/internal/api/router"
	"go-gin-api/internal/api/server"
	"go-gin-api/internal/initx"
)

func init()  {
	initx.NewInitApp().Run()
}

func main()  {
	r := &router.Router{}
	server.NewHttpServer(r)
}
