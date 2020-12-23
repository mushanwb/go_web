package route

import (
	"github.com/gorilla/mux"
	"go_web/app/http/controllers/article_controller"
	"go_web/app/http/controllers/auth_controller"
)

func RegisterApiRoutes(r *mux.Router) {
	ac := new(article_controller.ArticleController)

	// 通过命名路由获取 URL 示例
	//homeURL, _ := router.Get("home").URL()
	//fmt.Println("homeURL: ", homeURL)
	//articleURL, _ := router.Get("articles.show").URL("id", "23")
	//fmt.Println("articleURL: ", articleURL)

	r.HandleFunc("/articles/{id:[0-9]+}", ac.ArticlesShowHandler).Methods("GET").Name("home")
	r.HandleFunc("/articles", ac.ArticlesIndexHandler).Methods("GET").Name("home")
	r.HandleFunc("/articles", ac.ArticlesStoreHandler).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.ArticlesUpdateHandler).Methods("PUT").Name("articles.update")
	// 同名的路由,根据请求的方式不同，选择进入不同的函数
	r.HandleFunc("/articles/{id:[0-9]+}", ac.ArticlesDeleteHandler).Methods("DELETE").Name("articles.delete")

	auth := new(auth_controller.AuthController)
	r.HandleFunc("/register", auth.DoRegister).Methods("POST")

}
