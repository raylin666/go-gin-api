package logic

import (
	"context"
	"go-gin-api/internal/grpc/hello/rpc/client"
)

type HelloLogic struct {
	ctx context.Context
}

func NewHelloLogic(ctx context.Context) *HelloLogic {
	return &HelloLogic{
		ctx: ctx,
	}
}

func (l *HelloLogic) GetSpeak(req *client.GetSpeakRequest) (*client.GetSpeakResponse, error) {
	return &client.GetSpeakResponse{}, nil
}
