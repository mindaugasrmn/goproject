package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/mindaugasrmn/goproject/core/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
)


type User struct {
	Id   bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
}

func getUser(db *mgo.Database) *User{
	var aUser User
	users := db.C("users")
	err := users.Find(bson.M{"username": "test"}).One(&aUser)
	if err != nil {
		log.Print(err)
	}
	return &aUser
}
func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    db := mongo.GetDb(r)
    user :=getUser(db)
    data, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", data)

	
	
}
