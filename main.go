package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ishrivatsa/user/users"
)

//
func handleRequest() {

	// Init Router
	router := gin.Default()

	// Added
	authGroup := router.Group("/")
	{
		authGroup.POST("/register", users.RegisterUser)
		authGroup.POST("/login", users.LoginUser)
		authGroup.GET("/users", users.GetUsers)
		authGroup.GET("/users/:id", users.GetUser)
		authGroup.DELETE("/users/:id", users.DeleteUser)
	}

	router.Run(":8088")
}

func main() {

	dbsession := users.ConnectDB("acmefit", "users")
	log.Printf("Successfully connected to mongodb")

	handleRequest()

	users.CloseDB(dbsession)

}
