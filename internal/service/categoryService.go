package service

import (
	"flow-blog/internal/api/dto"
	"flow-blog/internal/global"
	"flow-blog/internal/model"

	"gorm.io/gorm"
)

// GetCategoryList 查询分类列表
func GetCategoryList() ([]model.CategoryModel, error) {
	var categories []model.CategoryModel

	err := global.DB.Find(&categories).Error
	if err != nil {
		return []model.CategoryModel{}, err
	}
	return categories, nil
}

// GetCategoryPostList 分类id获取文章
func GetCategoryPostList(p *dto.CategoryPost) ([]model.PostModel, int64, bool, error) {
	db := global.DB.Model(&model.PostModel{})
	db = db.Where("category_id = ?", p.CategoryId)

	return Paginate[model.PostModel](db, p.Page, p.PageSize, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Author").Order("id ASC")
	})
}
