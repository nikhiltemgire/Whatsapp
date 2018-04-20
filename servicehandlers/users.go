package servicehandlers

import (
	"Whatsapp/dao"
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type UserHandler struct {
}

func (p UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : UserHandler Get Method
func (p UserHandler) Get(r *http.Request) (string, int) {
	return "GET Called", 200
}

// Put : UserHandler Put Method
func (p UserHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : UserHandler Post Method
func (p UserHandler) Post(r *http.Request) (string, int) {

	var payload dao.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
		// return "BAD Request", 400
		panic(err)
	}

	user := dao.GetUserByEmail(payload.Email)
	if (dao.User{}) != user {
		return "Unprocessable Entity", 422
	}

	id := bson.NewObjectId()
	success := dao.CreateUser(id, payload.Email, payload.Password)

	if success {
		return id.Hex(), 200
	}
	return "Internal Server Error", 500

}
