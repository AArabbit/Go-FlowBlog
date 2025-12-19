package router

import (
	"flow-blog/internal/controller"

	"github.com/gin-gonic/gin"
)

// InitPostsRouter 初始化文章路由
func InitPostsRouter() {

	// 文章列表
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		public.POST("/posts", controller.PostsList)
	})

	// 文章详情
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		public.GET("/posts/:id", controller.PostsDetail)
	})

	// 更新文章浏览量
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		public.POST("/up_views", controller.PostsViews)
	})

	// 添加文章
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		auth.POST("/add_post", controller.AddPosts)
	})

	// 更新文章
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		auth.PUT("/up_post/:id", controller.UpdatePosts)
	})

	// 删除文章
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		auth.DELETE("/delete_post/:id", controller.DeletePosts)
	})

	// 搜索文章
	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		public.POST("/search", controller.SearchPosts)
	})
}
