package requests

import "github.com/thedevsaddam/govalidator"

type BaseRequest struct {
}

func (*BaseRequest) Options(data interface{}, rules map[string][]string, messages map[string][]string) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // Struct 标签标识符
		Messages:      messages,
	}

	// 4. 开始认证
	errs := govalidator.New(opts).ValidateStruct()

	return errs
}
