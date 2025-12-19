package service

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flow-blog/internal/global"
	"flow-blog/internal/model"
	"flow-blog/pkg/utils"
	"fmt"
	"io"
	"log"
	"time"

	"crypto/rand"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// GenerateRandomState 生成随机字符串，存入redis
func GenerateRandomState() (string, error) {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	state := hex.EncodeToString(b)
	redisKey := global.ThirdStringKey + state
	err := utils.RedisSet(redisKey, state, time.Minute*5)
	if err != nil {
		log.Fatalf("第三方登录redis缓存失败：%v", err.Error())
		return "", err
	}
	return state, nil
}

// GetOauthConf 创建oauth2配置
func GetOauthConf() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     viper.GetString("github.clientId"), // Client ID
		ClientSecret: viper.GetString("github.secret"),   // Client Secret
		Scopes:       []string{"user"},                   // scope为user
		Endpoint:     github.Endpoint,
		RedirectURL:  viper.GetString("github.redirectUrl"), // 回调地址，必须与GitHub设置一致
	}
}

// CompareState 校验state
func CompareState(state string) error {
	_, err := utils.RedisGet(global.ThirdStringKey + state)
	if err != nil {
		return errors.New("state验证不通过" + err.Error())
	}
	// 验证成功后删除
	err = utils.RedisDel(global.ThirdStringKey + state)
	if err != nil {
		return err
	}
	return nil
}

type GitHubEmail struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility"`
}

// GetGithubUserInfo 获取用户信息
func GetGithubUserInfo(oauthConf *oauth2.Config, token *oauth2.Token) (*model.UserModel, string, string, error) {
	client := oauthConf.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, "", "", err
	}
	defer resp.Body.Close()

	userData, _ := io.ReadAll(resp.Body)

	var userMap map[string]interface{}
	_ = json.Unmarshal(userData, &userMap)

	finalEmail := ""
	if userMap["email"] != nil {
		finalEmail = userMap["email"].(string)
	} else {
		respEmails, emailErr := client.Get("https://api.github.com/user/emails")
		if emailErr == nil {
			defer respEmails.Body.Close()
			emailsData, _ := io.ReadAll(respEmails.Body)
			var emails []GitHubEmail
			_ = json.Unmarshal(emailsData, &emails)

			for _, e := range emails {
				if e.Primary && e.Verified {
					finalEmail = e.Email
					break
				}
			}
		}
	}

	userLogin := userMap["login"].(string)
	userAvatar := userMap["avatar_url"].(string)
	ghID := userMap["id"].(float64)

	tempGithubUser := &model.UserModel{
		Username:  userLogin,
		Email:     finalEmail,
		Avatar:    userAvatar,
		Role:      "user",
		LoginType: "github",
		GithubID:  fmt.Sprintf("%.0f", ghID),
	}

	finalUser, tempErr := LoginOrRegisterGithubUser(tempGithubUser)
	if tempErr != nil {
		return nil, "", "", tempErr
	}

	accessToken, refreshToken, _ := LoginToken(finalUser.ID, finalUser.Username)
	return finalUser, accessToken, refreshToken, nil
}

// LoginOrRegisterGithubUser 处理GitHub用户的登录/注册
func LoginOrRegisterGithubUser(ghUser *model.UserModel) (*model.UserModel, error) {
	var dbUser model.UserModel
	var err error

	// GithubID查找
	err = global.DB.Where("github_id = ?", ghUser.GithubID).First(&dbUser).Error

	if err == nil {
		// -> 更新信息
		if upErr := updateGithubUser(&dbUser, ghUser); upErr != nil {
			return nil, err
		}
		// 返回最终数据
		return loadFullUserData(dbUser.ID)
	}

	// 邮箱合并通过Email查找已存在的账号
	if ghUser.Email != "" {
		errEmail := global.DB.Where("email = ?", ghUser.Email).First(&dbUser).Error

		if errEmail == nil {
			// 邮箱相同，合并账号
			dbUser.GithubID = ghUser.GithubID

			// 执行更新
			if upErr := updateGithubUser(&dbUser, ghUser); upErr != nil {
				return nil, upErr
			}
			return loadFullUserData(dbUser.ID)
		}
	}

	// 判定为新用户
	// 检查用户名冲突
	var conflictUser model.UserModel
	if dbErr := global.DB.Where("username = ?", ghUser.Username).First(&conflictUser).Error; dbErr == nil {
		ghUser.Username = fmt.Sprintf("%s%s", ghUser.Username, "(github)")
	}

	// 创建新用户
	if dbErr := global.DB.Create(ghUser).Error; dbErr != nil {
		return nil, dbErr
	}

	// 返回最终数据
	return loadFullUserData(ghUser.ID)
}

// 更新用户
func updateGithubUser(dbUser *model.UserModel, ghUser *model.UserModel) error {
	// 只更新GitHub可能变动的字段
	updates := map[string]interface{}{
		"github_id":  ghUser.GithubID,
		"avatar":     ghUser.Avatar,
		"email":      ghUser.Email,
		"login_type": "github",
	}
	return global.DB.Model(dbUser).Updates(updates).Error
}

// 加载完整的用户关联数据
func loadFullUserData(uid uint) (*model.UserModel, error) {
	var user model.UserModel

	err := global.DB.Preload("Bookmarks").First(&user, uid).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
