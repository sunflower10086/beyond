package logic

import (
	"beyond/application/like/mq/internal/svc"
	"context"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type ThumbupLogic struct {
	ctx        context.Context
	svcContext svc.ServiceContext
	logx.Logger
}

func NewThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbupLogic {
	return &ThumbupLogic{
		ctx:        ctx,
		svcContext: *svcCtx,
		Logger:     logx.WithContext(ctx),
	}
}

func (t *ThumbupLogic) Consume(ctx context.Context, key, val string) error {
	fmt.Printf("get key: %s val: %s\n", key, val)
	return nil
}

func Consumers(ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcContext.Config.KqConsumerConf, NewThumbupLogic(ctx, svcContext)),
	}
}
