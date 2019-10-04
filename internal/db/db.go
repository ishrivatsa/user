package db

import (
	"fmt"
	"os"

	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"github.com/vmwarecloudadvocacy/user/pkg/logger"
//	"github.com/vmwarecloudadvocacy/user/pkg/logger"
) 

var (
	// Session stores mongo session
	//session *mgo.Session

	// Mongo stores the mongodb connection string information
	mongo *mgo.DialInfo

	DB *mgo.Database

	Collection *mgo.Collection
)


// GetEnv accepts the ENV as key and a default string
// If the lookup returns false then it uses the default string else it leverages the value set in ENV variable
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	
	logger.Logger.Info("Setting default values for ENV variable " + key)
	return fallback
}

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

	DB = Session.DB(dbName)

	error = DB.Session.Ping()
	if error != nil {
		logger.Errorf("Unable to connect to database %s", dbName)
	}

	Collection = DB.C(collectionName)

	logger.Info("Connected to database and the collection")

	return Session
}

// CloseDB accepst Session as input to close Connection to the database
func CloseDB(s *mgo.Session, logger *logrus.Logger) {

	defer s.Close()
	logger.Info("Closed connection to db")
}
