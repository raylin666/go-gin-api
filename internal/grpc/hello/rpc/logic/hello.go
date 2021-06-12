package logic

import (
	"context"
	"go-gin-api/internal/grpc/hello/rpc/client"
	"go-gin-api/internal/grpc/hello/rpc/svc"
)

type HelloLogic struct {
	ctx    context.Context
	svcCtx *svc.Context
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.Context) *HelloLogic {
	return &HelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) GetSpeak(req *client.GetSpeakRequest) (*client.GetSpeakResponse, error) {
	var message string
	res_content := l.svcCtx.Model.KeywordTest.GetFirst(req.Content)
	if res_content.Keyword == "" {
		message = "没有查询到数据哦 ~"
	} else {
		message = res_content.ResContent
	}

	return &client.GetSpeakResponse{
		Message: message,
	}, nil
}
