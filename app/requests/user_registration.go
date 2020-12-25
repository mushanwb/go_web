package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type RegisterFromData struct {
	Name     string `valid:"name"`
	Email    string `valid:"email"`
	Password string `valid:"password"`
}

func ValidateRegisterFrom(userFromData RegisterFromData) map[string][]string {
	baseRequest := BaseRequest{}
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"name":     baseRequest.NameRules(),
		"email":    baseRequest.EmailRules(),
		"password": baseRequest.PasswordRules(),
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name":     baseRequest.NameMessages(),
		"email":    baseRequest.EmailMessages(),
		"password": baseRequest.PasswordMessages(),
	}

	return baseRequest.Options(userFromData, rules, messages)

}
