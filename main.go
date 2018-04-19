package main

import (
	"Whatsapp/servicehandlers"
	"log"
	"net/http"
)

func main() {

	p := servicehandlers.PingHandler{}
	a := servicehandlers.AuthenticateHandler{}
	u := servicehandlers.UserHandler{}

	http.Handle("/ping", p)
	http.Handle("/authenticate", a)
	http.Handle("/users", u)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// users := dao.GetAllUsers()
	// for _, u := range users {
	// 	fmt.Println(u)
	// }

	// user := dao.GetUserByEmail("alim@gmail.com")
	// fmt.Println(user)

	// fmt.Println(dao.CreateSession(user.ID))

	// fmt.Println(dao.AuthenticateSession("4ad48350-1dc9-4ac0-82ac-2e99709800de"))

	// fmt.Println(dao.CheckPassword("pass", "alim@gmail.com"))

}
