package service

import (
	"errors"
	"flow-blog/internal/global"
	"flow-blog/internal/model"
)

// GetDocListData 查询常用文档列表
func GetDocListData() ([]model.DocCategoriesModel, error) {
	var docList []model.DocCategoriesModel
	err := global.DB.Model(&model.DocCategoriesModel{}).Preload("Docs").Find(&docList).Error

	if err != nil {
		return nil, errors.New("查询错误")
	}
	return docList, nil
}
