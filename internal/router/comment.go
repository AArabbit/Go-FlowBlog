package router

import (
	"flow-blog/internal/controller"

	"github.com/gin-gonic/gin"
)

// InitCommentRouter 初始化评论模块路由
func InitCommentRouter() {

	// 文章评论列表
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		public.POST("/comments", controller.PostsComments)
	})

	// 新增评论
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		auth.POST("/add_comment", controller.AddComment)
	})

	// 删除评论
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		auth.POST("/delete_comment", controller.DeleteComment)
	})
}
