package dao

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

var userProfileCollection = GetUserProfilesCollection(session)

// UserProfile : UserProfile Structure
type UserProfile struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Email       string        `json:"email" bson:"email"`
	DisplayName string        `json:"display_name" bson:"display_name"`
	LastSeen    time.Time     `json:"last_seen" bson:"last_seen"`
	About       string        `json:"about" bson:"about"`
}

// CreateUserProfile : Creates a profile associated with the userID
func CreateUserProfile(userID bson.ObjectId, email string,
	displayName string, about string) bool {

	err := userProfileCollection.Insert(
		&UserProfile{
			ID:          userID,
			Email:       email,
			DisplayName: displayName,
			About:       about,
			LastSeen:    time.Now()})

	if err != nil {
		log.Println("User Profile Insert error: ", err)
		return false
	}

	return true
}

// GetUserProfileByID : Returns the User Profile Associated with userID
func GetUserProfileByID(userID bson.ObjectId) UserProfile {
	var up UserProfile

	err := userProfileCollection.Find(bson.M{"_id": userID}).One(&up)

	if err == mgo.ErrNotFound {
		// fmt.Print("lel")
		return up
	} else if err != nil {
		log.Println("Error in GetUserProfileByID")
		panic(err)
	}
	return up
}

// GetUserProfileByEmail : Returns the User Profile Associated with email
func GetUserProfileByEmail(email string) UserProfile {
	var up UserProfile

	err := userProfileCollection.Find(bson.M{"email": email}).One(&up)

	if err == mgo.ErrNotFound {
		return up
	} else if err != nil {
		log.Println("Error in GetUserProfileByEmail")
		panic(err)
	}
	return up
}
