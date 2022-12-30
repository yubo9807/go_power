package service

import "server/src/middleware"

func ErrorCapture(code int, msg string) {
	middleware.State.Code = code
	middleware.State.Message = msg
}

// 参数错误
func ErrorParams() {
	middleware.State.Code = 406
	middleware.State.Message = "params error"
}

// 未授权
func ErrorUnauthorized() {
	middleware.State.Code = 401
	middleware.State.Message = "unauthorized"
}

// 自定义错误消息
func ErrorCustom(msg string) {
	middleware.State.Code = 500
	middleware.State.Message = msg
}

// 请求成功，并返回数据
func SuccessData(data interface{}) {
	middleware.State.Code = 200
	middleware.State.Message = "success"
	middleware.State.Data = data
}

// 请求成功
func Success() {
	SuccessData("success")
}

func GetData() interface{} {
	return middleware.State.Data
}
