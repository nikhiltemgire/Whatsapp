package servicehandlers

import (
	"net/http"
)

type PingHandler struct {
}

func (p PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p PingHandler) Get(r *http.Request) (string, int) {
	return "GET Called", 200
}

func (p PingHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

func (p PingHandler) Post(r *http.Request) (string, int) {
	return "POST Called", 200
}
