package requests

func (*BaseRequest) EmailRules() []string {
	return []string{"required", "min:4", "max:30", "email"}
}

func (*BaseRequest) NameRules() []string {
	return []string{"required", "alpha_num", "between:3,20"}
}

func (*BaseRequest) PasswordRules() []string {
	return []string{"required", "min:6"}
}
