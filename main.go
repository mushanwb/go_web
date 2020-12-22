package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go_web/app/http/middlewares"
	"go_web/bootstrap"
	"net/http"
	"strings"
)

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	// 中间件：强制内容类型为 JSON
	router.Use(middlewares.ForceJson)

	// 通过命名路由获取 URL 示例
	//homeURL, _ := router.Get("home").URL()
	//fmt.Println("homeURL: ", homeURL)
	//articleURL, _ := router.Get("articles.show").URL("id", "23")
	//fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
