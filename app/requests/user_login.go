package requests

import "github.com/thedevsaddam/govalidator"

type LoginFromData struct {
	Email    string `valid:"email"`
	Password string `valid:"password"`
}

func ValidateLoginFrom(loginFromData LoginFromData) map[string][]string {
	baseRequest := BaseRequest{}
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"email":    baseRequest.EmailRules(),
		"password": baseRequest.PasswordRules(),
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"email":    baseRequest.EmailMessages(),
		"password": baseRequest.PasswordMessages(),
	}

	return baseRequest.Options(&loginFromData, rules, messages)
}
