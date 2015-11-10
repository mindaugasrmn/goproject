package mongo

import (
    "github.com/mindaugasrmn/goproject/settings"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"github.com/codegangsta/negroni"
	"net/http"
)

type key int

const db key = 0

func GetDb(r *http.Request) *mgo.Database {
	if rv := context.Get(r, db); rv != nil {
		return rv.(*mgo.Database)
	}
	return nil
}

func SetDb(r *http.Request, val *mgo.Database) {
	context.Set(r, db, val)
}

func MongoMiddleware() negroni.HandlerFunc {
	var conf = settings.ReadConfig()
	database := conf.Mongo_db
	session, err := mgo.Dial(conf.Mongo_svr)

	if err != nil {
		panic(err)
	}

	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		reqSession := session.Clone()
		defer reqSession.Close()
		db := reqSession.DB(database)
		SetDb(r, db)
		next(rw, r)
	})
}
