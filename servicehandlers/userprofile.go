package servicehandlers

import (

	"fmt"
	"net/http"
	_ "Whatsapp/dao"
	_ "encoding/json"
)

type UserProfileHandler struct {
}

func (p UserProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p UserProfileHandler) Get(r *http.Request) (string, int) {
	fmt.Println("GET params were:", r.URL.Query()); 
	return "GET Called", 200	
}

func (p UserProfileHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}



func (p UserProfileHandler) Post(r *http.Request) (string, int) {
	/*
	var payload dao.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	if err != nil {
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