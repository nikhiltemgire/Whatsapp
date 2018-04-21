package dao

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var chatListCollection = GetChatListCollection(session)

// ChatListEntry : View of one box in the ChatListchatList
// Contains information regarding Other user in chat with
type ChatListEntry struct {
	ChatID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DisplayName    string
	DisplayPicture string
}

// ChatList : Structure of Chatlist in MongoDB
type ChatList struct {
	UserID      bson.ObjectId   `json:"user_id" bson:"_id,omitempty" `
	ListOfChats []ChatListEntry `json:"chatlist" bson:"chatlist" `
}

// GetChatList : Returns Chatlist of given userID
func GetChatList(userID bson.ObjectId) ChatList {
	var chatlist ChatList

	err := chatListCollection.Find(
		bson.M{"UserID": userID}).One(&chatlist)

	if err == mgo.ErrNotFound {
		return ChatList{}
	} else if err != nil {
		panic(err)
	}

	return chatlist
}

// CreateEmptyChatList : Creates an empty chatList for a new User
func CreateEmptyChatList(ownerID bson.ObjectId) {
	var chatlistbox []ChatListEntry

	chatlist := ChatList{
		UserID:      ownerID,
		ListOfChats: chatlistbox}

	err := chatListCollection.Insert(&chatlist)

	if err != nil {
		// panic(err)
		log.Println("Chatlist for client already exists")
	}

	fmt.Println(chatlist)
}

// CreateChatListEntry : Creates a ChatListEnry Object
func CreateChatListEntry(ID bson.ObjectId,
	displayName string, displayPicture string) ChatListEntry {

	return ChatListEntry{ID, displayName, displayPicture}
}

// AddChat : Adds a chat instance between user1 and user2
func AddChat(user1 UserProfile, user2 UserProfile) bool {

	chatID := bson.NewObjectId()

	user1ListEntry := CreateChatListEntry(chatID, user2.DisplayName, "")
	user2ListEntry := CreateChatListEntry(chatID, user1.DisplayName, "")

	err := chatListCollection.Update(
		bson.M{"_id": user1.ID},
		bson.M{"$push": bson.M{"chatlist": user1ListEntry}})

	if err != nil {
		return false
	}

	err = chatListCollection.Update(
		bson.M{"_id": user2.ID},
		bson.M{"$push": bson.M{"chatlist": user2ListEntry}})

	if err != nil {
		return false
	}
	return true
}
