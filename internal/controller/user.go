package controller

import (
	"flow-blog/internal/model"
	"flow-blog/internal/model/request"
	"flow-blog/internal/service"
	"flow-blog/pkg/netRequest"
	"flow-blog/pkg/utils"

	"github.com/gin-gonic/gin"
)

// RegisterUser 用户注册
func RegisterUser(c *gin.Context) {
	var userRegisterInfo request.UserRegister
	if err := c.ShouldBindJSON(&userRegisterInfo); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "注册失败")
		return
	}
	codeErr := service.ValidateCode(&userRegisterInfo)
	if codeErr != nil {
		netRequest.Fail(c, netRequest.Unauthorized, codeErr.Error())
		return
	}
	netRequest.Success(c, gin.H{
		"msg": "注册成功",
	})
}

// GetEmailValidate 注册发送邮箱验证码
func GetEmailValidate(c *gin.Context) {
	// 获取邮箱
	var registerEmail request.EmailCode
	if err := c.ShouldBindJSON(&registerEmail); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	// 发送验证码
	sendNum, sendErr := service.SendEmailValidate([]string{registerEmail.Email})
	if sendErr != nil {
		netRequest.Fail(c, netRequest.ServerError, "验证码发送失败")
		return
	}
	if redisErr := service.ValidateCodeRedis(sendNum); redisErr != nil {
		netRequest.Fail(c, netRequest.ServerError, "验证码缓存失败")
		return
	}
	netRequest.Success(c, gin.H{"code": sendNum})
}

// LoginUser 用户登录
func LoginUser(c *gin.Context) {
	var loginInfo request.UserLogin
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	// 数据库比对登录信息
	user, err := service.ValidateLoginInfo(&loginInfo)
	if err != nil {
		netRequest.Fail(c, netRequest.Unauthorized, err.Error())
		return
	}
	accessToken, refreshToken, tokenErr := service.LoginToken(user.ID, loginInfo.UserName)
	if tokenErr != nil {
		netRequest.Fail(c, netRequest.ParamError, tokenErr.Error())
		return
	}

	netRequest.Success(c, gin.H{
		"msg":           "登陆成功",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"userInfo":      user,
	})
}

// RefreshToken 刷新token
func RefreshToken(c *gin.Context) {
	var token request.RefreshRequest
	if err := c.ShouldBindJSON(&token); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	// 验证RefreshToken
	claims, err := utils.ParseToken(token.RefreshToken, true) // true表示验证RefreshToken
	if err != nil {
		netRequest.Fail(c, netRequest.Unauthorized, "无效的刷新令牌")
		return
	}
	// 确认是否refresh
	if claims.TokenType != "refresh" {
		netRequest.Fail(c, netRequest.Unauthorized, "不是刷新令牌")
		return
	}
	// 生成一对新的token
	var userInfo model.User
	newAccess, newRefresh, tokenErr := service.LoginToken(userInfo.ID, userInfo.Username)
	if tokenErr != nil {
		netRequest.Fail(c, netRequest.ParamError, tokenErr.Error())
		return
	}
	netRequest.Success(c, gin.H{
		"msg":           "刷新令牌成功",
		"access_token":  newAccess,
		"refresh_token": newRefresh,
	})
}

// UserList 获取用户列表
func UserList(c *gin.Context) {
	var page request.PostsPageRequest
	if err := c.ShouldBindJSON(&page); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	userList, total, isEnd, err := service.GetUserList(&page)
	if err != nil {
		netRequest.Fail(c, netRequest.DBError, err.Error())
		return
	}
	netRequest.Success(c, gin.H{
		"total":    total,
		"has_more": isEnd,
		"useList":  userList,
	})
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	var user request.UpdateUser
	userId := c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}

	err := service.UpdateUserData(userId, &user)
	if err != nil {
		netRequest.Fail(c, netRequest.DBError, err.Error())
		return
	}

	netRequest.Success(c, gin.H{"msg": "更新成功"})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	err := service.DeleteUserData(userId)
	if err != nil {
		netRequest.Fail(c, netRequest.DBError, err.Error())
		return
	}

	netRequest.Success(c, gin.H{"msg": "删除成功"})
}
