package controller

import (
	"flow-blog/internal/model/request"
	"flow-blog/internal/service"
	"flow-blog/pkg/netRequest"

	"github.com/gin-gonic/gin"
)

// AddBookmark 添加收藏
func AddBookmark(c *gin.Context) {
	var AddBookmarks request.BookmarkStatus
	if err := c.ShouldBindJSON(&AddBookmarks); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	err := service.AddBookmarkData(&AddBookmarks)
	if err != nil {
		netRequest.Fail(c, netRequest.DBError, err.Error())
		return
	}
	netRequest.Success(c, gin.H{"msg": "收藏成功"})
}

// DeleteBookmark 取消收藏
func DeleteBookmark(c *gin.Context) {
	var deleteBookmarks request.BookmarkStatus
	if err := c.ShouldBindJSON(&deleteBookmarks); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	err := service.DeleteBookmarkData(&deleteBookmarks)
	if err != nil {
		netRequest.Fail(c, netRequest.DBError, err.Error())
		return
	}
	netRequest.Success(c, gin.H{"msg": "删除成功"})
}
