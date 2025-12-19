package dto

// 接收参数结构体

// PageRequest 分页查询参数结构
type PageRequest struct {
	Page     int `json:"page,omitempty" binding:"required"`
	PageSize int `json:"page_size,omitempty" binding:"required"`
}

// CommentsList 评论列表参数
type CommentsList struct {
	PostsId int `json:"posts_id,omitempty" binding:"required"`
}

// UpdateViews 更新文章浏览量请求体
type UpdateViews struct {
	CommentsList
	Views int `json:"views,omitempty" binding:"required"`
}

// DraftPost 文章缓存参数
type DraftPost struct {
	Draft  string `json:"draft" binding:"required"`
	UserId int    `json:"userId,omitempty" binding:"required"`
}

// DeleteComment 删除评论参数
type DeleteComment struct {
	CommentId int `json:"comment_id,omitempty" binding:"required"`
}

// SearchPosts 搜索文章
type SearchPosts struct {
	Keyword string `json:"keyword,omitempty" binding:"required"`
	PageRequest
}
