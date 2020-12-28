package entity

import (
	"github.com/dgrijalva/jwt-go"
	"go_web/app/http/models"
	"go_web/app/http/models/auth_model"
	"go_web/pkg/logger"
	"time"
)

var jwtKey = []byte("mushan")

var expireTime = time.Now().Add(7 * 24 * time.Hour)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func JwtGetToken(user auth_model.User) string {
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "mu_shan", // 签名颁发者
			Subject:   "go_web",  //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		logger.LogError(err)
	}
	return tokenString
}

func ParseToken(tokenString string) (auth_model.User, bool) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return auth_model.User{}, false
	}

	user := auth_model.User{
		BaseModel: models.BaseModel{
			ID: claims.UserId,
		},
	}
	_user, err := user.GetUserById()
	if err != nil {
		return _user, false
	}
	return _user, true
}
