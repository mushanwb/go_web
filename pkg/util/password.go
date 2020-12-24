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
