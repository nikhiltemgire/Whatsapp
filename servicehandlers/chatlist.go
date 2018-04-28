package servicehandlers

import (
	"Whatsapp/dao"
	"Whatsapp/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ChatListHandler : Restful ChatList Handler
type ChatListHandler struct {
}

func (p ChatListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : ChatListHandler Get Method
func (p ChatListHandler) Get(r *http.Request) (string, int) {

	token := r.Header.Get("token")

	if token == "" {
		return "Bad Request", 400
	}

	if !AuthenticateUser(r) {
		return "Unauthorized Access", 401
	}

	userID := dao.GetUserIDByToken(token)
	chatList := dao.GetChatList(userID)

	return utils.StructToJSONString(chatList), 200
}

// Put : ChatListHandler Put Method
func (p ChatListHandler) Put(r *http.Request) (string, int) {

	token := r.Header.Get("token")

	if token == "" {
		return "Bad Request", 400
	}

	if !AuthenticateUser(r) {
		return "Unauthorized Access", 401
	}

	decoder := json.NewDecoder(r.Body)
	var payload struct {
		Email string `json:"email" bson:"email"`
	}
	err := decoder.Decode(&payload)
	if err == io.EOF {
		return "Bad Request", 400
	}
	if err != nil {
		fmt.Println(payload, r.Body)
		panic(err)
	}

	user1 := dao.GetUserIDByToken(token)
	u1 := dao.GetUserProfileByID(user1)
	u2 := dao.GetUserProfileByEmail(payload.Email)

	if (u1 == dao.UserProfile{} || u2 == dao.UserProfile{}) {
		fmt.Println(u1, u2)
		return "User Does not Exist.", 422
	}

	s := dao.AddChat(u1, u2)

	if s {
		return "", 200
	}
	return "Internal Server Error", 500
}

// Post : UserProfileHandler Post Method
func (p ChatListHandler) Post(r *http.Request) (string, int) {
	return "ChatList Post Called", 200
}
