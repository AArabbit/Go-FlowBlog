package service

import (
	"errors"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type QueryCallback func(db *gorm.DB) *gorm.DB

// Paginate 分页公共方法
func Paginate[T any](baseDB *gorm.DB, page, pageSize int,
	listOps ...QueryCallback) ([]T, int64, bool, error) {
	var (
		list  []T
		total int64
		g     errgroup.Group
	)

	// 规范化分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 12
	}

	// 协程1, 总数
	g.Go(func() error {
		return baseDB.Count(&total).Error
	})

	// 协程2, 列表
	g.Go(func() error {
		// 复制一个DB会话
		q := baseDB.Session(&gorm.Session{})

		for _, queryCallback := range listOps {
			// 处理其他条件查询
			q = queryCallback(q)
		}

		// 分页
		offset := (page - 1) * pageSize
		return q.Offset(offset).Limit(pageSize).Find(&list).Error
	})

	// 等待协程
	if err := g.Wait(); err != nil {
		return nil, 0, false, errors.New("查询失败: " + err.Error())
	}

	// true,还有更多数据
	hasMore := total > int64(page*pageSize)

	return list, total, hasMore, nil
}
