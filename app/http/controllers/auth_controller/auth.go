package auth_controller

import (
	"go_web/app/http/entity"
	"go_web/app/http/models/auth_model"
	"go_web/app/requests"
	"net/http"
)

type AuthController struct {
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {

	userFromData := requests.RegisterFromData{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}

	errors := requests.ValidateRegisterFrom(userFromData)

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(entity.ReturnJson("请求参数错误", errors))
	} else {
		user := auth_model.User{
			Name:     userFromData.Name,
			Email:    userFromData.Email,
			Password: userFromData.Password,
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