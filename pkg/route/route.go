package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

// 获取路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
