package controller

import (
	"flow-blog/internal/api/dto"
	"flow-blog/internal/model"
	"flow-blog/internal/service"
	"flow-blog/pkg/app"
	"flow-blog/pkg/errcode"
	"flow-blog/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// NewUserController 构造函数
func NewUserController() *UserController {
	return &UserController{}
}

// RegisterUser 用户注册
func (*UserController) RegisterUser(c *gin.Context) {
	var userRegisterInfo dto.UserRegister
	if err := c.ShouldBindJSON(&userRegisterInfo); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	codeErr := service.ValidateCode(&userRegisterInfo)
	if codeErr != nil {
		app.FailWithMsg(c, errcode.Unauthorized, codeErr.Error())
		return
	}
	app.Success(c, gin.H{"msg": "注册成功"})
}

// GetEmailValidate 注册发送邮箱验证码
func (*UserController) GetEmailValidate(c *gin.Context) {
	// 获取邮箱
	var registerEmail dto.EmailCode
	if err := c.ShouldBindJSON(&registerEmail); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	// 发送验证码
	sendNum, sendErr := service.SendEmailValidate([]string{registerEmail.Email})
	if sendErr != nil {
		app.FailWithMsg(c, errcode.ServerError, "验证码发送失败")
		return
	}
	if redisErr := service.ValidateCodeRedis(sendNum); redisErr != nil {
		app.FailWithMsg(c, errcode.ServerError, "验证码缓存失败")
		return
	}
	app.Success(c, gin.H{"code": sendNum})
}

// LoginUser 用户登录
func (*UserController) LoginUser(c *gin.Context) {
	var loginInfo dto.UserLogin
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	// 数据库比对登录信息
	user, err := service.ValidateLoginInfo(&loginInfo)
	if err != nil {
		app.FailWithMsg(c, errcode.Unauthorized, err.Error())
		return
	}
	accessToken, refreshToken, tokenErr := service.LoginToken(user.ID, loginInfo.UserName)
	if tokenErr != nil {
		app.FailWithMsg(c, errcode.ParamError, tokenErr.Error())
		return
	}
	app.Success(c, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"userInfo":      user,
	})
}

// RefreshToken 刷新token
func (*UserController) RefreshToken(c *gin.Context) {
	var token dto.RefreshRequest
	if err := c.ShouldBindJSON(&token); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	// 验证RefreshToken
	claims, err := utils.ParseToken(token.RefreshToken, true) // true表示验证RefreshToken
	if err != nil {
		app.FailWithMsg(c, errcode.Unauthorized, "无效的刷新令牌")
		return
	}
	// 确认是否refresh
	if claims.TokenType != "refresh" {
		app.FailWithMsg(c, errcode.Unauthorized, "不是刷新令牌")
		return
	}
	// 生成一对新的token
	var userInfo model.UserModel
	newAccess, newRefresh, tokenErr := service.LoginToken(userInfo.ID, userInfo.Username)
	if tokenErr != nil {
		app.FailWithMsg(c, errcode.ParamError, tokenErr.Error())
		return
	}
	app.Success(c, gin.H{
		"access_token":  newAccess,
		"refresh_token": newRefresh,
	})
}

// GetUserInfo ID获取用户信息
func (*UserController) GetUserInfo(c *gin.Context) {
	var userId dto.UserInfo
	if err := c.ShouldBindJSON(&userId); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	userInfo, err := service.GetUserInfoData(&userId)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{"userInfo": userInfo})
}

// UserList 获取用户列表
func (*UserController) UserList(c *gin.Context) {
	var page dto.PageRequest
	if err := c.ShouldBindJSON(&page); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	userList, total, isEnd, err := service.GetUserList(&page)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{
		"total":    total,
		"has_more": isEnd,
		"useList":  userList,
	})
}

// UpdateUser 更新用户信息
func (*UserController) UpdateUser(c *gin.Context) {
	var user dto.UpdateUser
	userId := c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}

	err := service.UpdateUserData(userId, &user)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}

	app.Success(c, gin.H{"msg": "更新成功"})
}

// UpdateUserPass 更新用户密码
func (*UserController) UpdateUserPass(c *gin.Context) {
	var userPass dto.UpdatePass
	if err := c.ShouldBindJSON(&userPass); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	fmt.Println("参数成功接收:", userPass)

	err := service.UpdateUserPassData(&userPass)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{"msg": "密码修改成功"})
}

// DeleteUser 删除用户
func (*UserController) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	err := service.DeleteUserData(userId)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}

	app.Success(c, gin.H{"msg": "删除成功"})
}
