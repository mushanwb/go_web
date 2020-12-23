package auth_controller

import (
	"go_web/app/http/entity"
	"go_web/app/http/models/auth_model"
	"net/http"
)

type AuthController struct {
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	errors := validateUserFormData(name, email, password)

	if len(errors) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(entity.ReturnJson("请求参数错误", errors))
	} else {
		user := auth_model.User{
			Name:     name,
			Email:    email,
			Password: password,
		}

		err := user.Create()

		if err == nil {
			w.Write(entity.ReturnJson("用户注册成功", user))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(entity.ReturnJson("注册用户失败", nil))
		}
	}

}

func validateUserFormData(name string, email string, password string) map[string]string {
	errors := make(map[string]string)

	if name == "" {
		errors["name"] = "名字不能为空"
	}

	if email == "" {
		errors["email"] = "邮箱不能为空"
	}

	if password == "" {
		errors["password"] = "密码不能为空"
	}
	return errors
}
