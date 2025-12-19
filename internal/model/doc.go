package model

import "time"

// DocModel 文档链接
type DocModel struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string    `gorm:"type:varchar(255);not null" json:"title"`
	Category   string    `gorm:"type:varchar(255);not null" json:"category"`
	CategoryID uint      `gorm:"not null;index" json:"category_id"`
	Url        string    `gorm:"type:varchar(255);not null" json:"url"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"-"`
}

// TableName 指定表名
func (*DocModel) TableName() string {
	return "docs"
}
