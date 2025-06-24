package helper

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code: 200,
		Msg:  "请求成功",
		Data: data,
	}
}

func Fail(err string) *Response {
	return &Response{
		Code: 500,
		Msg:  err,
		Data: nil,
	}
}

func OkHandler(_ context.Context, v interface{}) any {
	return Success(v)
}

func ErrHandler(name string) func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		// 日志记录
		logx.WithContext(ctx).Errorf("【%s】 err %v", name, err)
		return http.StatusOK, Fail(err.Error())
	}
}
