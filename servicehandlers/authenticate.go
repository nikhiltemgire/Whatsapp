package servicehandlers

import (
	"Whatsapp/dao"
	"encoding/json"
	"net/http"
)

// AuthenticateHandler : AuthenticateHandler of Chatservice
type AuthenticateHandler struct {
}

func (p AuthenticateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : Authenciate Get Method
func (p AuthenticateHandler) Get(r *http.Request) (string, int) {
	return "Method Not Allowed", 200
}

// Put : Authenciate Put Method
func (p AuthenticateHandler) Put(r *http.Request) (string, int) {
	return "Method Not Allowed", 200
}

// Post : Authenciate Post Method
func (p AuthenticateHandler) Post(r *http.Request) (string, int) {

	var payload dao.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
		panic(err)
	}

	if !dao.AuthenticateUser(payload.Email, payload.Password) {
		return "Invalid Email or Password", 400
	}
	user := dao.GetUserByEmail(payload.Email)
	token := dao.CreateSession(user.ID)

	return token, 200
}
