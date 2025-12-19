package request

// PostsPageRequest 分页查询参数结构
type PostsPageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// UpdateViews 更新文章浏览量请求体
type UpdateViews struct {
	PostsId int `json:"posts_id,omitempty"`
	Views   int `json:"views,omitempty"`
}

// CommentsList 评论列表参数
type CommentsList struct {
	PostsId int `json:"posts_id,omitempty"`
}

// DeleteComment 删除评论参数
type DeleteComment struct {
	CommentId int `json:"comment_id,omitempty"`
}

// SearchPosts 搜索文章
type SearchPosts struct {
	Keyword string `json:"keyword,omitempty"`
	PostsPageRequest
}
