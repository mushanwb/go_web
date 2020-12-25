package bootstrap

import (
	"github.com/gorilla/mux"
	"go_web/route"
	"net/http"
	"strings"
)

func SetupRoute() http.Handler {
	router := mux.NewRouter()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Split(r.URL.Path, "/")[1] == "api" {
			route.RegisterApiRoutes(router)
		} else {
			route.RegisterWebRoutes(router)
		}

		router.ServeHTTP(w, r)
	})

}
