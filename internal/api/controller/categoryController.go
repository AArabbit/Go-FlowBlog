package controller

import (
	"flow-blog/internal/api/dto"
	"flow-blog/internal/service"
	"flow-blog/pkg/app"
	"flow-blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

// NewCategoryController 构造函数
func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

// CategoryList 查询分类列表
func (*CategoryController) CategoryList(c *gin.Context) {
	list, err := service.GetCategoryList()
	if err != nil {
		app.FailWithMsg(c, errcode.NotFound, err.Error())
		return
	}
	app.Success(c, gin.H{"categoryList": list})
}

// CategoryPostList 根据分类ID分页查询文章列表
func (*CategoryController) CategoryPostList(c *gin.Context) {
	var param dto.CategoryPost
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}

	list, total, isEnd, err := service.GetCategoryPostList(&param)
	if err != nil {
		app.FailWithMsg(c, errcode.NotFound, err.Error())
		return
	}
	app.Success(c, gin.H{
		"postList": list,
		"has_more": isEnd,
		"total":    total,
	})
}
