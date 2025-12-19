package router

import (
	"flow-blog/internal/controller"

	"github.com/gin-gonic/gin"
)

// InitBookmarkRouter 初始化收藏模块路由
func InitBookmarkRouter() {

	// 添加收藏
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		auth.POST("/add_bookmark", controller.AddBookmark)
	})

	// 取消收藏
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		auth.POST("/delete_bookmark", controller.DeleteBookmark)
	})
}
