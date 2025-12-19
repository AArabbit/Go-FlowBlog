package service

import (
	"flow-blog/internal/api/dto"
	"flow-blog/internal/global"
	"flow-blog/internal/model"
	"flow-blog/pkg/utils"

	"gorm.io/gorm"
)

func VisitorTrafficData(ip, userAgent string) error {
	var visitor model.Visitor

	// 获取地理位置
	location := utils.GetLocation(ip)

	// IP是否已存在
	if result := global.DB.Where("ip = ?", ip).First(&visitor); result.Error != nil {
		// 新访客
		visitor = model.Visitor{
			IP:        ip,
			Count:     1,
			UserAgent: userAgent,
			Location:  location,
		}
		if err := global.DB.Create(&visitor).Error; err != nil {
			return err
		}
		return nil
	} else {
		// 老访客，次数+1，更新时间
		visitor.Count++
		visitor.UserAgent = userAgent // 更新设备信息
		visitor.Location = location
		if err := global.DB.Save(&visitor).Error; err != nil {
			return err
		}
		return nil
	}
}

// GetVisitorTrafficData 获取访客列表
func GetVisitorTrafficData(p *dto.PageRequest) ([]model.Visitor, int64, bool, error) {
	db := global.DB.Model(&model.Visitor{})
	listOption := func(db *gorm.DB) *gorm.DB {
		return db.Order("updated_at DESC")
	}
	return Paginate[model.Visitor](db, p.Page, p.PageSize, listOption)
}

// DeleteVisitorTrafficData 删除记录
func DeleteVisitorTrafficData(id string) error {
	err := global.DB.Where("id = ?", id).Delete(&model.Visitor{}).Error
	if err != nil {
		return err
	}
	return nil
}
