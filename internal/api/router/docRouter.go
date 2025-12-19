package router

import (
	"flow-blog/internal/api/controller"

	"github.com/gin-gonic/gin"
)

func DocsRouters(public *gin.RouterGroup) {
	docController := controller.NewDocController()

	// 获取常用文档列表
	public.GET("/docs", docController.GetDocList)
}
