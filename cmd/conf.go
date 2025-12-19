package cmd

import (
	"flow-blog/pkg/utils"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// InitConfig 初始化读取配置文件
func InitConfig() {

	if viper.GetString("model.env") != "production" {
		//本地开发用
		_ = godotenv.Load("./config/.env.dev")
	}

	viper.SetDefault("server.port", 8080)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	if viperReadErr := viper.ReadInConfig(); viperReadErr != nil {
		utils.RecordError("加载配置文件失败:", viperReadErr)
	}
	fmt.Println("配置读取成功...")
}
