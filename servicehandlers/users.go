package servicehandlers

import (
	"fmt"
	"net/http"
	"Whatsapp/dao"
	"encoding/json"
)

type UserHandler struct {
}

func (p UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p UserHandler) Get(r *http.Request) (string, int) {
	return "GET Called", 200	
}

func (p UserHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

func (p UserHandler) Post(r *http.Request) (string, int) {
	
	var payload dao.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
		// return "BAD Request", 400
		panic(err)
	}

	// if !dao.AuthenticateUser(payload.Email, payload.Password) {
	// 	return "Invalid Email or Password", 400
	// }

	user := dao.GetUserByEmail(payload.Email)
	if (dao.User{}) != user {
		return "Unprocessable Entity" , 422
	}


	// return token, 200
	
	return "POST Called", 200
}