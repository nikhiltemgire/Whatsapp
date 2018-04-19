package servicehandlers

import (
	"eagoauthentication/dao"
	"encoding/json"
	"net/http"
)

type AuthenticateHandler struct {
}

func (p AuthenticateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p AuthenticateHandler) Get(r *http.Request) (string, int) {
	return "AUTHENTICATE GET Called", 200
}

func (p AuthenticateHandler) Put(r *http.Request) (string, int) {
	return "AUTHENTICATE PUT Called", 200
}

func (p AuthenticateHandler) Post(r *http.Request) (string, int) {

	type authenticateBody struct {
		Email    string
		Password string
	}

	var payload authenticateBody
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
