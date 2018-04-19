package main

import "Whatsapp/dao"

func main() {

	// fmt.Println(bson.ObjectIdHex("5a9cfc741b2a6104c061b24f"))

	dao.CreateEmptyChatList("5a9cfc741b2a6104c061b24f")

	// p := servicehandlers.PingHandler{}
	// a := servicehandlers.AuthenticateHandler{}
	// u := servicehandlers.UserHandler{}

	// http.Handle("/ping", p)
	// http.Handle("/authenticate", a)
	// http.Handle("/users", u)
	// log.Fatal(http.ListenAndServe(":8080", nil))

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
