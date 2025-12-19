package dto

// 接收参数结构体

// CategoryPost 搜索文章
type CategoryPost struct {
	CategoryId int `json:"category_id,omitempty" binding:"required"`
	PageRequest
}
