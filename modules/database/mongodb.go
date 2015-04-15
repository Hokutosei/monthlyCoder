package database

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"

	// use for query
	_ "gopkg.in/mgo.v2/bson"

	"monthlyCoder/modules/config"
)

var (
	session *mgo.Session

	mongodbClusterKey = "mongodb_cluster1"
)

// GetMongodbCluster retrieve mongodb cluster host
func GetMongodbCluster(host chan string) {
	mongodbCluster, err := config.EtcdRawGetValue(mongodbClusterKey)
	if err != nil {
		panic(err)
	}

	host <- mongodbCluster
}

// StartMongoDb start mongodb connection
func StartMongoDb() {
	hostChan := make(chan string)
	go GetMongodbCluster(hostChan)

	host := <-hostChan
	fmt.Println(host)
	fmt.Println("connecting to mongodb..")
	currentSession, err := mgo.Dial(host)
	if err != nil {
		log.Println("err connecting to mongodb!")
		log.Println("error: ", err)
		return
	}
	fmt.Println("connected to mongos!")
	session = currentSession
}
