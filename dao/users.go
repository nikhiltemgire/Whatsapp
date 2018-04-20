package dao

import (
	"Whatsapp/utils"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User : User db Structure
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
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
		return user
	} else if err != nil {
		panic(err)
	}

	return user
}

// AuthenticateUser : CheckPassword compares the entered password with pass of user
func AuthenticateUser(email string, password string) bool {
	user := GetUserByEmail(email)

	pass := utils.GetMD5Hash(password)
	fmt.Println("In Authenticate:", password, pass)
	fmt.Println(pass, user.Password)

	if pass == user.Password {
		return true
	}
	return false
}

// CreateUser : Create user in Users Collection
func CreateUser(id bson.ObjectId, email string, password string) bool {
	var user User
	user.ID = id
	user.Email = email
	user.Password = utils.GetMD5Hash(password)

	fmt.Println("In Create User:", password, user.Password)
	err := userCollection.Insert(&user)
	if err != nil {
		return false
	}
	return true
}
