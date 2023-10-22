package codex

type ResCode int

const (
	CodeSuccess ResCode = 200

	CodeInvalidParams ResCode = 400
	CodeInternalErr   ResCode = 500

	CodeMobilePhoneIsEmpty ResCode = 40001
	CodeSMSCodeIsEmpty     ResCode = 40002

	CodeUserIsExist ResCode = 50001
	CodeMobileExist ResCode = 50002
)
