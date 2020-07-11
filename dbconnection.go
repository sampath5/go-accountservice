package main

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	host := "MONGO_HOST"
	dbName := "go_ms_test"
	fmt.Println("conn info:", host, dbName)
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("session err:", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
}

// GetMongoDB function to return DB connection
func GetMongoDB() *mgo.Database {
	return db
}
