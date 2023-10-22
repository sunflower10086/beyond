package user

import (
	"beyond/application/applet/internal/svc"
	"beyond/application/applet/internal/types"
	"beyond/application/user/rpc/user"
	"beyond/pkg/codex"
	"beyond/pkg/errorx"
	"beyond/pkg/jwt"
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, errorx.WithCode("login", codex.CodeMobilePhoneIsEmpty)
	}

	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, errorx.WithCode("login", codex.CodeSMSCodeIsEmpty)
	}

	if err = VerificationCode(req.Mobile, req.VerificationCode, l.svcCtx.Redis); err != nil {
		err = errorx.Internal(err, err.Error()).WithError(err).WithMetadata(errorx.Metadata{
			"Mobile":           req.Mobile,
			"VerificationCode": req.VerificationCode,
		})
		return
	}

	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{Mobile: req.Mobile})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u == nil || u.UserId == 0 {
		return nil, errorx.WithCode("login", codex.CodeUserIsExist)
	}
	token, err := jwt.CreateToken(l.svcCtx.Config.Auth.AccessSecret, int(u.UserId))
	if err != nil {
		return nil, errorx.WithCode("login", codex.CodeInternalErr).WithError(err).WithMetadata(errorx.Metadata{
			"userId": u.UserId,
		})
	}

	_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.Redis)

	return &types.LoginResponse{
		UserId: u.UserId,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
