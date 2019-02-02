package main

import (
	"fmt"
	"os"

	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
)

var (
	// Session stores mongo session
	//session *mgo.Session

	// Mongo stores the mongodb connection string information
	mongo *mgo.DialInfo

	db *mgo.Database

	collection *mgo.Collection
)

// ConnectDB accepts name of database and collection as a string
func ConnectDB(dbName string, collectionName string, logger *logrus.Logger) *mgo.Session {

	dbUsername := os.Getenv("USERS_DB_USERNAME")
	dbSecret := os.Getenv("USERS_DB_PASSWORD")

	// Get ENV variable or set to default value
	dbIP := GetEnv("USERS_DB_HOST", "0.0.0.0")
	dbPort := GetEnv("USERS_DB_PORT", "27017")

	mongoDBUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", dbUsername, dbSecret, dbIP, dbPort)

	Session, error := mgo.Dial(mongoDBUrl)

	if error != nil {
		fmt.Printf(error.Error())
		logger.Fatalf(error.Error())
		os.Exit(1)
	}

	db = Session.DB(dbName)

	error = db.Session.Ping()
	if error != nil {
		logger.Errorf("Unable to connect to database %s", dbName)
	}

	collection = db.C(collectionName)

	logger.Info("Connected to database and the collection")

	return Session
}

// CloseDB accepst Session as input to close Connection to the database
func CloseDB(s *mgo.Session, logger *logrus.Logger) {

	defer s.Close()
	logger.Info("Closed connection to db")
}
