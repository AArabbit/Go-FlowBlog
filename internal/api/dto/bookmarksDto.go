package dto

// 接收参数结构体

// BookmarkStatus 收藏状态
type BookmarkStatus struct {
	UserId int `json:"user_id,omitempty" binding:"required"`
	PostId int `json:"post_id,omitempty" binding:"required"`
}
