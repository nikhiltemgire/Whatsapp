package dao

import mgo "gopkg.in/mgo.v2"

const dbName = "ChatService"
const hostname = "127.0.0.1"

const (
	userdb         = "users"
	userProfilesdb = "userProfiles"
	sessiondb      = "sessions"
	chatListdb     = "chatlist"
	chatsdb        = "chats"
	messagesdb     = "messages"
)

var session = getDatabaseSession()

// defer closeDatabaseSession(session)

func getDatabaseSession() *mgo.Session {
	s, err := mgo.Dial(hostname)

	if err != nil {
		panic(err)
	}
	return s
}

func closeDatabaseSession(s *mgo.Session) {
	s.Close()
}

// GetUserCollection : Returns User Collection handler
func GetUserCollection(s *mgo.Session) *mgo.Collection {
	return s.DB(dbName).C(userdb)
}

// GetUserProfilesCollection : Returns Chat List Collection handler
func GetUserProfilesCollection(s *mgo.Session) *mgo.Collection {
	return s.DB(dbName).C(userProfilesdb)
}

// GetSessionCollection : Returns Session Collection handler
func GetSessionCollection(s *mgo.Session) *mgo.Collection {
	return s.DB(dbName).C(sessiondb)
}

// GetChatListCollection : Returns Chat List Collection handler
func GetChatListCollection(s *mgo.Session) *mgo.Collection {
	return s.DB(dbName).C(chatListdb)
}

// GetChatsCollection : Returns Chat List Collection handler
func GetChatsCollection(s *mgo.Session) *mgo.Collection {
	return s.DB(dbName).C(chatsdb)
}

// GetMessagesCollection : Returns Messages Collection handler
func GetMessagesCollection(s *mgo.Session) *mgo.Collection {
	return s.DB(dbName).C(messagesdb)
}
