package main

import (
	"Whatsapp/servicehandlers"
	"log"
	"net/http"
)

func main() {

	// p := servicehandlers.PingHandler{}
	a := servicehandlers.AuthenticateHandler{}
	u := servicehandlers.UserHandler{}
	up := servicehandlers.UserProfileHandler{}
	cl := servicehandlers.ChatListHandler{}
	m := servicehandlers.MessageHandler{}

	// http.Handle("/ping", p)
	http.Handle("/chatservice/authenticate", a)
	http.Handle("/chatservice/users", u)
	http.Handle("/chatservice/userprofile", up)
	http.Handle("/chatservice/chatlist", cl)
	http.Handle("/chatservice/messages", m)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
