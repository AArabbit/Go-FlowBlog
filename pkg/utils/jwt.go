package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// 生成token的密钥Secret
var (
	accessSecret  = []byte("access_secret_ARabbit")
	refreshSecret = []byte("refresh_secret_ARabbit")
)

// token 过期时间
const (
	AccessTokenDuration  = 15 * time.Minute   // Access 15分钟
	RefreshTokenDuration = 7 * 24 * time.Hour // Refresh 7天
)

// JwtClaims jwt结构体
type JwtClaims struct {
	UserID    uint
	Username  string
	TokenType string // 区分是 access 还是 refresh，防止攻击
	jwt.RegisteredClaims
}

// 工厂函数
func generateToken(userID uint, username string,
	DurationType time.Duration, tokenType string, secret []byte) (string, error) {

	tokenClaims := JwtClaims{
		UserID:    userID,
		Username:  username,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置过期时间与签发时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(DurationType)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    viper.GetString("app.name"),
		},
	}
	resToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	// 使用密钥生成token
	resStr, err := resToken.SignedString(secret)
	if err != nil {
		return "", err
	}
	return resStr, nil
}

// GenerateAccessToken 生成AccessToken
func GenerateAccessToken(userID uint, username string) (string, error) {
	accessToken, err := generateToken(userID, username,
		AccessTokenDuration, "access", accessSecret)

	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// GenerateRefreshToken  生成RefreshToken
func GenerateRefreshToken(userID uint, username string) (string, error) {
	refreshToken, err := generateToken(userID, username,
		RefreshTokenDuration, "refresh", refreshSecret)

	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

// ParseToken 解析 Token
// isRefresh: 是否解析 Refresh Token (使用不同的 Secret)
func ParseToken(tokenString string, isRefresh bool) (*JwtClaims, error) {
	secret := accessSecret
	if isRefresh {
		secret = refreshSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}
