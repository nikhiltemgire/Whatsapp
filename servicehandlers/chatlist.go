package servicehandlers

import (
	"net/http"
)

type ChatListHandler struct {
}

func (p ChatListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p ChatListHandler) Get(r *http.Request) (string, int) {
	return "ChatList GET Called", 200
}

func (p ChatListHandler) Put(r *http.Request) (string, int) {
	return "ChatList PUT Called", 200
}

func (p ChatListHandler) Post(r *http.Request) (string, int) {

	return "ChatList Post Called", 200
}
