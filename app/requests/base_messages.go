package requests

func (*BaseRequest) EmailMessages() []string {
	return []string{
		"required:Email 为必填项",
		"min:Email 长度需大于 4",
		"max:Email 长度需小于 30",
		"email:Email 格式不正确，请提供有效的邮箱地址",
	}
}

func (*BaseRequest) NameMessages() []string {
	return []string{
		"required:用户名为必填项",
		"alpha_num:格式错误，只允许数字和英文",
		"between:用户名长度需在 3~20 之间",
	}
}

func (*BaseRequest) PasswordMessages() []string {
	return []string{
		"required:密码为必填项",
		"min:长度需大于 6",
	}
}
