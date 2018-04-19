package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Chats : Chat Structure in MongoDB
type Chats struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	User1    bson.ObjectId
	User2    bson.ObjectId
	Messages []bson.ObjectId
}

var chatsCollection = GetChatsCollection(session)

// CreateChat : Creates a new Chat
func CreateChat(user1 bson.ObjectId,
	user2 bson.ObjectId) bson.ObjectId {

	var messages []bson.ObjectId
	id := bson.NewObjectId()

	err := chatsCollection.Insert(
		Chats{
			ID:       id,
			User1:    user1,
			User2:    user2,
			Messages: messages})

	if err != nil {
		panic(err)
	}
	return id
}

//GetChat : Returns the Chat Object associated with ChatID
func GetChat(chatID bson.ObjectId) Chats {
	var chat Chats

	err := chatsCollection.Find(bson.M{"_id": chatID}).One(&chat)

	if err == mgo.ErrNotFound {
		return chat
	} else if err != nil {
		panic(err)
	}

	return chat
}
