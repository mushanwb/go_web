package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

// 使用 精准匹配 的 gorilla/mux
var Router *mux.Router

// Initialize 初始化路由
func Initialize() {
	Router = mux.NewRouter()
}

// 获取路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
