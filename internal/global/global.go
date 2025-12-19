package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量
var (
	Log         *zap.SugaredLogger
	DB          *gorm.DB
	RedisClient *redis.Client
)

// ValidateRedisKey 邮箱验证码redisKey
const ValidateRedisKey = "send_num"

// DailyRedisKey 每日推荐文章redisKey
const DailyRedisKey = "midnight_posts"

// PostsDraftKey 草稿通用key
const PostsDraftKey = "draft_posts_"

// ThirdStringKey 第三方登录随机字符串key
const ThirdStringKey = "oauth_key_"
