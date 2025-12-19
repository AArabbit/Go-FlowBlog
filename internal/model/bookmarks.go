package model

import "time"

// BookmarkModel 收藏表 - 多对多关联
type BookmarkModel struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement:false" json:"-"`             // 用户ID
	PostID    uint      `gorm:"primaryKey;autoIncrement:false;index" json:"post_id"` // 文章ID
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`                             // 收藏时间
}

// TableName 指定表名
func (BookmarkModel) TableName() string {
	return "bookmarks"
}
