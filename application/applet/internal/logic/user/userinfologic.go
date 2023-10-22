package user

import (
	"beyond/application/applet/internal/svc"
	"beyond/application/applet/internal/types"
	"beyond/application/user/rpc/user"
	"beyond/pkg/codex"
	"beyond/pkg/errorx"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	var (
		userIdNum json.Number
		ok        bool

		userId int64
	)

	if userIdNum, ok = l.ctx.Value("userId").(json.Number); !ok {
		return nil, errorx.WithCode("获取用户信息失败", codex.CodeInternalErr)
	}

	if userId, err = userIdNum.Int64(); err != nil {
		err = errorx.WithCode("获取用户信息失败", codex.CodeInternalErr)
		return
	}

	userInfo, err := l.svcCtx.UserRPC.FindById(l.ctx, &user.FindByIdRequest{
		UserId: userId,
	})

	if err != nil {
		return nil, err
	}
	resp = new(types.UserInfoResponse)
	resp.UserId = userInfo.UserId
	//resp.Avatar = userInfo.Avatar
	resp.Username = userInfo.Username

	return resp, nil
}
