package utils

import (
	"github.com/mindaugasrmn/goproject/apicore/models/auth"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"

)

func GetUserByUsername(db *mgo.Database,requestUser *models.User) *models.User{
    var aUser models.User
    user := db.C("users")
    log.Print(requestUser.Username)
	err := user.Find(bson.M{"username": requestUser.Username}).One(&aUser)
	if err != nil {
		log.Print(err)
	} else {
	    db.FindRef(aUser.Position).One(&aUser.Pos)
    }
    return &aUser
}


func GetProfileByUsername(db *mgo.Database,username string) *models.Profile{
    var aUser models.Profile
    user := db.C("users")
	err := user.Find(bson.M{"username": username}).One(&aUser)
	if err != nil {
		return &aUser
	} else {
	    db.FindRef(aUser.Position).One(&aUser.Pos)
    }
    return &aUser
}
//func UserExist(db *mgo.Database,requestUser *models.User) bool {
//    var aUser models.User
//    err := user.Find(bson.M{"username": requestUser.Username}).One(&aUser)
//    if err != nil {
//        return false
//    } else {
//        return true
//    }
//}

func GetUserById(db *mgo.Database,userId string) *models.User{
   var aUser models.User
   user := db.C("users")
	err := user.Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&aUser)
	if err != nil {
		log.Print(err)
	}
	return &aUser
}

func GetUserByIdMe(db *mgo.Database,userId string) *models.Profile{
   var aUser models.Profile
   user := db.C("users")
	err := user.Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&aUser)
	if err != nil {
		log.Print(err)
	}
	return &aUser
}

