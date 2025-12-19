package router

import (
	"flow-blog/internal/api/controller"

	"github.com/gin-gonic/gin"
)

func ThirdLoginRouters(public *gin.RouterGroup) {
	thirdLoginController := controller.NewThirdLoginController()

	public.GET("/github_login", thirdLoginController.GithubLogin)

	public.POST("/github_callback", thirdLoginController.GithubCallback)
}
