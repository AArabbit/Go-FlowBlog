package service

import (
	"errors"
	"flow-blog/internal/api/dto"
	"flow-blog/internal/global"
	"flow-blog/internal/model"
)

// AddBookmarkData 添加收藏处理
func AddBookmarkData(d *dto.BookmarkStatus) error {
	addMark := model.BookmarkModel{
		UserID: uint(d.UserId),
		PostID: uint(d.PostId),
	}
	err := global.DB.Create(&addMark).Error
	if err != nil {
		return errors.New("添加失败")
	}
	return nil
}

// DeleteBookmarkData 取消收藏处理
func DeleteBookmarkData(d *dto.BookmarkStatus) error {
	err := global.DB.Where("user_id = ?", d.UserId).Where("post_id = ?", d.PostId).
		Delete(&model.BookmarkModel{}).Error
	if err != nil {
		return errors.New("删除失败")
	}
	return nil
}
