package initialize

import (
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger 初始化日志模块
func InitLogger() *zap.SugaredLogger {
	// 设置日志等级
	level := zapcore.DebugLevel // 开发环境使用
	if viper.GetString("model.env") == "production" {
		level = zapcore.InfoLevel // 生产环境使用
	}

	// 创建 core
	core := zapcore.NewCore(getEncoder(), getLogWriterSync(), level)

	return zap.New(core).Sugar()
}

// 自定义配置项
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()       // 拿到配置项实例
	encoderConfig.TimeKey = "time"                          // 修改time字段名
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 把日志级别名称大写
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime)) // 时间格式
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriterSync() zapcore.WriteSyncer {
	const logDirName = "logs"
	rootDir, _ := os.Getwd()
	logDirPath := filepath.Join(rootDir, logDirName)

	// 创建logs目录
	if _, err := os.Stat(logDirPath); os.IsNotExist(err) {
		// 权限设为0755
		_ = os.Mkdir(logDirPath, 0755)
	}

	// 拼接文件名
	logFilePath := filepath.Join(logDirPath, time.Now().Format(time.DateOnly)+".txt")

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"),
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"),
		Compress:   false,
	}

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
}
