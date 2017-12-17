package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func main() {
	//session, err := mgo.Dial("localhost")
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Timeout:  60 * time.Second,
		Database: "bookmarkdb",
		Username: "shijuvar",
		Password: "password123",
	}
	session, err := mgo.DialWithInfo(mongoDialInfo)

}
