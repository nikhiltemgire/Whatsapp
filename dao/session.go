package dao

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

// UserSession : Structure of usersession in Mongo
type UserSession struct {
	UserID    bson.ObjectId `json:"userid" bson:"_id"`
	Token     string        `json:"token" bson:"token"`
	LastLogin time.Time     `json:"lastlogin" bson:"lastlogin"`
}

var sessionCollection = getSessionCollection(session)

// GetAllSession : Returns all the User Sessions
func GetAllSession() []UserSession {
	var usersessions []UserSession

	err := sessionCollection.Find(nil).All(&usersessions)

	if err == mgo.ErrNotFound {
		print("Session Collection: Not Found")
		return usersessions
	} else if err != nil {
		panic(err)
	}

	return usersessions
}

//CreateSession : Creates a new session for a user
func CreateSession(userObjectID bson.ObjectId) string {

	us := UserSession{
		UserID:    userObjectID,
		Token:     createUID().String(),
		LastLogin: time.Now()}

	err := sessionCollection.Insert(&us)

	if err != nil {
		panic(err)
	}
	return us.Token
}

func createUID() uuid.UUID {
	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return uid
	}
	return uid
}

// AuthenticateSession : Authenticates the token
func AuthenticateSession(token string) bool {
	var us UserSession
	err := sessionCollection.Find(bson.M{"token": token}).One(&us)

	if err == mgo.ErrNotFound {
		return false
	}
	return true
}
