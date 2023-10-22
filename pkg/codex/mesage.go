package codex

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "操作成功",

	CodeInternalErr: "服务器开小差啦，稍后再来试一试",

	CodeMobilePhoneIsEmpty: "手机号码为空",
	CodeSMSCodeIsEmpty:     "验证码为空",

	CodeUserIsExist: "用户不存在",
	CodeMobileExist: "手机号已存在",
}

func (code ResCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeInternalErr]
	}
	return msg
}
