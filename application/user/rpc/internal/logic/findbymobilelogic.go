package logic

import (
	"context"
	"fmt"

	"beyond/application/user/rpc/internal/svc"
	"beyond/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByMobileLogic) FindByMobile(in *pb.FindByMobileRequest) (*pb.FindByMobileResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	fmt.Println(user, err)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return &pb.FindByMobileResponse{}, nil
		}
		logx.Errorf("FindById userId: %s error: %v", in.Mobile, err)
		return nil, err
	}

	return &pb.FindByMobileResponse{
		UserId:   user.Id,
		Username: user.Username,
		Avatar:   user.Avatar,
	}, nil
}
