package servicehandlers

import (
	"Whatsapp/dao"
	"Whatsapp/utils"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// UserProfileHandler : UserProfile Handler
type UserProfileHandler struct {
}

func (p UserProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

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
func (p UserProfileHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : UserProfileHandler Post Method
func (p UserProfileHandler) Post(r *http.Request) (string, int) {

	var payload dao.UserProfile
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
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
