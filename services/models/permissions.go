package models

import (
    "gopkg.in/mgo.v2/bson"
)

type Group struct {
    Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Name   string `bson:"username"`
}




