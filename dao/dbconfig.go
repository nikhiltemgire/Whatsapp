package dao

import mgo "gopkg.in/mgo.v2"

const dbName = "go_authenticator"
const hostname = "127.0.0.1"

const (
	userdb    = "users"
	sessiondb = "sessions"
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

func getUserCollection(s *mgo.Session) *mgo.Collection {
	return s.DB(dbName).C(userdb)
}

func getSessionCollection(s *mgo.Session) *mgo.Collection {
	return s.DB(dbName).C(sessiondb)
}
