package router

import (
	"flow-blog/internal/global"
	"flow-blog/internal/middleware"
	"flow-blog/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RegisterType = func(auth *gin.RouterGroup, public *gin.RouterGroup)

var registers []RegisterType

// RegisterRouter 路由注册函数
func RegisterRouter(fun RegisterType) {
	if fun == nil {
		return
	}
	registers = append(registers, fun)
}

// InitRouter 初始化总路由
func InitRouter() {
	gin.ForceConsoleColor()
	r := gin.Default()

	if err := r.SetTrustedProxies([]string{viper.GetString("trust.local")}); err != nil {
		global.Log.Error(err)
	}
	// 基础路径
	baseRouterPath := r.Group(viper.GetString("app.basePath"))
	// 基础路径下带token验证
	authRouterPath := baseRouterPath.Group("/auth")
	// 基础路径下开放接口
	publicRouterPath := baseRouterPath.Group("")

	// 挂载中间件
	authRouterPath.Use(middleware.JWTAuth())

	initModuleRouter()

	// 循环注册路由
	for _, routerFunction := range registers {
		routerFunction(authRouterPath, publicRouterPath)
	}

	err := r.Run(viper.GetString("app.port"))
	if err != nil {
		utils.RecordError("Run Server Error:", err)
	}
}

// 初始化模块路由
func initModuleRouter() {
	InitUsersRouter()
	InitCategoryRouter()
	InitPostsRouter()
	InitCommentRouter()
	InitBookmarkRouter()
}
