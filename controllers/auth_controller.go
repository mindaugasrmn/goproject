package controllers

import (
	"github.com/mindaugasrmn/goproject/services"
	"github.com/mindaugasrmn/goproject/services/models"
	"github.com/mindaugasrmn/goproject/core/mongo"
	"github.com/mindaugasrmn/goproject/utils"
	"encoding/json"
	"net/http"

)


func Login(w http.ResponseWriter, r *http.Request) {
	db := mongo.GetDb(r)
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
    user := utils.GetUserByUser(db,requestUser)
    db.FindRef(user.Position).One(&user.Pos)
	responseStatus, token := services.Login(requestUser,user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	db := mongo.GetDb(r)
    user := utils.GetUserById(db,services.GetUserFromToken(r))
    db.FindRef(user.Position).One(&user.Pos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(services.RefreshToken(user))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := services.Logout(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
