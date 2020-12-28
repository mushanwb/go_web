package middlewares

import (
	"context"
	"go_web/app/http/entity"
	"net/http"
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		ctx := context.WithValue(r.Context(), "user", user)
		// 继续处理请求
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
