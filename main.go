package main

import (
	"go_web/app/http/middlewares"
	"go_web/bootstrap"
	"net/http"
)

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	// 中间件：强制内容类型为 JSON
	router.Use(middlewares.ForceJson)

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
