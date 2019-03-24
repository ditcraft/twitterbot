package database

import (
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session
var databaseAddress string
var databaseName = "twitterbot"

func getSession() (*mgo.Session, error) {
	if mgoSession == nil {
		var mgoErr error
		mgoSession, mgoErr = mgo.DialWithTimeout(os.Getenv("MONGO_DB_ADDRESS"), 5*time.Second)
		if mgoErr != nil {
			return nil, mgoErr
		}
		mgoSession.SetMode(mgo.Monotonic, true)
	}
	return mgoSession.Clone(), nil
}

func mgoRequest(collection string, s func(*mgo.Collection) error) error {
	session, mgoErr := getSession()
	if mgoErr != nil {
		return mgoErr
	}
	defer session.Close()
	c := session.DB(databaseName).C(collection)
	return s(c)
}
