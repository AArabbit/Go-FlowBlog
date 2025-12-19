package model

import "time"

type DocCategoriesModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`

	Docs []DocModel `gorm:"foreignKey:CategoryID" json:"docs"`
}

// TableName 指定表名
func (*DocCategoriesModel) TableName() string {
	return "doc_categories"
}
