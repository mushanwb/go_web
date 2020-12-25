package auth_controller

import (
	"go_web/app/http/entity"
	"go_web/app/http/models/auth_model"
	"go_web/app/requests"
	"go_web/pkg/util"
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
		passwordHash := util.Hash(userFromData.Password)
		user := auth_model.User{
			Name:     userFromData.Name,
			Email:    userFromData.Email,
			Password: passwordHash,
		}
		_user, _ := user.GetUserByNameOrEmail()
		if _user.ID != 0 {
			w.WriteHeader(http.StatusConflict)
			w.Write(entity.ReturnJson("名称或邮箱已经被创建", nil))
		} else {
			errCreate := user.Create()

			if errCreate == nil {
				w.Write(entity.ReturnJson("用户注册成功", user))
			} else {

				w.WriteHeader(http.StatusInternalServerError)
				w.Write(entity.ReturnJson("注册用户失败", nil))
			}
		}
	}

}

func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	loginFromData := requests.LoginFromData{
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}

	errors := requests.ValidateLoginFrom(loginFromData)
	user := auth_model.User{
		Email: loginFromData.Email,
	}

	if len(errors) == 0 {
		_user, err := user.GetUserByNameOrEmail()

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(entity.ReturnJson("请求参数错误", errors))
	}
}
