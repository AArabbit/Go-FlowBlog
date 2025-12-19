package router

import (
	"flow-blog/internal/controller"

	"github.com/gin-gonic/gin"
)

// InitCategoryRouter 初始化分类路由
func InitCategoryRouter() {

	RegisterRouter(func(auth *gin.RouterGroup, public *gin.RouterGroup) {
		public.GET("/categories", controller.CategoryList)
	})
}
