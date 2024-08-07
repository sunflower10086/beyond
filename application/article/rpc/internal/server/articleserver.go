// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package server

import (
	"context"

	"beyond/application/article/rpc/internal/logic"
	"beyond/application/article/rpc/internal/svc"
	"beyond/application/article/rpc/pb"
)

type ArticleServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedArticleServer
}

func NewArticleServer(svcCtx *svc.ServiceContext) *ArticleServer {
	return &ArticleServer{
		svcCtx: svcCtx,
	}
}

func (s *ArticleServer) Publish(ctx context.Context, in *pb.PublishRequest) (*pb.PublishResponse, error) {
	l := logic.NewPublishLogic(ctx, s.svcCtx)
	return l.Publish(in)
}
