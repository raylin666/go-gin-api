package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	Network        string
	Host           string
	Port           uint16
	RegisterServer func(*grpc.Server)
}

// 创建 gRPC 服务
func NewServer(server Server) {
	// 监听本地端口
	lis, err := net.Listen(server.Network, fmt.Sprintf("%s:%d", server.Host, server.Port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer() // 创建 gRPC 服务器

	// 在 gRPC 服务端注册服务
	if server.RegisterServer != nil {
		server.RegisterServer(s)
	}

	reflection.Register(s) // 在给定的 gRPC 服务器上注册服务器反射服务
	// Serve 方法在 lis 上接受传入连接，为每个连接创建一个 ServerTransport 和 server 的 goroutine
	// 该 goroutine 读取 gRPC 请求，然后调用已注册的处理程序来响应它们
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
