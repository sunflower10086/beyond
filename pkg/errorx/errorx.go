package errorx

import (
	"beyond/pkg/codex"
	"fmt"
	"net/http"
)

type CodeError struct {
	Type     string   `json:"type"` // 业务类型
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
	Metadata Metadata `json:"-"`
	Err      error    `json:"-"`
}
type Metadata map[string]interface{}

func New(t string, code int, message string) *CodeError {
	return &CodeError{
		Type: t,
		Code: code,
		Msg:  message,
	}
}

func WithCode(t string, code codex.ResCode) *CodeError {
	return &CodeError{
		Type: t,
		Code: int(code),
		Msg:  code.Msg(),
	}
}

func Internal(err error, format string, args ...interface{}) *CodeError {
	message := format
	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	}
	return New(http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError, message).WithError(err)
}

// 实现了error的Error()方法，CodeError就是一个error
func (e *CodeError) Error() string {
	if e.Err != nil {
		return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}

func (e *CodeError) WithError(err error) *CodeError {
	e.Err = err
	return e
}

func (e *CodeError) WithMetadata(metadata Metadata) *CodeError {
	e.Metadata = metadata
	return e
}

func From(err error) *CodeError {
	if err == nil {
		return nil
	}
	if errx, ok := err.(*CodeError); ok {
		return errx
	}
	return Internal(err, codex.CodeInternalErr.Msg())
}

func NotFound(format string, args ...any) *CodeError {
	message := fmt.Sprintf(format, args...)
	return New(http.StatusText(http.StatusNotFound), http.StatusNotFound, message)
}
