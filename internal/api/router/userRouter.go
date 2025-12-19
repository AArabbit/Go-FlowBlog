package router

import (
	"flow-blog/internal/api/controller"

	"github.com/gin-gonic/gin"
)

// UserRoutes 用户相关路由
func UserRoutes(public *gin.RouterGroup, private *gin.RouterGroup) {
	userController := controller.NewUserController()

	// 注册用户
	public.POST("/register", userController.RegisterUser)

	// 发送邮件
	public.POST("/email", userController.GetEmailValidate)

	// 登录
	public.POST("/login", userController.LoginUser)

	// 刷新token
	public.POST("/refresh", userController.RefreshToken)

	// 更新密码
	public.POST("/up_userPass", userController.UpdateUserPass)

	// ======鉴权组======
	// 用户id获取用户信息
	private.POST("/user_info", userController.GetUserInfo)

	// 获取用户列表
	private.POST("/user_list", userController.UserList)

	// 修改用户信息
	private.PUT("/up_user/:id", userController.UpdateUser)

	// 删除用户信息
	private.DELETE("/delete_user/:id", userController.DeleteUser)
}
