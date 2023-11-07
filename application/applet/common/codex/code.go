package codex

import "beyond/pkg/codex"

var (
	CodeSuccess = codex.New(200, "操作成功")

	CodeInvalidParams = codex.New(400, "参数错误")
	CodeInternalErr   = codex.New(500, "服务器开小差啦，稍后再来试一试")

	CodeMobilePhoneIsEmpty = codex.New(40001, "手机号码为空")
	CodeSMSCodeIsEmpty     = codex.New(40002, "验证码为空")

	CodeUserNotExist = codex.New(50001, "用户不存在")
	CodeMobileExist  = codex.New(50002, "手机号已存在")
)
