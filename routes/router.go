package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range ApplicationRoutes {
		var handler http.Handler

		handler = route.Handler
		//handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
