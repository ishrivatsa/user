package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

//
func handleRequest() {

	// Init Router
	router := gin.Default()

	// Added
	authGroup := router.Group("/")
	{
		authGroup.POST("/register", RegisterUser)
		authGroup.POST("/login", LoginUser)
		authGroup.GET("/users", GetUsers)
		authGroup.GET("/users/:id", GetUser)
		authGroup.DELETE("/users/:id", DeleteUser)
	}

	router.Run(":8088")
}

func main() {

	dbsession := ConnectDB("acmefit", "users")
	log.Printf("Successfully connected to mongodb")

	handleRequest()

	CloseDB(dbsession)

}
