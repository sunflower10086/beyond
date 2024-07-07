package logic

import (
	"beyond/application/user/rpc/internal/code"
	"beyond/application/user/rpc/internal/model"
	"beyond/application/user/rpc/internal/svc"
	"beyond/application/user/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if in.Username == "" {
		return nil, code.RegisterNameEmpty
	}

	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username: in.Username,
		Avatar:   in.Avatar,
		Mobile:   in.Mobile,
	})

	if err != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, err
	}

	userId, err := ret.LastInsertId()
	if err != nil {
		logx.Errorf("LastInsertId error: %v", err)
		return nil, err
	}

	return &pb.RegisterResponse{UserId: userId}, nil
}
