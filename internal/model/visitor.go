package model

import "time"

type Visitor struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	IP        string    `gorm:"type:varchar(50);uniqueIndex" json:"ip"` // IP地址
	Location  string    `gorm:"type:varchar(100)" json:"location"`      // 地理位置
	Count     int       `gorm:"default:1" json:"count"`                 // 访问次数
	UserAgent string    `gorm:"type:varchar(255)" json:"user_agent"`    // 设备信息
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (*Visitor) TableName() string {
	return "visitors"
}
