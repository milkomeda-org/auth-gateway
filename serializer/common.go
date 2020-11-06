package serializer

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"oa-auth/configs"
)

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "未登录",
	}
}

// NoAccess 无权访问
func NoAccess() Response {
	return Response{
		Code: CodeNoRightErr,
		Msg:  "无权访问",
	}
}

// Err 通用错误处理
func Err(code int, msg string, err error) Response {
	res := Response{
		Code: code,
	}
	if "" == msg {
		res.Msg = "操作失败"
	} else {
		res.Msg = msg
	}
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// ParamErr 参数错误
func ParamErr(msg string, err error) Response {
	if "" == msg {
		msg = "参数错误"
	}
	return Err(401, msg, err)
}

// Success 处理成功
func Success(data interface{}) Response {
	res := Response{
		Code: 0,
		Msg:  "操作成功",
		Data: data,
	}
	return res
}

// Failed 处理失败
func Failed(err error) Response {
	res := Response{
		Code:  -1,
		Msg:   "操作失败",
		Error: err.Error(),
	}
	return res
}

// I18Error 返回错误消息
func I18Error(err error) Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := configs.T(fmt.Sprintf("Field.%s", e.Field()))
			tag := configs.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
			return ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ParamErr("JSON类型不匹配", err)
	}

	return ParamErr("参数错误", err)
}
