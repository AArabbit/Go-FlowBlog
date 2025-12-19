package service

import (
	"errors"
	"flow-blog/internal/global"
	"flow-blog/internal/model"
	"flow-blog/internal/model/request"

	"golang.org/x/sync/errgroup"
)

// GetPostsList 查询文章列表
func GetPostsList(p *request.PostsPageRequest) ([]model.Post, int64, bool, error) {
	var (
		posList []model.Post
		total   int64
		g       errgroup.Group
	)
	g.Go(func() error {
		// 总条数
		return global.DB.Model(&model.Post{}).Count(&total).Error
	})
	g.Go(func() error {
		// 分页
		return global.DB.Preload("Author").Order("id DESC").
			Offset((p.Page - 1) * p.PageSize).Limit(p.PageSize).Find(&posList).Error
	})
	// 等待协程
	if err := g.Wait(); err != nil {
		return nil, 0, false, errors.New("查询失败: " + err.Error())
	}

	isEnd := total <= int64(p.Page*p.PageSize)
	return posList, total, !isEnd, nil
}

// GetPostsDetail 查询文章详情
func GetPostsDetail(id string) (*model.Post, error) {
	var postsDetail model.Post
	err := global.DB.Preload("Author").Where("id = ?", id).First(&postsDetail).Error
	if err != nil {
		return nil, err
	}
	return &postsDetail, nil
}

// UpdatePostsViews 更新文章浏览量
func UpdatePostsViews(postsId int, views int) error {
	err := global.DB.Model(&model.Post{}).Where("id = ?", postsId).
		Update("views", views).Error
	if err != nil {
		return err
	}
	return nil
}

// AddPostsData 存储文章
func AddPostsData(p *model.Post) error {
	err := global.DB.Create(p).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdatePostsData 更新文章处理
func UpdatePostsData(p *model.Post, postId string) error {
	err := global.DB.Where("id = ?", postId).Updates(p).Error
	if err != nil {
		return err
	}
	return nil
}

// DeletePostsData 删除文章处理
func DeletePostsData(postId string) error {
	err := global.DB.Where("id = ?", postId).Delete(&model.Post{}).Error
	if err != nil {
		return err
	}
	return nil
}

// SearchPostsData 搜索文章
func SearchPostsData(s *request.SearchPosts) ([]model.Post, int64, bool, error) {
	var posts []model.Post
	var g errgroup.Group
	var total int64

	g.Go(func() error {
		// 总条数
		return global.DB.Model(&model.Post{}).Where("title like ?", "%"+s.Keyword+"%").
			Where("content like ?", "%"+s.Keyword+"%").Count(&total).Error
	})

	g.Go(func() error {
		return global.DB.Preload("Author").Where("title like ?", "%"+s.Keyword+"%").
			Where("content like ?", "%"+s.Keyword+"%").
			Offset((s.Page - 1) * s.PageSize).Limit(s.PageSize).Find(&posts).Error
	})

	if err := g.Wait(); err != nil {
		return nil, 0, false, errors.New("查询失败: " + err.Error())
	}

	isEnd := total <= int64(s.Page*s.PageSize)
	return posts, total, !isEnd, nil
}
