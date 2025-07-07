package helper

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/zeromicro/go-zero/core/logx"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(data interface{}) any {
	if data == nil {
		return Response{
			Code: 200,
			Msg:  "请求成功",
		}
	}

	// 判断是否有 Rows 字段
	val := reflect.Indirect(reflect.ValueOf(data))
	if val.Kind() == reflect.Struct && val.FieldByName("Rows").IsValid() {
		base := map[string]interface{}{
			"code": 200,
			"msg":  "请求成功",
		}
		for k, v := range structToMap(data) {
			base[k] = v
		}
		return base
	}

	// 普通响应
	return Response{
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
		logx.WithContext(ctx).Errorf("【%s】 err %v", name, err)
		return http.StatusOK, Fail(err.Error())
	}
}

func structToMap(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	if obj == nil {
		return m
	}
	b, _ := json.Marshal(obj)
	_ = json.Unmarshal(b, &m)
	return m
}
