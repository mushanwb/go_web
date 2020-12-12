package main

import (
	"fmt"
	"net/http"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	url := r.URL.Path
	switch url {
	case "/":
		fmt.Fprint(w, "<h1>Hello, 欢迎来到 Golang！</h1>")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

// 使用 go 自带的 NewServeMux, 需要判断 GET 或者 POST 方法
func articlesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "访问文章列表")
	case "POST":
		fmt.Fprint(w, "创建文章")
	}
}

// 使用 go 自带的 NewServeMux, 需要用切割的方法判断参数
func articleHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	fmt.Fprint(w, "文章 id："+id)

}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)
	router.HandleFunc("/articles", articlesHandler)
	router.HandleFunc("/article/", articleHandler)
	http.ListenAndServe(":3000", router)
}
