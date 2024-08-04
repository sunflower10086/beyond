package logic

import (
	"beyond/application/like/rpc/internal/types"
	"beyond/application/like/rpc/pb"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"

	"beyond/application/like/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbupLogic {
	return &ThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbupLogic) Thumbup(in *pb.ThumbupRequest) (*pb.ThumbupResponse, error) {
	// TODO 逻辑暂时忽略
	// 1. 查询是否点过赞
	// 2. 计算当前内容的总点赞数和点踩数

	msg := &types.ThumbupMsg{
		BizId:    in.BizId,
		LikeType: in.LikeType,
		ObjId:    in.ObjId,
		UserId:   in.UserId,
	}
	threading.GoSafe(func() {
		data, err := json.Marshal(msg)
		if err != nil {
			l.Logger.Errorf("[Thumbup] marshal msg: %v error: %v", msg, err)
			return
		}

		if err = l.svcCtx.KqPusherClient.Push(l.ctx, string(data)); err != nil {
			l.Logger.Errorf("[Thumbup] kq push data: %s error: %v", data, err)
		}
	})

	return &pb.ThumbupResponse{}, nil
}
