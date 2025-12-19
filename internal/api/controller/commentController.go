package controller

import (
	"flow-blog/internal/api/dto"
	"flow-blog/internal/model"
	"flow-blog/internal/service"
	"flow-blog/pkg/app"
	"flow-blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

// NewCommentsController 构造函数
func NewCommentsController() *CommentController {
	return &CommentController{}
}

// PostsComments 文章评论列表
func (*CommentController) PostsComments(c *gin.Context) {
	var comment dto.CommentsList
	if err := c.ShouldBindJSON(&comment); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	comments, err := service.GetPostsComments(comment.PostsId)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{
		"comments": comments,
	})
}

// AddComment 添加评论
func (*CommentController) AddComment(c *gin.Context) {
	var comment model.CommentModel
	if err := c.ShouldBindJSON(&comment); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	err := service.AddCommentData(&comment)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{"msg": "评论成功"})
}

// DeleteComment 删除评论
func (*CommentController) DeleteComment(c *gin.Context) {
	var deleteId dto.DeleteComment
	if err := c.ShouldBindJSON(&deleteId); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	err := service.DeleteCommentData(&deleteId)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{"msg": "删除成功"})
}
