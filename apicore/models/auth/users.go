package models

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type User struct {
    Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Username   string `bson:"username"`
    Password   string `bson:"password"`
    FirstName  string `bson:"first_name"`
    MidName    string `bson:"mid_name"`
    LastName   time.Time `bson:"last_name"`
    Gender     time.Time `bson:"gender"`
    Position   *mgo.DBRef `bson:"position"`
    Pos        Pos `bson:pos`

}

type Me struct {
    Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Username   string `bson:"username"`
    FirstName  string `bson:"first_name"`
    MidName    string `bson:"mid_name"`
    LastName   string `bson:"last_name"`
    Gender     string `bson:"gender"`
    Created_at time.Time `bson:"created_at"`
    Updated_at time.Time `bson:"updated_at"`
    Phone      []Phone `bson:"phone"`
    Position   *mgo.DBRef `bson:"position"`
    Pos        Pos `bson:pos`
}

type Phone struct {
    Name       string `bson:"name"`
    Phone      string `bson:"phone"`
}


type Position struct {
    Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Color         string `bson:"color"`
    PositionName  string `bson:"positionname"`
}

type Pos struct {
    Color         string `bson:"color"`
    PositionName  string `bson:"positionname"`
}


