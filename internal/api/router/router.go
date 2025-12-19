package router

import (
	"flow-blog/internal/api/middleware"
	"flow-blog/pkg/utils"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// InitRouter 初始化总路由
func InitRouter() {
	var err error
	gin.ForceConsoleColor() // 开启彩色日志
	r := gin.Default()
	// 信任docker转发地址，为了实现获取外网真实IP
	err = r.SetTrustedProxies([]string{"127.0.0.1", "172.16.0.0/12"})
	if err != nil {
		fmt.Println("设置信任代理失败:", err)
	}

	// 生产环境配置
	if viper.GetString("mode.env") == "production" {
		gin.SetMode(gin.ReleaseMode)
		// 跨域
		r.Use(cors.New(cors.Config{
			// 只允许前端域名
			AllowOrigins:     []string{"https://www.rabbitwebsite.top", "https://rabbitwebsite.top"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	// 统一路由组
	blogApi := r.Group(viper.GetString("app.basePath"))
	// 开放接口组
	publicGroup := blogApi.Group("")
	// 鉴权接口组
	privateGroup := blogApi.Group("/auth")
	// 挂载鉴权中间件
	privateGroup.Use(middleware.JWTAuth())
	{
		// 初始化模块路由
		UserRoutes(publicGroup, privateGroup)
		PostsRouters(publicGroup, privateGroup)
		CommentRouters(publicGroup, privateGroup)
		VisitorRouters(publicGroup, privateGroup)
		DocsRouters(publicGroup)
		CategoryRouters(publicGroup)
		BookmarkRouters(privateGroup)
		ThirdLoginRouters(publicGroup)
	}

	if err = r.Run(viper.GetString("app.port")); err != nil {
		utils.RecordError("Run Server Error:", err)
	}
	fmt.Println("服务启动成功...")
}
