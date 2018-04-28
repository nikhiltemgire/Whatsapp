package utils

import "net/http"

// AuthenticateUser : Get token from header and authenticate
func AuthenticateUser(r *http.Request) bool {

	token := r.Header["token"]
	if token == nil {
		return false
	}

	return true
}
