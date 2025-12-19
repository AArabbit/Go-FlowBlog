package service

import (
	"errors"
	"flow-blog/internal/global"
	"flow-blog/internal/model"
	"flow-blog/internal/model/request"
)

// GetPostsComments 获取文章的评论列表
func GetPostsComments(postId int) ([]model.Comment, error) {
	var comments []model.Comment
	err := global.DB.Where("post_id = ?", postId).Find(&comments).Error
	if err != nil {
		return nil, errors.New("数据库查询错误")
	}

	return comments, nil
}

// AddCommentData 添加评论
func AddCommentData(c *model.Comment) error {
	err := global.DB.Create(c).Error
	if err != nil {
		return errors.New("评论失败:数据库错误")
	}
	return nil
}

// DeleteCommentData 删除评论
func DeleteCommentData(d *request.DeleteComment) error {
	err := global.DB.Delete(&model.Comment{}, d.CommentId).Error
	if err != nil {
		return errors.New("删除失败")
	}
	return nil
}
