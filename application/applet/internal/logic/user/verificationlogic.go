package user

import (
	"beyond/application/applet/internal/common/cache"
	"beyond/application/applet/internal/common/constants"
	"beyond/application/applet/internal/svc"
	"beyond/application/applet/internal/types"
	"beyond/application/user/rpc/user"
	"beyond/pkg/utils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type VerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerificationLogic {
	return &VerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Verification 发送验证码
func (l *VerificationLogic) Verification(req *types.VerificationRequest) (resp *types.VerificationResponse, err error) {
	// 校验手机号格式是否正确

	// 发送验证码

	// 30分钟内验证码不再变化
	code, err := getActivationCache(req.Mobile, l.svcCtx.Redis)
	if err != nil {
		logx.Errorf("getActivationCache mobile: %s error: %v", req.Mobile, err)
	}
	if len(code) == 0 {
		code = utils.RandomNumeric(6)
	}

	_, err = l.svcCtx.UserRPC.SendSms(l.ctx, &user.SendSmsRequest{
		Mobile: req.Mobile,
	})

	if err != nil {
		logx.Errorf("sendSms mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}

	// 保存验证码
	err = saveActivationCache(req.Mobile, code, l.svcCtx.Redis)
	if err != nil {
		logx.Errorf("saveActivationCache mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}

	return &types.VerificationResponse{}, nil

}

func getActivationCache(mobile string, rds *redis.Redis) (string, error) {
	key := cache.GetSMSCodeKey(mobile)
	return rds.Get(key)
}

func saveActivationCache(mobile, code string, rds *redis.Redis) error {
	key := cache.GetSMSCodeKey(mobile)
	return rds.Setex(key, code, constants.SMSCodeExpireTime)
}

func delActivationCache(mobile, code string, rds *redis.Redis) error {
	key := cache.GetSMSCodeKey(mobile)
	_, err := rds.Del(key)
	return err
}
