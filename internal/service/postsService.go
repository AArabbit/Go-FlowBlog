package service

import (
	"encoding/json"
	"errors"
	"flow-blog/internal/api/dto"
	"flow-blog/internal/global"
	"flow-blog/internal/model"
	"flow-blog/pkg/utils"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// GetPostsList 查询文章列表
func GetPostsList(p *dto.PageRequest) ([]model.PostModel, int64, bool, error) {
	db := global.DB.Model(&model.PostModel{})
	listOption := func(db *gorm.DB) *gorm.DB {
		return db.Preload("Author").Order("id DESC")
	}

	// 调用公共方法
	return Paginate[model.PostModel](db, p.Page, p.PageSize, listOption)
}

// GetPostsDetail 查询文章详情
func GetPostsDetail(id string) (*model.PostModel, error) {
	var postsDetail model.PostModel
	err := global.DB.Preload("Author").Where("id = ?", id).First(&postsDetail).Error
	if err != nil {
		return nil, err
	}
	return &postsDetail, nil
}

// DailyPostsDetail 数据库查询每日推荐文章
func DailyPostsDetail() (*model.PostModel, error) {
	var postList []model.PostModel
	err := global.DB.Preload("Author").Where("is_curated = ?", true).
		Find(&postList).Error
	if err != nil {
		return nil, errors.New("查询失败" + err.Error())
	}

	if len(postList) < 0 {
		return nil, errors.New("没有数据")
	}
	// 随机返回一篇精选文章
	reanNum := rand.Intn(len(postList))
	return &postList[reanNum], nil
}

// GetDailyPostsData 查询缓存的每日推荐文章
func GetDailyPostsData() (*model.PostModel, error) {
	postsJson, err := utils.RedisGet(global.DailyRedisKey)
	if err != nil {
		return nil, errors.New("没有精选文章" + err.Error())
	}
	var posts model.PostModel
	jsonErr := json.Unmarshal([]byte(postsJson), &posts)
	if jsonErr != nil {
		return nil, errors.New("缓存数据格式错误" + jsonErr.Error())
	}

	return &posts, nil
}

// DraftPostsRedis 缓存文章草稿
func DraftPostsRedis(d *dto.DraftPost) error {
	userId := global.PostsDraftKey + strconv.Itoa(d.UserId)
	draft := d.Draft
	// 缓存48小时
	err := utils.RedisSet(userId, draft, time.Hour*48)
	if err != nil {
		return errors.New("草稿保存失败" + err.Error())
	}
	return nil
}

// GetDraftPostsData 获取缓存的草稿
func GetDraftPostsData(id string) (string, error) {
	redisRes, err := utils.RedisGet(global.PostsDraftKey + id)
	if err != nil {
		return "", errors.New("草稿过期或不存在")
	}
	return redisRes, nil
}

// UpdatePostsViews 更新文章浏览量
func UpdatePostsViews(postsId int, views int) error {
	err := global.DB.Model(&model.PostModel{}).Where("id = ?", postsId).
		Update("views", views).Error
	if err != nil {
		return err
	}
	return nil
}

// AddPostsData 存储文章
func AddPostsData(p *model.PostModel) error {
	err := global.DB.Create(p).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdatePostsData 更新文章处理
func UpdatePostsData(p *model.PostModel, postId string) error {
	err := global.DB.Where("id = ?", postId).Updates(p).Error
	if err != nil {
		return err
	}
	return nil
}

// DeletePostsData 删除文章处理
func DeletePostsData(postId string) error {
	err := global.DB.Where("id = ?", postId).Delete(&model.PostModel{}).Error
	if err != nil {
		return err
	}
	return nil
}

// SearchPostsData 搜索文章
func SearchPostsData(s *dto.SearchPosts) ([]model.PostModel, int64, bool, error) {
	db := global.DB.Model(&model.PostModel{})
	db = db.Where("title LIKE ?", "%"+s.Keyword+"%").
		Where("content LIKE ?", "%"+s.Keyword+"%")

	return Paginate[model.PostModel](db, s.Page, s.PageSize, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Author").Order("id DESC")
	})
}
