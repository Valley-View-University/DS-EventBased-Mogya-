package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"./handlers"
	"./models"
	"./logging"
)

func NewRouter() *mux.Router {
	router:= mux.NewRouter().StrictSlash(true)
	for _, route := range routes{
		var handler http.Handler

		handler = route.HandlerFunction
		handler = logging.Logger(handler, route.Name)

		//todo: add the logger handler function

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

var routes = models.Routes{
	//the index route
	models.Route{"Index", http.MethodGet,"/",handlers.Index},
}
