package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Message : Message Schema
type Message struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	ChatID     bson.ObjectId `json:"chat_id" bson:"chat_id"`
	SenderName string        `json:"sender_name" bson:"sender_name"`
	Type       int           `json:"type" bson:"type"`
	Status     int           `json:"status" bson:"status"`
	Content    string        `json:"content" bson:"content"`
}

var messagesCollection = GetMessagesCollection(session)

// GetMessageByID : Return Message associted by msgID
func GetMessageByID(msgID bson.ObjectId) Message {
	var msg Message

	err := messagesCollection.Find(bson.M{"_id": msgID}).One(&msg)

	if err == mgo.ErrNotFound {
		return msg
	} else if err != nil {
		log.Println("Error in GetUserProfileByID")
		panic(err)
	}
	return msg
}

// GetAllMessages returns all users from messages collection
// associted with ChatID
func GetAllMessages(chatID bson.ObjectId) []Message {
	var messages []Message

	err := messagesCollection.Find(bson.M{"chat_id": chatID}).All(&messages)
	if err != nil {
		panic(err)
	}

	return messages
}

// CreateMessage : Creates a Message instance in Collection
// Returns Message ID
func CreateMessage(chatID bson.ObjectId, senderName string,
	msgtype int, content string) bson.ObjectId {

	id := bson.NewObjectId()

	message := Message{
		ID:         id,
		ChatID:     chatID,
		SenderName: senderName,
		Type:       msgtype,
		Content:    content,
		Status:     0}

	err := messagesCollection.Insert(&message)

	if err != nil {
		panic(err)
	}

	return id
}
