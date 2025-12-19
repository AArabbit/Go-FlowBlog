package service

import (
	"errors"
	"flow-blog/internal/global"
	"flow-blog/internal/model"
	"flow-blog/internal/model/request"
	"flow-blog/pkg/utils"
	"fmt"
	"math/rand"
	"net/smtp"
	"strings"
	"time"

	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// SendEmailValidate 发送邮箱验证码
func SendEmailValidate(recipientEmail []string) (string, error) {
	e := email.NewEmail()
	blogName := viper.GetString("app.name")
	sendEmail := viper.GetString("email.sendEmail")
	emailCode := viper.GetString("email.emailCode")

	//发件人
	e.From = fmt.Sprintf("%s<%s>", blogName, sendEmail)
	// 收件人
	e.To = recipientEmail
	// 标题
	e.Subject = blogName + "注册验证"

	// 生成六位随机数
	rndNum := rand.New(rand.NewSource(time.Now().UnixNano()))
	sendNum := fmt.Sprintf("%06v", rndNum.Int31n(1000000))

	// 发送时间
	sendTime := time.Now().Format("2006-01-02 15:04:05")
	currentYear := time.Now().Year()

	sendContent := fmt.Sprintf(utils.HTMLTemplate, blogName, sendNum,
		recipientEmail[0], sendTime, currentYear, blogName)

	e.HTML = []byte(sendContent)

	// 邮件服务器配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", sendEmail, emailCode, "smtp.qq.com"))
	if !strings.Contains(err.Error(), "short response") {
		return "", errors.New("邮件发送错误")
	}
	return sendNum, nil
}

// ValidateCode 校验验证码与邮箱是否存在，增加数据
func ValidateCode(r *request.UserRegister) error {
	var user model.User
	err := global.DB.Where("email = ?", r.Email).First(&user).Error
	if err == nil {
		return errors.New("邮箱已存在")
	}

	sendNum, err := utils.RedisGet(global.ValidateRedisKey)
	if err != nil {
		return err
	}
	if sendNum != r.Code {
		return errors.New("验证码过期")
	}
	// 存到数据库
	if DbErr := RegisterUserDB(r); DbErr != nil {
		return err
	}
	return nil
}

// ValidateCodeRedis redis缓存验证码
func ValidateCodeRedis(sendNum string) error {
	// 验证码存入redis, 5分钟过期
	redisErr := utils.RedisSet(global.ValidateRedisKey, sendNum, 5*time.Minute)
	if redisErr != nil {
		return redisErr
	}
	return nil
}

// RegisterUserDB 用户注册存储
func RegisterUserDB(u *request.UserRegister) error {
	// 插入数据库
	user := model.User{
		Username: u.UserName,
		Password: u.PassWord,
		Email:    u.Email,
		Avatar:   "",
		Role:     "user",
	}
	// 设置随机头像
	user.Avatar = fmt.Sprintf("https://api.dicebear.com/7.x/avataaars/svg?seed=%s", u.UserName)
	err := global.DB.Create(&user).Error
	if err != nil {
		global.Log.Error("用户注册保存数据库失败,")
		return err
	}
	return nil
}

// ValidateLoginInfo 登录校验信息
func ValidateLoginInfo(u *request.UserLogin) (*model.User, error) {
	var user model.User
	err := global.DB.Preload("Bookmarks").Where("username = ?", u.UserName).First(&user).Error
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	if !utils.Decrypt(user.Password, u.PassWord) {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

// LoginToken 生成token
func LoginToken(userId uint, userName string) (string, string, error) {
	// 生成双 Token
	accessToken, accErr := utils.GenerateAccessToken(userId, userName)
	refreshToken, refErr := utils.GenerateRefreshToken(userId, userName)
	if accErr != nil || refErr != nil {
		return "", "", errors.New("token获取失败")
	}
	return accessToken, refreshToken, nil
}

// GetUserList 获取用户列表
func GetUserList(p *request.PostsPageRequest) ([]model.User, int64, bool, error) {
	var (
		user  []model.User
		total int64
		g     errgroup.Group
	)

	g.Go(func() error {
		// 总条数
		return global.DB.Model(&model.User{}).Count(&total).Error
	})
	g.Go(func() error {
		return global.DB.Order("id DESC").Offset((p.Page - 1) * p.PageSize).
			Limit(p.PageSize).Find(&user).Error
	})
	// 等待协程
	if err := g.Wait(); err != nil {
		return nil, 0, false, errors.New("查询失败: " + err.Error())
	}
	isEnd := total <= int64(p.Page*p.PageSize)
	return user, total, !isEnd, nil
}

// UpdateUserData 更新用户信息处理
func UpdateUserData(userId string, u *request.UpdateUser) error {
	var user = model.User{
		Username: u.UserName,
		Email:    u.Email,
		Role:     u.Role,
		Avatar:   u.Avatar,
	}
	err := global.DB.Where("id = ?", userId).Updates(&user).Error
	if err != nil {
		return errors.New("更新失败:" + err.Error())
	}
	return nil
}

// DeleteUserData 删除用户
func DeleteUserData(userId string) error {
	var user model.User
	err := global.DB.Where("id = ?", userId).Delete(&user).Error
	if err != nil {
		return errors.New("删除失败:" + err.Error())
	}
	return nil
}
