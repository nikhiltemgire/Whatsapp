package dao

import (
	"gopkg.in/mgo.v2/bson"
)

// Chats : Chat Structure in MongoDB
type Chats struct {
	ID       bson.ObjectId
	User1    bson.ObjectId
	User2    bson.ObjectId
	Messages []bson.ObjectId
}
