package auth_model

import "go_web/app/http/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique" json:"name"`
	Email    string `gorm:"column:email;type:varchar(255);not null;unique" json:"email"`
	Password string `gorm:"column:password;type:varchar(255);" json:"-"`
}
