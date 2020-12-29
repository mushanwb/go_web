package middlewares

import (
	"context"
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
		// 将用户信息存到 context 中
		ctx := context.WithValue(r.Context(), "user", user)
		// 在别的文件中，从 context 获取到用户信息
		//user := r.Context().Value("user").(auth_model.User)

		// 继续处理请求
		next(w, r.WithContext(ctx))
	}
}
