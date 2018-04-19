package servicehandlers

import (
	"fmt"
	"net/http"
)

// HttpServiceHandler : Restful Http Handler
type HTTPServiceHandler interface {
	Get(*http.Request) (string, int)
	Put(*http.Request) (string, int)
	Post(*http.Request) (string, int)
}

func methodRouter(p HTTPServiceHandler, w http.ResponseWriter, r *http.Request) {

	var response string
	var status int
	if r.Method == "GET" {
		response, status = p.Get(r)
	} else if r.Method == "PUT" {
		response, status = p.Put(r)
	} else if r.Method == "POST" {
		response, status = p.Post(r)
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, response)

}
