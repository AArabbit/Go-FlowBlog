package controller

import (
	"flow-blog/internal/model"
	"flow-blog/internal/model/request"
	"flow-blog/internal/service"
	"flow-blog/pkg/netRequest"

	"github.com/gin-gonic/gin"
)

// PostsComments 文章评论列表
func PostsComments(c *gin.Context) {
	var comment request.CommentsList
	if err := c.ShouldBindJSON(&comment); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	comments, err := service.GetPostsComments(comment.PostsId)
	if err != nil {
		netRequest.Fail(c, netRequest.ServerError, err.Error())
		return
	}
	netRequest.Success(c, gin.H{
		"comments": comments,
	})
}

// AddComment 添加评论
func AddComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	err := service.AddCommentData(&comment)
	if err != nil {
		netRequest.Fail(c, netRequest.ServerError, err.Error())
		return
	}
	netRequest.Success(c, gin.H{"msg": "评论成功"})
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	var deleteId request.DeleteComment
	if err := c.ShouldBindJSON(&deleteId); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	err := service.DeleteCommentData(&deleteId)
	if err != nil {
		netRequest.Fail(c, netRequest.ServerError, err.Error())
		return
	}
	netRequest.Success(c, gin.H{"msg": "删除成功"})
}
