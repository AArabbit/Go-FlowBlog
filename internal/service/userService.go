package service

import (
	"crypto/tls"
	"errors"
	"flow-blog/internal/api/dto"
	"flow-blog/internal/global"
	"flow-blog/internal/model"
	"flow-blog/pkg/utils"
	"fmt"
	"math/rand"
	"net/smtp"
	"strings"
	"time"

	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
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
	auth := smtp.PlainAuth("", sendEmail, emailCode, "smtp.qq.com")
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.qq.com",
	}
	err := e.SendWithTLS("smtp.qq.com:465", auth, tlsConfig)
	if err != nil {
		if strings.Contains(err.Error(), "short response") {
			return sendNum, nil
		}
		fmt.Println("邮件发送失败:", err)
		return "", errors.New("邮件发送错误")
	}
	return sendNum, nil
}

// ValidateCode 校验验证码与邮箱是否存在，增加数据
func ValidateCode(r *dto.UserRegister) error {
	var user model.UserModel
	err := global.DB.Where("email = ?", r.Email).First(&user).Error
	if err == nil {
		return errors.New("邮箱已存在")
	}

	err = global.DB.Where("username = ?", r.UserName).First(&user).Error
	if err == nil {
		return errors.New("用户名已存在")
	}

	sendNum, sendErr := utils.RedisGet(global.ValidateRedisKey)
	if sendErr != nil {
		return err
	}
	if sendNum != r.Code {
		return errors.New("验证码过期或错误")
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
func RegisterUserDB(u *dto.UserRegister) error {
	// 插入数据库
	user := model.UserModel{
		Username:  u.UserName,
		Password:  u.PassWord,
		Email:     u.Email,
		Avatar:    "",
		Role:      "user",
		LoginType: "pwd",
	}
	// 设置随机头像
	user.Avatar = fmt.Sprintf("https://api.dicebear.com/9.x/bottts-neutral/svg?seed=%s", u.UserName)
	err := global.DB.Create(&user).Error
	if err != nil {
		global.Log.Error("用户注册保存数据库失败,")
		return err
	}
	return nil
}

// ValidateLoginInfo 登录校验信息
func ValidateLoginInfo(u *dto.UserLogin) (*model.UserModel, error) {
	var user model.UserModel
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

// GetUserInfoData ID获取用户信息
func GetUserInfoData(u *dto.UserInfo) (*model.UserModel, error) {
	var userInfo model.UserModel
	err := global.DB.Preload("Bookmarks").Where("id = ?", u.UserId).
		Find(&userInfo).Error
	if err != nil {
		return nil, errors.New("用户不存在" + err.Error())
	}
	return &userInfo, nil
}

// GetUserList 获取用户列表
func GetUserList(p *dto.PageRequest) ([]model.UserModel, int64, bool, error) {
	db := global.DB.Model(&model.UserModel{})
	listOption := func(db *gorm.DB) *gorm.DB {
		return db.Order("id ASC")
	}
	return Paginate[model.UserModel](db, p.Page, p.PageSize, listOption)
}

// UpdateUserData 更新用户信息处理
func UpdateUserData(userId string, u *dto.UpdateUser) error {
	var user = model.UserModel{
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

// UpdateUserPassData 更新用户密码
func UpdateUserPassData(u *dto.UpdatePass) error {
	sendNum, err := utils.RedisGet(global.ValidateRedisKey)
	if err != nil {
		return errors.New("验证码错误")
	}
	if sendNum != u.Code {
		return errors.New("验证码错误或已失效")
	}
	newPassword, _ := utils.Encrypt(u.Password)

	upErr := global.DB.Where("email = ?", u.Email).
		Find(&model.UserModel{}).Update("password", newPassword).Error
	if upErr != nil {
		return errors.New("更新密码失败" + upErr.Error())
	}
	return nil
}

// DeleteUserData 删除用户
func DeleteUserData(userId string) error {
	var user model.UserModel
	err := global.DB.Where("id = ?", userId).Delete(&user).Error
	if err != nil {
		return errors.New("删除失败:" + err.Error())
	}
	return nil
}
