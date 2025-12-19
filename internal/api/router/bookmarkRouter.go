package router

import (
	"flow-blog/internal/api/controller"

	"github.com/gin-gonic/gin"
)

// BookmarkRouters 初始化收藏模块路由
func BookmarkRouters(private *gin.RouterGroup) {
	bookmarkController := controller.NewBookmarkController()
	// ======鉴权组======
	// 添加收藏
	private.POST("/add_bookmark", bookmarkController.AddBookmark)

	// 取消收藏
	private.POST("/delete_bookmark", bookmarkController.DeleteBookmark)
}
