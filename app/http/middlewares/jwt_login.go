package middlewares

import (
	"context"
	"fmt"
	"go_web/app/http/entity"
	"net/http"
)

func JwtAuth(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(entity.ReturnJson("用户未登录", nil))
			return
		}

		user, bl := entity.ParseToken(token)
		if !bl {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(entity.ReturnJson("token 验证失败", nil))
			return
		}
		fmt.Println(user)
		ctx := context.WithValue(r.Context(), "user", user)
		// 继续处理请求
		next(w, r.WithContext(ctx))
	}
}
