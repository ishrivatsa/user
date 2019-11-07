package db

import (
	"fmt"
	"os"

	"github.com/globalsign/mgo"
	redis "github.com/go-redis/redis/v7"
	"github.com/vmwarecloudadvocacy/user/pkg/logger"
)

var (
	// Session stores mongo session
	//session *mgo.Session

	// Mongo stores the mongodb connection string information
	mongo *mgo.DialInfo

	DB *mgo.Database

	Collection *mgo.Collection

	RedisClient *redis.Client
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

// ConnectRedisDB returns a redis client
func ConnectRedisDB() *redis.Client {

	redisHost := GetEnv("REDIS_DB_HOST", "0.0.0.0")
	redisPort := GetEnv("REDIS_DB_PORT", "6379")

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping().Result()
	logger.Logger.Infof("Reply from Redis %s", pong)
	if err != nil {
		fmt.Errorf(err.Error())
		logger.Logger.Fatalf("Failed connecting to redis db %s", err.Error())
		os.Exit(1)
	}
	logger.Logger.Infof("Successfully connected to redis database")
	return RedisClient
}

// ConnectDB accepts name of database and collection as a string
func ConnectDB(dbName string, collectionName string) *mgo.Session {

	dbUsername := os.Getenv("USERS_DB_USERNAME")
	dbSecret := os.Getenv("USERS_DB_PASSWORD")

	// Get ENV variable or set to default value
	dbIP := GetEnv("USERS_DB_HOST", "0.0.0.0")
	dbPort := GetEnv("USERS_DB_PORT", "27017")

	mongoDBUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", dbUsername, dbSecret, dbIP, dbPort)

	Session, error := mgo.Dial(mongoDBUrl)

	if error != nil {
		fmt.Printf(error.Error())
		logger.Logger.Fatalf(error.Error())
		os.Exit(1)
	}

	DB = Session.DB(dbName)

	error = DB.Session.Ping()
	if error != nil {
		logger.Logger.Errorf("Unable to connect to database %s", dbName)
	}

	Collection = DB.C(collectionName)

	logger.Logger.Info("Connected to database and the collection")

	return Session
}

// CloseDB accepst Session as input to close Connection to the database
func CloseDB(s *mgo.Session) {

	defer s.Close()
	logger.Logger.Info("Closed connection to db")
}
