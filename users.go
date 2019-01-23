package main

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"
)

var (
	ErrNoCustomerInResponse = errors.New("Response has no matching customer")
	ErrMissingField         = "Error missing %v"
)

type registerRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type UserResponse struct {
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
	ID        bson.ObjectId `json:"id" bson:"_id"`
}

// User Struct
type User struct {
	FirstName string        `json:"firstname" bson:"firstname"`
	LastName  string        `json:"lastname" bson:"lastname"`
	Email     string        `json:"email" bson:"email"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"password" bson:"password"`
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Salt      string        `json:"-" bson:"salt"`
}

// Validate if the fields are available
func (u *User) Validate() error {
	if u.FirstName == "" {
		return fmt.Errorf(ErrMissingField, "FirstName")
	}
	if u.LastName == "" {
		return fmt.Errorf(ErrMissingField, "LastName")
	}
	if u.Username == "" {
		return fmt.Errorf(ErrMissingField, "Username")
	}
	if u.Password == "" {
		return fmt.Errorf(ErrMissingField, "Password")
	}
	return nil
}

func (u *User) NewSalt() {
	h := sha1.New()
	io.WriteString(h, strconv.Itoa(int(time.Now().UnixNano())))
	u.Salt = fmt.Sprintf("%x", h.Sum(nil))
}
