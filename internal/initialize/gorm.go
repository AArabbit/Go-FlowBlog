package initialize

import (
	"flow-blog/pkg/utils"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitGormDB() *gorm.DB {
	// 获取数据库配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.Get("mysql.user"),
		viper.GetString("mysql.password"),
		viper.Get("mysql.host"),
		viper.Get("mysql.dataDB"),
		viper.Get("mysql.charset"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true, // 忽略日志ErrRecordNotFound错误
		},
	)

	var db *gorm.DB
	var err error
	maxRetries := 6
	// 尝试连接6 次，每次间隔2秒，总共等待12秒
	for i := 1; i <= maxRetries; i++ {
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:               dsn,
			DefaultStringSize: 171, // utf8mb4默认长度
		}), &gorm.Config{
			Logger:                 newLogger,
			SkipDefaultTransaction: false, // 不使用默认事务
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false, // 使用复数表名
			},
			DisableForeignKeyConstraintWhenMigrating: true, // 禁止生成实体外键
		})
		if err == nil {
			fmt.Printf("第%d次尝试，数据库连接成功..\n", i)
			break
		}
		fmt.Printf("第%d次尝试，连接失败:%v, 2秒后重试...\n", i, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		utils.RecordError("数据库尝试连接失败：", err)
	}

	return db
}
