package route

import (
	"github.com/gorilla/mux"
	"go_web/app/http/controllers/article_controller"
)

func RegisterApiRoutes(r *mux.Router) {
	ac := new(article_controller.ArticleController)

	r.HandleFunc("/articles/{id:[0-9]+}", ac.ArticlesShowHandler).Methods("GET").Name("home")
	r.HandleFunc("/articles", ac.ArticlesIndexHandler).Methods("GET").Name("home")

}
