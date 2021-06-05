package main

import (
	"fmt"
	"go-gin-api/internal/grpc/hello/rpc/client"
	"go-gin-api/internal/grpc/hello/rpc/server"
	"go-gin-api/internal/grpc/hello/rpc/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	Network = "tcp"
	Host = "127.0.0.1"
	Port = 11000
)

func main() {
	// 监听本地的 11000 端口
	lis, err := net.Listen(Network, fmt.Sprintf("%s:%d", Host, Port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer() // 创建 GRPC 服务器
	c := svc.NewContext() // 创建服务器上下文切换
	client.RegisterHelloServer(s, server.NewHelloServer(c))	// 在 GRPC 服务端注册服务

	reflection.Register(s) // 在给定的 GRPC 服务器上注册服务器反射服务

	fmt.Printf("success to serve, listen: %s:%d\n", Host, Port)

	defer s.Stop()

	// Serve 方法在 lis 上接受传入连接，为每个连接创建一个 ServerTransport 和 server 的 goroutine。
	// 该 goroutine 读取 GRPC 请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}