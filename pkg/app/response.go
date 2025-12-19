package app

import (
	"flow-blog/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应结构体
type Response struct {
	Code int         `json:"code"` // 业务码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 数据
}

// result 基础响应
func result(c *gin.Context, code int, msg string, data interface{}) {
	// 总是返回 HTTP 200，前端根据内部 code 判断业务逻辑
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	result(c, int(errcode.Success), errcode.Success.Msg(), data)
}

// SuccessWithMsg 带自定义消息的成功响应
func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	result(c, int(errcode.Success), msg, data)
}

// Fail 失败响应 (使用定义好的 ErrorCode)
func Fail(c *gin.Context, errorCode errcode.ErrorCode) {
	result(c, int(errorCode), errorCode.Msg(), nil)
}

// FailWithMsg 失败响应 (自定义消息)
func FailWithMsg(c *gin.Context, errorCode errcode.ErrorCode, msg string) {
	result(c, int(errorCode), msg, nil)
}
