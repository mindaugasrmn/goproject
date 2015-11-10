package utils

import (
	"github.com/mindaugasrmn/goproject/services/models"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"

)

func GetUserByUser(db *mgo.Database,requestUser *models.User) *models.User{
   var aUser models.User
   user := db.C("users")
	err := user.Find(bson.M{"username": requestUser.Username}).One(&aUser)
	if err != nil {
		log.Print(err)
	}
	return &aUser
}

func GetUserById(db *mgo.Database,userId string) *models.User{
   var aUser models.User
   user := db.C("users")
	err := user.Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&aUser)
	if err != nil {
		log.Print(err)
	}
	return &aUser
}

func GetUserByIdMe(db *mgo.Database,userId string) *models.Me{
   var aUser models.Me
   user := db.C("users")
	err := user.Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&aUser)
	if err != nil {
		log.Print(err)
	}
	return &aUser
}

