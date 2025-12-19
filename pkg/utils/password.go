package utils

import "golang.org/x/crypto/bcrypt"

// Encrypt BCrypt密码加密
func Encrypt(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), err
}

// Decrypt 解密
func Decrypt(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
