package routers

import (
	"github.com/mindaugasrmn/goproject/apicore/controllers/auth"
	"github.com/mindaugasrmn/goproject/apicore/jwt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	subrouter := router.PathPrefix("/api/v1/auth").Subrouter()
	subrouter.HandleFunc("/login/", controllers.Login).Methods("POST")
	subrouter.Handle("/refresh-token/",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")
	subrouter.Handle("/logout/",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")
	return router
}
