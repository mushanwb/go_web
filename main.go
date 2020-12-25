package main

import (
	"go_web/app/http/middlewares"
	"go_web/bootstrap"
	"net/http"
)

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
