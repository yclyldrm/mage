package models

import (
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    int64              `json:"userId" bson:"userId"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

var (
	USERNAME_NULL_Error = "Username can't be null"
	PASSWORD_NULL_Error = "Password can't be null"
)

func (u *User) IsValidInfo() (bool, error) {
	if u.Username == "" {
		return false, errors.New(USERNAME_NULL_Error)
	}

	if u.Password == "" {
		return false, errors.New(PASSWORD_NULL_Error)
	}

	return true, nil
}

func (u *User) HashedPassword() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("While password crypting, an error occured", err.Error())
		return
	}

	u.Password = string(hashedPassword)
}
