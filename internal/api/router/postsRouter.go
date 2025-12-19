package router

import (
	"flow-blog/internal/api/controller"

	"github.com/gin-gonic/gin"
)

// PostsRouters 初始化文章路由
func PostsRouters(public *gin.RouterGroup, private *gin.RouterGroup) {
	postsController := controller.NewPostsController()

	// 文章列表
	public.POST("/posts", postsController.PostsList)

	// 文章详情
	public.GET("/posts/:id", postsController.PostsDetail)

	// 获取每日推荐文章
	public.GET("/postsDaily", postsController.GetDailyPosts)

	// 更新文章浏览量
	public.POST("/up_views", postsController.PostsViews)

	// 搜索文章
	public.POST("/search", postsController.SearchPosts)

	// ======鉴权组======
	// redis缓存文章草稿
	private.POST("/draft_posts", postsController.DraftPosts)

	// 获取草稿
	private.GET("/draft_posts/:id", postsController.GetDraftPosts)

	// 添加文章
	private.POST("/add_post", postsController.AddPosts)

	// 更新文章
	private.PUT("/up_post/:id", postsController.UpdatePosts)

	// 删除文章
	private.DELETE("/delete_post/:id", postsController.DeletePosts)
}
