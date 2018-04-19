package dao

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var chatListCollection = GetChatListCollection(session)

// ChatListBox : View of one box in the ChatListchatList
// Contains information regarding Other user in chat with
type ChatListBox struct {
	ChatID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DisplayName    string
	DisplayPicture string
	IsContact      bool
}

// ChatList : Structure of Chatlist in MongoDB
type ChatList struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty" `
	ListOfChats []ChatListBox `json:"chatlist" bson:"chatlist" `
}

// GetChatList : Returns Chatlist of given userID
func GetChatList(userID string) ChatList {
	var chatlist ChatList

	err := chatListCollection.Find(
		bson.M{"OwnerID": userID}).One(&chatlist)

	if err == mgo.ErrNotFound {
		return ChatList{}
	} else if err != nil {
		panic(err)
	}

	return chatlist
}

// CreateEmptyChatList : Creates an empty chatList for a new User
func CreateEmptyChatList(ownerID string) {
	var chatlistbox []ChatListBox

	chatlist := ChatList{
		ID:          bson.ObjectIdHex(ownerID),
		ListOfChats: chatlistbox}

	err := chatListCollection.Insert(&chatlist)

	if err != nil {
		// panic(err)
		log.Println("Chatlist for client already exists")
	}

	fmt.Println(chatlist)
}
