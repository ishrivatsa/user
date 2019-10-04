package service

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/vmwarecloudadvocacy/user/internal/auth"
	"github.com/vmwarecloudadvocacy/user/internal/db"
	"github.com/vmwarecloudadvocacy/user/pkg/logger"

)


// GetUsers accepts a context and returns all the users in json format
func GetUsers(c *gin.Context) {
	var users []auth.UserResponse

	logger.Logger.Infof("Retrieving All Users")

	error := db.Collection.Find(nil).All(&users)

	if error != nil {
		message := "Users " + error.Error()
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

// GetUser accepts context, User ID as param and returns user info
func GetUser(c *gin.Context) {
	var user auth.UserResponse

	userID := c.Param("id")

	if bson.IsObjectIdHex(userID) {

		error := db.Collection.FindId(bson.ObjectIdHex(userID)).One(&user)

		if error != nil {
			message := "User " + error.Error()
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
			return
		}
	} else {
		message := "Incorrect Format for UserID"
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

// RegisterUser accepts context and inserts data to the db
func RegisterUser(c *gin.Context) {

	var user auth.User

	error := c.ShouldBindJSON(&user)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Incorrect Field Name(s)/ Value(s)"})
		return
	}

	error = user.Validate()

	if error != nil {
		message := "User " + error.Error()
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
		return
	}

	// Inserts ID for the user object
	user.ID = bson.NewObjectId()

	user.Password = auth.CalculatePassHash(user.Password, user.Salt)

	error = db.Collection.Insert(&user)

	if error != nil {
		message := "User " + error.Error()
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!", "resourceId": user.ID})

}

// LoginUser 
func LoginUser(c *gin.Context) {
	var user auth.User

	error := c.ShouldBindJSON(&user)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Incorrect Field Name(s)"})
		return
	}

	userpass := user.Password

	error = db.Collection.Find(bson.M{"username": user.Username}).One(&user)

	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Username"})
		return
	}

	if user.Password != auth.CalculatePassHash(userpass, user.Salt) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Password"})
		return
	}

	// Declare the expiration time of the token
	// Here the expiration is 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &auth.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			// The expiry time is in milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(auth.JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "token": tokenString})

}

// DeleteUser accepts user id and deletes the specific user
func DeleteUser(c *gin.Context) {

	userID := c.Param("id")

	if bson.IsObjectIdHex(userID) != true {
		message := "Invalid User ID "
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
		return
	}

	error := db.Collection.RemoveId(bson.ObjectIdHex(userID))

	if error != nil {
		message := "User " + error.Error()
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}
