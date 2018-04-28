package servicehandlers

import (
	"Whatsapp/dao"
	"net/http"
)

// AuthenticateUser : Get token from header and authenticate
func AuthenticateUser(r *http.Request) bool {

	token := r.Header.Get("token")
	if token == "" {
		return false
	}

	return dao.AuthenticateSession(token)
}
