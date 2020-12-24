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
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"name":     []string{"required", "alpha_num", "between:3,20"},
		"email":    []string{"required", "min:4", "max:30", "email"},
		"password": []string{"required", "min:6"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
		},
	}

	opts := govalidator.Options{
		Data:          &userFromData,
		Rules:         rules,
		TagIdentifier: "valid", // Struct 标签标识符
		Messages:      messages,
	}

	// 4. 开始认证
	errs := govalidator.New(opts).ValidateStruct()

	return errs
}
