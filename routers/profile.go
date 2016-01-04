package routers

import (
	"github.com/mindaugasrmn/goproject/apicore/jwt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mindaugasrmn/goproject/controllers/profile"
)

func SetGetProfileRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/v1/profile/{username}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(mecontroller.GetMyProfileController),
		)).Methods("GET")

	return router
}
