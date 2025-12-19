package controller

import (
	"flow-blog/internal/api/dto"
	"flow-blog/internal/service"
	"flow-blog/pkg/app"
	"flow-blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type VisitorController struct{}

// NewVisitorController 构造函数
func NewVisitorController() *VisitorController {
	return &VisitorController{}
}

// VisitorTraffic 记录访客
func (*VisitorController) VisitorTraffic(c *gin.Context) {
	ip := c.ClientIP()                 // 获取真实IP
	userAgent := c.Request.UserAgent() // 访客设备标识

	err := service.VisitorTrafficData(ip, userAgent)
	if err != nil {
		app.Fail(c, errcode.DBError)
		return
	}
	app.Success(c, gin.H{})
}

// GetVisitorTraffic 获取访客列表
func (*VisitorController) GetVisitorTraffic(c *gin.Context) {
	var p *dto.PageRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	list, total, hasMore, err := service.GetVisitorTrafficData(p)
	if err != nil {
		app.Fail(c, errcode.DBError)
		return
	}
	app.Success(c, gin.H{
		"visitor":  list,
		"total":    total,
		"has_more": hasMore,
	})
}

// DeleteVisitorTraffic 删除访客记录
func (*VisitorController) DeleteVisitorTraffic(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteVisitorTrafficData(id)
	if err != nil {
		app.Fail(c, errcode.DBError)
		return
	}
	app.Success(c, gin.H{"msg": "删除成功"})
}
