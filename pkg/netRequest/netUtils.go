package netRequest

import (
	"flow-blog/internal/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess  = 0    // 成功
	ParamError   = 1001 // 参数错误
	NotFound     = 1002 // 数据不存在
	DBError      = 1003 // 数据库错误
	Unauthorized = 1004 // 未授权/未登录
	ServerError  = 1005 // 服务内部错误
)

// 响应封装
func resultJson(c *gin.Context, code int, msg string, data interface{}) {
	// 统一返回200，用内部码控制状态
	c.JSON(http.StatusOK, response.ResultResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Success 成功响应，带数据
func Success(c *gin.Context, data interface{}) {
	resultJson(c, CodeSuccess, "success", data)
}

// Fail 失败响应
func Fail(c *gin.Context, code int, msg string) {
	resultJson(c, code, msg, nil)
}
