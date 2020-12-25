package util

import (
	"go_web/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogError(err)
	return string(bytes)
}

// CheckHash 对比明文密码和数据库的哈希值
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	logger.LogError(err)
	return err == nil
}
