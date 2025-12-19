package middleware

import (
	"flow-blog/pkg/app"
	"flow-blog/pkg/errcode"
	"flow-blog/pkg/utils"

	"github.com/gin-gonic/gin"
)

// JWT 鉴权 & 登录过期验证

// JWTAuth 中间件
func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 获取 Authorization Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			app.FailWithMsg(c, errcode.Unauthorized, "没有携带token")
			c.Abort()
			return
		}

		// 解析token
		claims, err := utils.ParseToken(authHeader, false) // false 表示验证 AccessToken
		if err != nil {
			app.FailWithMsg(c, errcode.Unauthorized, "token过期")
			c.Abort()
			return
		}

		// 判断是accessToken才进行验证
		if claims.TokenType != "access" {
			app.FailWithMsg(c, errcode.Unauthorized, "不接受的token")
			c.Abort()
			return
		}

		// 用户信息存入上下文，通过gin实例调用
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		// 放行请求
		c.Next()
	}
}
