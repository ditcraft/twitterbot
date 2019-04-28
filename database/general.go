package database

import (
	"errors"
	"os"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mgoSession *mgo.Session
var databaseAddress string
var databaseName = "twitterbot"

// User struct for a twitter user
type User struct {
	TwitterID               string    `bson:"twitter_id"`
	TwitterScreenName       string    `bson:"twitter_screen_name"`
	ETHAddress              string    `bson:"eth_address"`
	PassedKYCDemo           bool      `bson:"passed_kyc_demo"`
	PassedKYCLive           bool      `bson:"passed_kyc_live"`
	SkipKYC                 bool      `bson:"skip_kyc"`
	DateOfContact           time.Time `bson:"date_of_contact"`
	HasUsedClient           bool      `bson:"used_client"`
	HasBeenAskedForFeedback bool      `bson:"asked_for_feedback"`
	HasProvidedFeedback     bool      `bson:"provided_feedback"`
}

// GetUser returns a user object when the user exists
func GetUser(_twitterID string) (*User, error) {
	var foundUsers User
	err := mgoRequest("users", func(c *mgo.Collection) error {
		return c.Find(bson.M{"twitter_id": _twitterID}).One(&foundUsers)
	})
	if err != nil {
		return nil, err
	}

	return &foundUsers, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(_existingUser User) error {
	where := bson.M{"twitter_id": _existingUser.TwitterID}
	change := bson.M{"$set": bson.M{
		"twitter_screen_name": _existingUser.TwitterScreenName,
		"eth_address":         _existingUser.ETHAddress,
		"passed_kyc_demo":     _existingUser.PassedKYCDemo,
		"passed_kyc_live":     _existingUser.PassedKYCLive,
		"date_of_contact":     _existingUser.DateOfContact,
		"used_client":         _existingUser.HasUsedClient,
		"asked_for_feedback":  _existingUser.HasBeenAskedForFeedback,
	}}
	err := mgoRequest("users", func(c *mgo.Collection) error {
		return c.Update(where, change)
	})
	if err != nil {
		return err
	}

	return nil
}

// CreateUser stores a new user in the database
func CreateUser(_newUser User) error {
	user, err := GetUser(_newUser.TwitterID)
	if user != nil || (err != nil && !(strings.Contains(err.Error(), "not found"))) {
		return errors.New("Failed to check whether this user already exists")
	}

	err = mgoRequest("users", func(c *mgo.Collection) error {
		return c.Insert(_newUser)
	})
	if err != nil {
		return err
	}

	return nil
}

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
