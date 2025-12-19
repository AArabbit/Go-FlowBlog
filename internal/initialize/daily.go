package initialize

import (
	"encoding/json"
	"flow-blog/internal/global"
	"flow-blog/internal/service"
	"flow-blog/pkg/utils"
	"fmt"

	"github.com/robfig/cron/v3"
)

// InitTaskScheduling 创建任务调度器
func InitTaskScheduling() *cron.Cron {
	// 初始执行一次
	getDailyPosts()
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("@daily", func() {
		// 每天凌晨执行一次
		getDailyPosts()
	})

	c.Start()

	if err != nil {
		utils.RecordError("定时任务添加失败", err)
	}

	return c
}

// 初始执行的数据库查询
func getDailyPosts() {
	posts, err := service.DailyPostsDetail()
	if err != nil {
		// 不停止程序，继续运行
		fmt.Println(err.Error())
	}
	postJson, _ := json.Marshal(posts)
	_ = utils.RedisSet(global.DailyRedisKey, postJson, 0)
}
