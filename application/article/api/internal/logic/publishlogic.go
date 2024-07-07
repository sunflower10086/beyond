package logic

import (
	"beyond/application/article/api/code"
	"beyond/application/article/rpc/article"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"

	"beyond/application/article/api/internal/svc"
	"beyond/application/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const minContentLen = 80

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishRequest) (resp *types.PublishResponse, err error) {
	if len(req.Title) == 0 {
		return nil, code.ArticleTitleEmpty
	}
	if len(req.Content) < minContentLen {
		return nil, code.ArticleContentTooFewWords
	}
	if len(req.Cover) == 0 {
		return nil, code.ArticleCoverEmpty
	}

	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		logc.Errorf(l.ctx, "[PublishLogic] error: %v， 获取用户userId失败", err)
		return nil, err
	}

	pret, err := l.svcCtx.ArticleRPC.Publish(l.ctx, &article.PublishRequest{
		UserId:      userId,
		Title:       req.Title,
		Cover:       req.Cover,
		Content:     req.Content,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &types.PublishResponse{ArticleId: pret.ArticleId}, nil
}
