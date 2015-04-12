package database

import (
	"fmt"
	"log"
	_ "time"

	mgo "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session
)

// StartMongoDb start mongodb connection
func StartMongoDb() {
	currentSession, err := mgo.Dial("104.155.227.195:27020")
	if err != nil {
		log.Println("err connecting to mongodb!")
		log.Println("error: ", err)
		return
	}
	fmt.Println("connected to mongos!")
	session = currentSession
}
