package request

// BookmarkStatus 收藏状态
type BookmarkStatus struct {
	UserId int `json:"user_id,omitempty"`
	PostId int `json:"post_id,omitempty"`
}
