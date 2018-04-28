package servicehandlers

import (
	"Whatsapp/dao"
	"Whatsapp/utils"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// ChatsHandler : Restful Chats Handler
type ChatsHandler struct {
}

func (p ChatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : ChatsHandlers GET Method
func (p ChatsHandler) Get(r *http.Request) (string, int) {
	keys, ok := r.URL.Query()["chat_id"]

	if !ok || len(keys) < 1 {
		log.Println("Url Param 'chat_id' is missing")
		return "Bad Request", 400
	}

	chatID := bson.ObjectIdHex(keys[0])
	chats := dao.GetAllMessages(chatID)

	return utils.StructToJSONString(chats), 200
}

// Put : ChatsHandlers Put Method
func (p ChatsHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : ChatsHandlers Post Method
func (p ChatsHandler) Post(r *http.Request) (string, int) {

	return "POST Called", 200
}
