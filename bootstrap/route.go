package bootstrap

import (
	"github.com/gorilla/mux"
	"go_web/route"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	route.RegisterWebRoutes(router)
	route.RegisterApiRoutes(router)
	return router
}
