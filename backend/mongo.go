package backend

import "labix.org/v2/mgo"

func DefaultDial() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return session
}
