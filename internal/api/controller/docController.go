package controller

import (
	"flow-blog/internal/service"
	"flow-blog/pkg/app"
	"flow-blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type DocController struct{}

// NewDocController 构造函数
func NewDocController() *DocController {
	return &DocController{}
}

// GetDocList 获取常用文档列表
func (*DocController) GetDocList(c *gin.Context) {
	docs, err := service.GetDocListData()
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{"docList": docs})
}
