package controller

import (
	"context"
	"flow-blog/internal/api/dto"
	"flow-blog/internal/service"
	"flow-blog/pkg/app"
	"flow-blog/pkg/errcode"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type ThirdLoginController struct{}

func NewThirdLoginController() *ThirdLoginController {
	return &ThirdLoginController{}
}

// GithubLogin 第三方Github登录
func (*ThirdLoginController) GithubLogin(c *gin.Context) {
	oauthConf := service.GetOauthConf()
	state, err := service.GenerateRandomState()
	if err != nil {
		app.FailWithMsg(c, errcode.ServerError, err.Error())
		return
	}
	// 生成授权 URL
	url := oauthConf.AuthCodeURL(state)

	app.Success(c, gin.H{"url": url})
}

// GithubCallback 第三方github回调接口
func (*ThirdLoginController) GithubCallback(c *gin.Context) {
	var thirdCallback dto.ThirdCallback
	oauthConf := service.GetOauthConf()
	if err := c.ShouldBindJSON(&thirdCallback); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}

	// 校验state
	if err := service.CompareState(thirdCallback.State); err != nil {
		app.Fail(c, errcode.ServerError)
		return
	}

	// code换取token
	token, codeErr := oauthConf.Exchange(context.Background(), thirdCallback.Code)
	if codeErr != nil {
		app.FailWithMsg(c, errcode.ServerError, "token获取失败")
		return
	}
	fmt.Println("token:", token)

	user, accToken, refToken, userErr := service.GetGithubUserInfo(oauthConf, token)
	if userErr != nil {
		app.FailWithMsg(c, errcode.ServerError, "获取用户信息失败")
		log.Fatalf("获取用户信息失败: %v", userErr)
		return
	}

	app.Success(c, gin.H{
		"access_token":  accToken,
		"refresh_token": refToken,
		"userInfo":      user,
	})

}
