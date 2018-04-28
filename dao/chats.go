package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Chats : Chat Structure in MongoDB
type Chats struct {
	ID       bson.ObjectId   `json:"id" bson:"_id"`
	Members  []bson.ObjectId `json:"members" bson:"members"`
	Messages []bson.ObjectId `json:"messages" bson:"messages"`
}

var chatsCollection = GetChatsCollection(session)

// CreateChat : Creates a new Chat
func CreateChat(user1 bson.ObjectId,
	user2 bson.ObjectId) bson.ObjectId {

	var messages []bson.ObjectId
	members := []bson.ObjectId{user1, user2}
	id := bson.NewObjectId()

	err := chatsCollection.Insert(
		Chats{
			ID:       id,
			Members:  members,
			Messages: messages})

	if err != nil {
		panic(err)
	}
	return id
}

// GetChat : Returns the Chat Object associated with ChatID
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

// AddMessageToChat : Adds message to chat
func AddMessageToChat(chatID bson.ObjectId, msgID bson.ObjectId) {

	err := chatListCollection.Update(
		bson.M{"_id": chatID},
		bson.M{"$push": bson.M{"messages": msgID}})

	if err != nil {
		panic(err)
	}
}
