package route

import (
	"github.com/gorilla/mux"
	"go_web/app/http/middlewares"
	"go_web/app/http/views"
	"net/http"
)

func RegisterWebRoutes(r *mux.Router) {
	pc := new(views.PagesController)

	// home 首页
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")

	// 自定义 404 页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)

	// 中间件：强制内容类型为 HTML
	r.Use(middlewares.ForceHtml)

}
