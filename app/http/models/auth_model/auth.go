package auth_model

import (
	"go_web/app/http/models"
	"go_web/pkg/logger"
	"go_web/pkg/model"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique" json:"name"`
	Email    string `gorm:"column:email;type:varchar(255);not null;unique" json:"email"`
	Password string `gorm:"column:password;type:varchar(255);" json:"-"`
}

type LoginInfo struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func (user *User) GetUserByNameOrEmail() (User, error) {
	var _user User
	err := model.DB.Where("name = ?", user.Name).Or("email = ?", user.Email).First(&_user).Error
	if err != nil {
		logger.LogError(err)
		return _user, err
	}
	return _user, nil
}
