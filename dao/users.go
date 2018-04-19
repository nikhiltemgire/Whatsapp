package dao

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User : User db Structure
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

var userCollection = GetUserCollection(session)

// GetAllUsers returns all users from users collection
func GetAllUsers() []User {
	var users []User

	err := userCollection.Find(nil).All(&users)
	if err != nil {
		panic(err)
	}

	return users
}

// GetUserByEmail return the user if exists
func GetUserByEmail(email string) User {
	var user User

	err := userCollection.Find(bson.M{"email": email}).One(&user)

	if err == mgo.ErrNotFound {
		fmt.Println("Not Found.")
		return user
	} else if err != nil {
		panic(err)
	}

	return user
}

// CheckPassword compares the entered password with pass of user
func AuthenticateUser(email string, password string) bool {
	user := GetUserByEmail(email)

	if password == user.Password {
		return true
	}
	return false
}

func CreateUser(email string, password string) {
	var user User
	user.Email = email
	user.Password = password
	err := userCollection.Insert(&user)
	if err != nil {
		panic(err)
	}
}
