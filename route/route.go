package route

import "github.com/gorilla/mux"

// 使用 精准匹配 的 gorilla/mux
var Router *mux.Router

// Initialize 初始化路由
func Initialize() {
	Router = mux.NewRouter()
}
