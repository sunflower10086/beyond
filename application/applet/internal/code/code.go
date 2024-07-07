package code

import "beyond/pkg/codex"

var (
	RegisterMobileEmpty   = codex.New(10001, "注册手机号不能为空")
	VerificationCodeEmpty = codex.New(100002, "验证码不能为空")
	MobileHasRegistered   = codex.New(100003, "手机号已经注册")
	LoginMobileEmpty      = codex.New(100003, "手机号不能为空")
	RegisterPasswdEmpty   = codex.New(100004, "密码不能为空")
)
