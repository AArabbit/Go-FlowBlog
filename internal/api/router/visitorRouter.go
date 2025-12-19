package router

import (
	"flow-blog/internal/api/controller"

	"github.com/gin-gonic/gin"
)

func VisitorRouters(public *gin.RouterGroup, private *gin.RouterGroup) {
	visitorController := controller.NewVisitorController()

	// 记录访客
	public.GET("/visit", visitorController.VisitorTraffic)

	// 获取访客列表
	private.POST("/visit_traffic", visitorController.GetVisitorTraffic)

	// 删除记录
	private.DELETE("/delete_traffic/:id", visitorController.DeleteVisitorTraffic)
}
