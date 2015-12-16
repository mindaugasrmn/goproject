package services

import (
	"github.com/mindaugasrmn/goproject/apicore/parameters"
	"github.com/mindaugasrmn/goproject/apicore/jwt"
	"github.com/mindaugasrmn/goproject/apicore/models/auth"
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"log"
)

func Login(requestUser *models.User,user *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	if authBackend.Authenticate(requestUser,user) {
		token, err := authBackend.GenerateToken(user)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(parameters.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}

func GetUserFromToken(req *http.Request) string{
	var err error
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{},error) {
		return authBackend.PublicKey,nil
	})
	if err != nil {
		return "err"
	}
    return tokenRequest.Claims["sub"].(string)
}
func RefreshToken(requestUser *models.User) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	log.Print(tokenRequest.Claims["sub"])
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}
