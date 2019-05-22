package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	stdopentracing "github.com/opentracing/opentracing-go"
	tracelog "github.com/opentracing/opentracing-go/log"
)

func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// GetUsers accepts a context and returns all the users in json format
func GetUsers(c *gin.Context) {
	var users []UserResponse

	tracer := stdopentracing.GlobalTracer()

	userSpanCtx, _ := tracer.Extract(stdopentracing.HTTPHeaders, stdopentracing.HTTPHeadersCarrier(c.Request.Header))

	userSpan := tracer.StartSpan("db_get_users", stdopentracing.FollowsFrom(userSpanCtx))
	defer userSpan.Finish()

	error := collection.Find(nil).All(&users)

	if error != nil {
		message := "Users " + error.Error()
		userSpan.LogFields(
			tracelog.String("event", "error"),
			tracelog.String("message", error.Error()),
		)
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

// GetUser accepts context, User ID as param and returns user info
func GetUser(c *gin.Context) {
	var user UserResponse

	tracer := stdopentracing.GlobalTracer()

	userSpanCtx, _ := tracer.Extract(stdopentracing.HTTPHeaders, stdopentracing.HTTPHeadersCarrier(c.Request.Header))

	userSpan := tracer.StartSpan("db_get_user", stdopentracing.FollowsFrom(userSpanCtx))

	defer userSpan.Finish()

	userID := c.Param("id")

	productSpan.LogFields(
		tracelog.String("event", "string-format"),
		tracelog.String("user.id", userID),
	)

	if bson.IsObjectIdHex(userID) {

		error := collection.FindId(bson.ObjectIdHex(userID)).One(&user)

		if error != nil {
			message := "User " + error.Error()
			userSpan.LogFields(
				tracelog.String("event", "error"),
				tracelog.String("message", error.Error()),
			)
			userSpan.SetTag("http.status_code", http.StatusNotFound)
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
			return
		}
	} else {
		message := "Incorrect Format for UserID"
		userSpan.LogFields(
			tracelog.String("event", "error"),
			tracelog.String("message", message),
		)
		userSpan.SetTag("http.status_code", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
		return
	}

	userSpan.SetTag("http.status_code", http.StatusOK)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

// CreateUser accepts context and inserts data to the db
func RegisterUser(c *gin.Context) {

	var user User

	tracer := stdopentracing.GlobalTracer()

	userSpanCtx, _ := tracer.Extract(stdopentracing.HTTPHeaders, stdopentracing.HTTPHeadersCarrier(c.Request.Header))

	userSpan := tracer.StartSpan("db_login_user", stdopentracing.FollowsFrom(userSpanCtx))

	defer userSpan.Finish()

	error := c.ShouldBindJSON(&user)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Incorrect Field Name(s)/ Value(s)"})
		return
	}

	error = user.Validate()

	if error != nil {
		message := "User " + error.Error()
		userSpan.LogFields(
			tracelog.String("event", "error"),
			tracelog.String("message", error.Error()),
		)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
		return
	}

	// Inserts ID for the user object
	user.ID = bson.NewObjectId()

	user.Password = calculatePassHash(user.Password, user.Salt)

	error = collection.Insert(&user)

	if error != nil {
		message := "User " + error.Error()
		userSpan.LogFields(
			tracelog.String("event", "error"),
			tracelog.String("message", error.Error()),
		)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
		return
	}

	userSpan.SetTag("http.status_code", http.StatusCreated)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!", "resourceId": user.ID})

}

func LoginUser(c *gin.Context) {
	var user User

	tracer := stdopentracing.GlobalTracer()

	userSpanCtx, _ := tracer.Extract(stdopentracing.HTTPHeaders, stdopentracing.HTTPHeadersCarrier(c.Request.Header))

	userSpan := tracer.StartSpan("db_login_user", stdopentracing.FollowsFrom(userSpanCtx))

	defer userSpan.Finish()

	error := c.ShouldBindJSON(&user)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Incorrect Field Name(s)"})
		return
	}

	userpass := user.Password

	error = collection.Find(bson.M{"username": user.Username}).One(&user)

	if error != nil {
		userSpan.SetTag("http.status_code", http.StatusNotFound)
		userSpan.LogFields(
			tracelog.String("event", "error"),
			tracelog.String("message", error.Error()),
		)
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Username"})
		return
	}

	if user.Password != calculatePassHash(userpass, user.Salt) {
		userSpan.SetTag("http.status_code", http.StatusNotFound)
		userSpan.LogFields(
			tracelog.String("event", "error"),
			tracelog.String("message", error.Error()),
		)
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Password"})
		return
	}

	userSpan.SetTag("http.status_code", http.StatusOK)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "token": user.ID})

}

// DeleteUser accepts user id and deletes the specific user
func DeleteUser(c *gin.Context) {

	userID := c.Param("id")

	if bson.IsObjectIdHex(userID) != true {
		message := "Invalid User ID "
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
		return
	}

	error := collection.RemoveId(bson.ObjectIdHex(userID))

	log.Println(error)

	if error != nil {
		message := "User " + error.Error()
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}
