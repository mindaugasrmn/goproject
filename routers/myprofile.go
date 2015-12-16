package routers

import (
	"github.com/mindaugasrmn/goproject/apicore/jwt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mindaugasrmn/goproject/controllers/myprofile"
)

func SetGetMyProfileRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/v1/myprofile/",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(mecontroller.GetMyProfileController),
		)).Methods("GET")

	return router
}
