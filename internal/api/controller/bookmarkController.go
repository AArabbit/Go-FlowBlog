package controller

import (
	"flow-blog/internal/api/dto"
	"flow-blog/internal/service"
	"flow-blog/pkg/app"
	"flow-blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type BookmarkController struct{}

// NewBookmarkController 构造函数
func NewBookmarkController() *BookmarkController {
	return &BookmarkController{}
}

// AddBookmark 添加收藏
func (*BookmarkController) AddBookmark(c *gin.Context) {
	var AddBookmarks dto.BookmarkStatus
	if err := c.ShouldBindJSON(&AddBookmarks); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	err := service.AddBookmarkData(&AddBookmarks)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{"msg": "收藏成功"})
}

// DeleteBookmark 取消收藏
func (*BookmarkController) DeleteBookmark(c *gin.Context) {
	var deleteBookmarks dto.BookmarkStatus
	if err := c.ShouldBindJSON(&deleteBookmarks); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	err := service.DeleteBookmarkData(&deleteBookmarks)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{"msg": "删除成功"})
}
