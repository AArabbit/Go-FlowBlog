package model

import (
	"time"

	"gorm.io/gorm"
)

// CategoryModel 分类表
type CategoryModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`                // 分类ID
	Name      string         `gorm:"type:varchar(64);not null" json:"name"`             // 分类名称
	Slug      string         `gorm:"type:varchar(64);uniqueIndex;not null" json:"slug"` // 分类别名(用于URL)
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`                  // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`                  // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                    // 软删除时间

	// 关联关系
	Posts []PostModel `gorm:"foreignKey:CategoryID" json:"-"` // 该分类下的文章
}

// TableName 指定表名
func (CategoryModel) TableName() string {
	return "categories"
}
