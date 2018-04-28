package dao

import (

	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserProfile struct {
	ID       	 	bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
	Email    		string        	`json:"email" bson:"email"`
	DisplayName 	string        	`json:"display_name" bson:"display_name"`
	DisplayPicture 	string		   	`json:"display_picture" bson:"display_picture"`
	LastSeen		time.Time		`json:"last_seen" bson:"last_seen"`			
	About			string			`json:"about" bson:"about"`
	
}


var userProfileCollection = getUserProfileCollection(session)


// GetAllUser returns all users from users collection
/*func GetAllUsers() []User {
	var users []User

	err := userCollection.Find(nil).All(&users)
	if err != nil {
		panic(err)
	}

	return users
}*/

// GetUserProfileByEmail returns the user profile if exists
func GetUserProfileByEmail(email string) UserProfile {
	var userProfile UserProfile

	err := userProfileCollection.Find(bson.M{"email": email}).One(&userProfile)

	if err == mgo.ErrNotFound {
		fmt.Println("Not Found")
		return userProfile
	} else if err != nil {
		panic(err)
	}

	return userProfile
}


// GetUserProfileByID returns the user profile if exists
func GetUserProfileByID(userID bson.ObjectId) UserProfile {
	var userProfile UserProfile

	err := userProfileCollection.Find(bson.M{"_id": userID}).One(&userProfile)

	if err == mgo.ErrNotFound {
		fmt.Println("Not Found")
		return userProfile
	} else if err != nil {
		panic(err)
	}

	return userProfile
}

// CheckPassword compares the entered password with pass of user
/*func AuthenticateUser(email string, password string) bool {
	user := GetUserByEmail(email)

	if password == user.Password {
		return true
	}
	return false
}*/

	/*Email    		string        	`json:"email" bson:"email"`
	DisplayName 	string        	`json:"display_name" bson:"display_name"`
	DisplayPicture 	string		   	`json:"display_picture" bson:"display_picture"`
	LastSeen		time.Time		`json:"last_seen" bson:"last_seen"`			
	About			string			`json:"about" bson:"about"`
	*/

func CreateUserProfile(email string, password string) {
	var user User
	user.Email = email
	user.Password = password
	err := userCollection.Insert(&user)
	if err != nil {
		panic(err)
	}
}