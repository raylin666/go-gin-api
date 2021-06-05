package server

import (
	"context"
	"go-gin-api/internal/grpc/hello/rpc/client"
	"go-gin-api/internal/grpc/hello/rpc/logic"
	"go-gin-api/internal/grpc/hello/rpc/svc"
)

type HelloServer struct {
	svcCtx *svc.Context
}

func NewHelloServer(ctx *svc.Context) *HelloServer {
	return &HelloServer{
		svcCtx: ctx,
	}
}

func (serv *HelloServer) GetSpeak(ctx context.Context, req *client.GetSpeakRequest) (*client.GetSpeakResponse, error) {
	l := logic.NewHelloLogic(ctx)
	return l.GetSpeak(req)
}