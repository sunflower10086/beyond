package user

import (
	"context"
	"errors"
	"strings"

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

	}

	return
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
		return errors.New("verification code failed")
	}

	return nil
}
