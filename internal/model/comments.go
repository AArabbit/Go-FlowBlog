package model

import (
	"time"

	"gorm.io/gorm"
)

// CommentModel 评论表
type CommentModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"` // 评论ID
	PostID    uint           `gorm:"not null;index" json:"post_id"`      // 关联文章ID
	UserID    uint           `gorm:"not null;index" json:"user_id"`      // 评论者ID
	UserName  string         `gorm:"not null;index" json:"user_name"`    // 评论者名称
	Avatar    string         `gorm:"not null;index" json:"avatar"`       // 评论者头像url
	Content   string         `gorm:"type:text;not null" json:"content"`  // 评论内容
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`   // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`   // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                     // 软删除时间
}

// TableName 指定表名
func (CommentModel) TableName() string {
	return "comments"
}
