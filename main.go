package main

import (
	"flow-blog/cmd"
	"flow-blog/internal/api/router"
	"flow-blog/internal/global"
	"flow-blog/internal/initialize"
	"flow-blog/pkg/utils"
)

func main() {
	// 日志初始化
	global.Log = initialize.InitLogger()

	// 初始化配置项
	cmd.InitConfig()

	// 初始化数据库
	global.DB = initialize.InitGormDB()
	// 自动建表
	//initializeErr := global.DB.AutoMigrate(
	//	&model.UserModel{},
	//	//&model.PostModel{},
	//	//&model.BookmarkModel{},
	//	//&model.CommentModel{},
	//	//&model.CategoryModel{},
	//	//&model.DocModel{},
	//	//&model.DocCategoriesModel{},
	//	//&model.Visitor{},
	//)
	//if initializeErr != nil {
	//	fmt.Println("自动创建表失败..", initializeErr)
	//	return
	//}
	//fmt.Println("自动创建表成功...")

	// 初始化redis
	global.RedisClient = initialize.InitRedis()

	// 初始化ip搜索库
	utils.InitIPDB("./config/ip2region_v4.xdb")

	// 启动任务调度
	c := initialize.InitTaskScheduling()
	defer c.Stop()

	// 初始化路由
	router.InitRouter()
}
