package router

import (
	"flow-blog/internal/controller"

	"github.com/gin-gonic/gin"
)

// InitUsersRouter 用户相关路由
func InitUsersRouter() {

	// 注册用户
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		public.POST("/register", controller.RegisterUser)
	})

	// 发送邮件
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		public.POST("/email", controller.GetEmailValidate)
	})

	// 登录
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		public.POST("/login", controller.LoginUser)
	})

	// 刷新token
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		public.POST("/refresh", controller.RefreshToken)
	})

	// 获取用户列表
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		auth.POST("/user_list", controller.UserList)
	})

	// 修改用户信息
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		auth.PUT("/up_user/:id", controller.UpdateUser)
	})

	// 删除用户信息
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		auth.DELETE("/delete_user/:id", controller.DeleteUser)
	})
}
