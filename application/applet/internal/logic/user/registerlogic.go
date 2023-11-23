package user

import (
	"beyond/application/user/rpc/user"
	"beyond/pkg/encrypt"
	"beyond/pkg/jwt"
	"context"
	"errors"
	"strings"

	"beyond/application/applet/internal/code"
	"beyond/application/applet/internal/svc"
	"beyond/application/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.RegisterMobileEmpty
	}

	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		req.Password = encrypt.EncPassword(req.Password)
	}

	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, code.RegisterPasswdEmpty
	}

	err = VerificationCode(req.Mobile, req.VerificationCode, l.svcCtx.Redis)
	if err != nil {
		return nil, err
	}

	// 看这个手机号是否绑定了其他人
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{
		Mobile: req.Mobile,
	})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u != nil && u.UserId > 0 {
		return nil, code.MobileHasRegistered
	}

	// 开始注册
	regRet, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterRequest{
		Username: req.Name,
		Mobile:   req.Mobile,
	})
	if err != nil {
		logx.Errorf("Register error: %v", err)
		return nil, err
	}

	token, err := jwt.CreateToken(l.svcCtx.Config.Auth.AccessSecret, int(regRet.UserId))
	if err != nil {
		logx.Errorf("Register error: %v", err)
		return nil, err
	}

	// 注册成功删除验证码
	_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.Redis)

	return &types.RegisterResponse{
		UserId: regRet.UserId,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}

func VerificationCode(mobile, code string, redis *redis.Redis) error {
	cacheCode, err := getActivationCache(mobile, redis)
	if err != nil {
		return err
	}

	if cacheCode == "" {
		return errors.New("verification code expired")
	}

	if cacheCode != code {
		return errors.New("verification code expired")
	}

	return nil
}
