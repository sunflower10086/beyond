package errorx

import (
	"context"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func ErrHandler(err error) (int, any) {
	var body Response
	codeFrom := From(err)
	body.Code = codeFrom.Code
	body.Msg = codeFrom.Msg

	return http.StatusOK, body
}

func SuccessHandler(ctx context.Context, data any) any {
	var body Response
	body.Code = 200
	body.Msg = "success"
	body.Data = data
	return body
}
