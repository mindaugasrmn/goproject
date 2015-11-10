package routers

import (
	"github.com/mindaugasrmn/goproject/controllers"
	"github.com/mindaugasrmn/goproject/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetGetMeRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/v1/me/",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.GetMeController),
		)).Methods("GET")

	return router
}
