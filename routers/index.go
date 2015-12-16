package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func StaticFileServer(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "static/")
}

func SetIndexRoutes(router *mux.Router) *mux.Router {
	router.PathPrefix("/static/").Handler(http.FileServer(http.Dir("static/"))).Methods("GET")
	router.PathPrefix("/{_:.*}").HandlerFunc(StaticFileServer).Methods("GET")
	return router
}
