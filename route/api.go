package route

import (
	"github.com/gorilla/mux"
	"go_web/app/http/controllers/article_controller"
)

func RegisterApiRoutes(r *mux.Router) {
	ac := new(article_controller.ArticleController)

	r.HandleFunc("/articles/{id:[0-9]+}", ac.ArticlesShowHandler).Methods("GET").Name("home")
	r.HandleFunc("/articles", ac.ArticlesIndexHandler).Methods("GET").Name("home")
	r.HandleFunc("/articles", ac.ArticlesStoreHandler).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.ArticlesUpdateHandler).Methods("PUT").Name("articles.update")
	// 同名的路由,根据请求的方式不同，选择进入不同的函数
	r.HandleFunc("/articles/{id:[0-9]+}", ac.ArticlesDeleteHandler).Methods("DELETE").Name("articles.delete")

}
