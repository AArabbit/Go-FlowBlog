package router

import (
	"flow-blog/internal/api/controller"

	"github.com/gin-gonic/gin"
)

// CategoryRouters 初始化分类路由
func CategoryRouters(public *gin.RouterGroup) {
	categoryController := controller.NewCategoryController()

	// 分类列表
	public.GET("/categories", categoryController.CategoryList)

	// 分类下的文章
	public.POST("/categoriesPosts", categoryController.CategoryPostList)
}
