package router

import (
	"github.com/gorilla/mux"
	"github.com/jittakal/go-apps/ams/city"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range city.Routes {
		router.PathPrefix("/ams/city").
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	return router
}
