package model

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

// PostModel 文章表
type PostModel struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`        // 文章ID
	Title      string    `gorm:"type:varchar(255);not null" json:"title"`   // 文章标题
	Desc       string    `gorm:"type:varchar(512);default:''" json:"desc"`  // 文章摘要
	Content    string    `gorm:"type:longtext" json:"content"`              // 文章内容(Markdown)
	Cover      string    `gorm:"type:varchar(255);default:''" json:"cover"` // 封面图片URL
	CategoryID uint      `gorm:"not null;index" json:"category_id"`         // 关联分类ID
	UserID     uint      `gorm:"not null;index" json:"user_id"`             // 作者ID
	Views      uint      `gorm:"type:int unsigned;default:0" json:"views"`  // 浏览量
	IsCurated  bool      `gorm:"type:bool;default:false" json:"is_curated"` // 是否精选
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`          // 创建时间
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`          // 更新时间

	// 关联关系
	Category  CategoryModel   `gorm:"foreignKey:CategoryID" json:"-"`  // 分类
	Author    UserModel       `gorm:"foreignKey:UserID" json:"author"` // 作者
	Comments  []CommentModel  `gorm:"foreignKey:PostID" json:"-"`      // 文章的评论
	Bookmarks []BookmarkModel `gorm:"foreignKey:PostID" json:"-"`      // 文章的收藏记录
}

// TableName 指定表名
func (*PostModel) TableName() string {
	return "posts"
}

// BeforeCreate 给数据添加随机封面图片url
func (u *PostModel) BeforeCreate(g *gorm.DB) error {
	u.Cover = fmt.Sprintf("https://picsum.photos/seed/%d/800/500", rand.Intn(200))
	return nil
}
