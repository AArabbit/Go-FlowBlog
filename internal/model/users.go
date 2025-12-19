package model

import (
	"errors"
	"flow-blog/pkg/utils"
	"time"

	"gorm.io/gorm"
)

// UserModel 用户表
type UserModel struct {
	ID        uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                                // 用户ID
	Username  string `gorm:"column:username;type:varchar(64);uniqueIndex;not null" json:"username"`       // 用户名
	Password  string `gorm:"column:password;type:varchar(128);not null" json:"-"`                         // 加密后的密码
	Email     string `gorm:"column:email;type:varchar(128);uniqueIndex;not null" json:"email"`            // 邮箱地址
	Avatar    string `gorm:"column:avatar;type:varchar(255);default:''" json:"avatar"`                    // 头像URL
	Role      string `gorm:"column:role;type:enum('admin','user');default:'user'" json:"role"`            // 角色: admin-管理员, user-普通用户
	LoginType string `gorm:"column:login_type;type:enum('pwd','github');default:'pwd'" json:"login_type"` // 登录方式
	GithubID  string `gorm:"column:github_id;type:varchar(64);index" json:"github_id"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"` // 更新时间

	// 关联关系
	Posts     []PostModel     `gorm:"foreignKey:UserID" json:"-"`         // 用户发布的文章
	Comments  []CommentModel  `gorm:"foreignKey:UserID" json:"-"`         // 用户的评论
	Bookmarks []BookmarkModel `gorm:"foreignKey:UserID" json:"bookmarks"` // 用户的收藏
}

// TableName 指定表名
func (u *UserModel) TableName() string {
	return "users"
}

// BeforeCreate 加密密码
func (u *UserModel) BeforeCreate(g *gorm.DB) error {
	if u.LoginType == "pwd" {
		hashPassword, err := utils.Encrypt(u.Password)
		u.Password = hashPassword
		if err != nil {
			return errors.New("加密失败")
		}
		return nil
	}
	return nil
}
