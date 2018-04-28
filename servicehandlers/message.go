package servicehandlers

import (
	"Whatsapp/dao"
	"Whatsapp/utils"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// MessageHandler : Restful Chats Handler
type MessageHandler struct {
}

func (p MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : MessageHandler GET Method
func (p MessageHandler) Get(r *http.Request) (string, int) {

	if !AuthenticateUser(r) {
		return "[Messages] Unauthorized Access", 401
	}

	keys, ok := r.URL.Query()["chat_id"]

	if !ok || len(keys) < 1 {
		log.Println("Url Param 'chat_id' is missing")
		return "Bad Request", 400
	}

	chatID := bson.ObjectIdHex(keys[0])
	chat := dao.GetAllMessages(chatID)

	return utils.StructToJSONString(chat), 200
}

// Put : MessageHandler Put Method
func (p MessageHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : MessageHandler Post Method
func (p MessageHandler) Post(r *http.Request) (string, int) {

	token := r.Header.Get("token")
	if token == "" {
		return "Bad Request", 400
	}

	if !AuthenticateUser(r) {
		return "Unauthorized Access", 401
	}

	var payload dao.Message
	decoder := json.NewDecoder(r.Body)

	if r.Body == nil {
		return "Bad Request", 400
	}
	err := decoder.Decode(&payload)

	if err != nil {
		panic(err)
	}

	id := dao.CreateMessage(payload.ChatID, payload.SenderName, payload.Type, payload.Content)

	dao.AddMessageToChat(payload.ChatID, id)

	return "Message Sent", 200
}
