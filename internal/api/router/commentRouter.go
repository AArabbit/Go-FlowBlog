package router

import (
	"flow-blog/internal/api/controller"

	"github.com/gin-gonic/gin"
)

// CommentRouters 初始化评论模块路由
func CommentRouters(public *gin.RouterGroup, private *gin.RouterGroup) {
	commentController := controller.NewCommentsController()

	// 文章评论列表
	public.POST("/comments", commentController.PostsComments)

	// ======鉴权组======
	// 新增评论
	private.POST("/add_comment", commentController.AddComment)

	// 删除评论
	private.POST("/delete_comment", commentController.DeleteComment)
}
