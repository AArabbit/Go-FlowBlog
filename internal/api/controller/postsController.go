package controller

import (
	"flow-blog/internal/api/dto"
	"flow-blog/internal/model"
	"flow-blog/internal/service"
	"flow-blog/pkg/app"
	"flow-blog/pkg/errcode"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PostsController struct{}

// NewPostsController 构造函数
func NewPostsController() *PostsController {
	return &PostsController{}
}

// PostsList 获取文章列表
func (*PostsController) PostsList(c *gin.Context) {
	var postPage dto.PageRequest
	if err := c.ShouldBindJSON(&postPage); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}

	list, total, isEnd, err := service.GetPostsList(&postPage)
	fmt.Println(err, "err")
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{
		"postList": list,
		"has_more": isEnd,
		"total":    total,
	})
}

// PostsDetail 获取文章详情
func (*PostsController) PostsDetail(c *gin.Context) {
	postId := c.Param("id")
	postsDetail, err := service.GetPostsDetail(postId)
	if err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	app.Success(c, gin.H{"postsDetail": postsDetail})
}

// GetDailyPosts 每日推荐文章
func (*PostsController) GetDailyPosts(c *gin.Context) {
	posts, err := service.GetDailyPostsData()
	if err != nil {
		app.FailWithMsg(c, errcode.ParamError, err.Error())
		return
	}

	app.Success(c, gin.H{"postsDaily": posts})
}

// DraftPosts 缓存文章草稿
func (*PostsController) DraftPosts(c *gin.Context) {
	var draftPost dto.DraftPost
	if err := c.ShouldBindJSON(&draftPost); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}

	err := service.DraftPostsRedis(&draftPost)
	if err != nil {
		app.FailWithMsg(c, errcode.ParamError, err.Error())
		return
	}
	app.Success(c, gin.H{"msg": "保存草稿成功"})
}

// GetDraftPosts 获取缓存的草稿
func (*PostsController) GetDraftPosts(c *gin.Context) {
	userId := c.Param("id")
	data, err := service.GetDraftPostsData(userId)
	if err != nil {
		app.FailWithMsg(c, errcode.NotFound, err.Error())
		return
	}

	app.Success(c, gin.H{"draft": data})
}

// PostsViews 更新文章浏览量
func (*PostsController) PostsViews(c *gin.Context) {
	var views dto.UpdateViews
	if err := c.ShouldBindJSON(&views); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	err := service.UpdatePostsViews(views.PostsId, views.Views)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{"msg": "更新成功"})
}

// AddPosts 添加文章
func (*PostsController) AddPosts(c *gin.Context) {
	var posts model.PostModel
	if err := c.ShouldBindJSON(&posts); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	err := service.AddPostsData(&posts)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{
		"msg": "添加成功",
	})
}

// UpdatePosts 更新文章
func (*PostsController) UpdatePosts(c *gin.Context) {
	var posts model.PostModel
	postId := c.Param("id")
	if err := c.ShouldBindJSON(&posts); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	err := service.UpdatePostsData(&posts, postId)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{
		"msg": "更新成功",
	})
}

// DeletePosts 删除文章
func (*PostsController) DeletePosts(c *gin.Context) {
	postId := c.Param("id")
	err := service.DeletePostsData(postId)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}
	app.Success(c, gin.H{"msg": "删除成功"})
}

// SearchPosts 搜索文章
func (*PostsController) SearchPosts(c *gin.Context) {
	var key dto.SearchPosts
	if err := c.ShouldBindJSON(&key); err != nil {
		app.Fail(c, errcode.ParamError)
		return
	}
	posts, total, isEnd, err := service.SearchPostsData(&key)
	if err != nil {
		app.FailWithMsg(c, errcode.DBError, err.Error())
		return
	}

	app.Success(c, gin.H{
		"total":    total,
		"has_more": isEnd,
		"postList": posts,
	})
}
