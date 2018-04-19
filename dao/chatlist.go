package dao

import "gopkg.in/mgo.v2/bson"

// ChatList : Structure of Chatlist in MongoDB
type ChatList struct {
	ID          bson.ObjectId
	ClientID    bson.ObjectId
	ListOfChats []bson.ObjectId
}
