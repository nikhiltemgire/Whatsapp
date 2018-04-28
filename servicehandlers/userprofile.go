package servicehandlers

import (
<<<<<<< HEAD

	"fmt"
	"net/http"
	_ "Whatsapp/dao"
	_ "encoding/json"
)

=======
	"Whatsapp/dao"
	"Whatsapp/utils"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// UserProfileHandler : UserProfile Handler
>>>>>>> 295d5d29d9ff944eec7cd64695e5eae6ef32c988
type UserProfileHandler struct {
}

func (p UserProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

<<<<<<< HEAD
func (p UserProfileHandler) Get(r *http.Request) (string, int) {
	fmt.Println("GET params were:", r.URL.Query()); 
	return "GET Called", 200	
}

=======
// Get : UserProfileHandler Get Method
func (p UserProfileHandler) Get(r *http.Request) (string, int) {

	if !AuthenticateUser(r) {
		return "Unauthorized Access", 401
	}

	keys, ok := r.URL.Query()["user_id"]

	if !ok || len(keys) < 1 {
		log.Println("Url Param 'user_id' is missing")
		return "Bad Request", 400
	}
	userID := bson.ObjectIdHex(keys[0])
	userProfile := dao.GetUserProfileByID(userID)

	if (userProfile == dao.UserProfile{}) {
		// fmt.Println(userID)
		return "Profile Not Found", 404
	}

	return utils.StructToJSONString(userProfile), 200
}

// Put : UserProfileHandler Put Method
>>>>>>> 295d5d29d9ff944eec7cd64695e5eae6ef32c988
func (p UserProfileHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

<<<<<<< HEAD


func (p UserProfileHandler) Post(r *http.Request) (string, int) {
	/*
	var payload dao.User
=======
// Post : UserProfileHandler Post Method
func (p UserProfileHandler) Post(r *http.Request) (string, int) {

	var payload dao.UserProfile
>>>>>>> 295d5d29d9ff944eec7cd64695e5eae6ef32c988
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
<<<<<<< HEAD
		//return "BAD Request", 400
		panic(err)
	}


	user := dao.GetUserByEmail(payload.Email)
	if (dao.User{}) != user {					// User exists
		return "Unprocessable Entity" , 422
	}


	dao.CreateUser(payload.Email, payload.Password)
	// return token, 200
	*/
	return "POST Called", 200 
}
=======
		panic(err)
	}

	if payload.ID == bson.ObjectId("") {
		return "Bad Request", 400
	}

	user := dao.GetUserProfileByEmail(payload.Email)
	if (dao.UserProfile{}) != user { // User Exists
		return "Unprocessable Entity", 422
	}

	success := dao.CreateUserProfile(payload.ID, payload.Email,
		payload.DisplayName, payload.About)

	dao.CreateEmptyChatList(payload.ID)

	if success {
		return "ok", 200
	}
	return "Internal Server Error", 500
}
>>>>>>> 295d5d29d9ff944eec7cd64695e5eae6ef32c988
