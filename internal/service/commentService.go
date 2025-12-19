package service

import (
	"errors"
	"flow-blog/internal/api/dto"
	"flow-blog/internal/global"
	"flow-blog/internal/model"
)

// GetPostsComments 获取文章的评论列表
func GetPostsComments(postId int) ([]model.CommentModel, error) {
	var comments []model.CommentModel
	err := global.DB.Where("post_id = ?", postId).Find(&comments).Error
	if err != nil {
		return nil, errors.New("数据库查询错误")
	}

	return comments, nil
}

// AddCommentData 添加评论
func AddCommentData(c *model.CommentModel) error {
	err := global.DB.Create(c).Error
	if err != nil {
		return errors.New("评论失败:数据库错误")
	}
	return nil
}

// DeleteCommentData 删除评论
func DeleteCommentData(d *dto.DeleteComment) error {
	err := global.DB.Delete(&model.CommentModel{}, d.CommentId).Error
	if err != nil {
		return errors.New("删除失败")
	}
	return nil
}
