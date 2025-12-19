package service

import (
	"flow-blog/internal/global"
	"flow-blog/internal/model"
)

// GetCategoryList 查询分类列表
func GetCategoryList() ([]model.Category, error) {
	var categories []model.Category

	err := global.DB.Find(&categories).Error
	if err != nil {
		return []model.Category{}, err
	}
	return categories, nil
}
