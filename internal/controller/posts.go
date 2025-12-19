package controller

import (
	"flow-blog/internal/model"
	"flow-blog/internal/model/request"
	"flow-blog/internal/service"
	"flow-blog/pkg/netRequest"

	"github.com/gin-gonic/gin"
)

// PostsList 获取文章列表
func PostsList(c *gin.Context) {
	var postPage request.PostsPageRequest
	if err := c.ShouldBindJSON(&postPage); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}

	list, total, isEnd, err := service.GetPostsList(&postPage)
	if err != nil {
		netRequest.Fail(c, netRequest.DBError, "数据库错误")
		return
	}
	netRequest.Success(c, gin.H{
		"postList": list,
		"has_more": isEnd,
		"total":    total,
	})
}

// PostsDetail 获取文章详情
func PostsDetail(c *gin.Context) {
	postId := c.Param("id")
	postsDetail, err := service.GetPostsDetail(postId)
	if err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	netRequest.Success(c, gin.H{"postsDetail": postsDetail})
}

// PostsViews 更新文章浏览量
func PostsViews(c *gin.Context) {
	var views request.UpdateViews
	if err := c.ShouldBindJSON(&views); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	err := service.UpdatePostsViews(views.PostsId, views.Views)
	if err != nil {
		netRequest.Fail(c, netRequest.NotFound, "没找到文章")
		return
	}
	netRequest.Success(c, gin.H{"msg": "更新成功"})
}

// AddPosts 添加文章
func AddPosts(c *gin.Context) {
	var posts model.Post
	if err := c.ShouldBindJSON(&posts); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	err := service.AddPostsData(&posts)
	if err != nil {
		netRequest.Fail(c, netRequest.DBError, "数据库错误")
		return
	}
	netRequest.Success(c, gin.H{
		"msg": "添加成功",
	})
}

// UpdatePosts 更新文章
func UpdatePosts(c *gin.Context) {
	var posts model.Post
	postId := c.Param("id")
	if err := c.ShouldBindJSON(&posts); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	err := service.UpdatePostsData(&posts, postId)
	if err != nil {
		netRequest.Fail(c, netRequest.DBError, "数据库错误")
		return
	}
	netRequest.Success(c, gin.H{
		"msg": "更新成功",
	})
}

// DeletePosts 删除文章
func DeletePosts(c *gin.Context) {
	postId := c.Param("id")
	err := service.DeletePostsData(postId)
	if err != nil {
		netRequest.Fail(c, netRequest.DBError, err.Error())
		return
	}
	netRequest.Success(c, gin.H{"msg": "删除成功"})
}

// SearchPosts 搜索文章
func SearchPosts(c *gin.Context) {
	var key request.SearchPosts
	if err := c.ShouldBindJSON(&key); err != nil {
		netRequest.Fail(c, netRequest.ParamError, "参数错误")
		return
	}
	posts, total, isEnd, err := service.SearchPostsData(&key)
	if err != nil {
		netRequest.Fail(c, netRequest.DBError, err.Error())
		return
	}

	netRequest.Success(c, gin.H{
		"total":    total,
		"has_more": isEnd,
		"search":   posts,
	})
}
