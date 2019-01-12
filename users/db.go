package users

import (
	"log"

	"github.com/globalsign/mgo"
)

var (
	// Session stores mongo session
	//session *mgo.Session

	// Mongo stores the mongodb connection string information
	mongo *mgo.DialInfo

	db *mgo.Database

	collection *mgo.Collection
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	//MongoDBUrl = "mongodb://dbapp:VMware1@0.0.0.0:27017/user?authSource=admin"
	MongoDBUrl = "mongodb://mongoadmin:secret@0.0.0.0:27017/?authSource=admin"
)

// ConnectDB accepts name of database and collection as a string
func ConnectDB(dbName string, collectionName string) *mgo.Session {

	Session, error := mgo.Dial(MongoDBUrl)
	if error != nil {
		log.Fatal(error)
	}

	db = Session.DB(dbName)

	collection = db.C(collectionName)

	return Session
}

// CloseDB accepst Session as input to close Connection to the database
func CloseDB(s *mgo.Session) {

	defer s.Close()
	log.Println("Closed connection to db")
}
