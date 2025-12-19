package controller

import (
	"flow-blog/internal/service"
	"flow-blog/pkg/netRequest"

	"github.com/gin-gonic/gin"
)

// CategoryList 查询分类列表
func CategoryList(c *gin.Context) {
	list, err := service.GetCategoryList()
	if err != nil {
		netRequest.Fail(c, netRequest.NotFound, "没有数据")
		return
	}
	netRequest.Success(c, gin.H{"categoryList": list})
}
