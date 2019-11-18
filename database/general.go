package database

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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

// NotificationProposalStarted struct
type NotificationProposalStarted struct {
	TwitterID         string    `json:"twitter_id" bson:"twitter_id"`
	LiveMode          bool      `json:"live_mode" bson:"live_mode"`
	RepositoryHash    string    `json:"repository_hash" bson:"repository_hash"`
	RepositoryName    string    `json:"repository_name" bson:"repository_name"`
	KNWVoteID         int       `json:"knw_vote_id" bson:"knw_vote_id"`
	KnowledgeLabel    string    `json:"knowledge_label" bson:"knowledge_label"`
	ProposerTwitterID string    `json:"proposer_twitter_id" bson:"proposer_twitter_id"`
	ProposerGithubID  string    `json:"proposer_github_handle" bson:"proposer_github_handle"`
	Description       string    `json:"description" bson:"description"`
	Identifier        string    `json:"identifier" bson:"identifier"`
	CommitUntil       time.Time `json:"commit_until" bson:"commit_until"`
	RevealUntil       time.Time `json:"reveal_until" bson:"reveal_until"`
}

// GetUser returns a user object when the user exists
func GetUser(_twitterID string) (*User, error) {
	var foundUsers User
	err := MgoRequest("users", func(c *mgo.Collection) error {
		return c.Find(bson.M{"twitter_id": _twitterID}).One(&foundUsers)
	})
	if err != nil {
		return nil, err
	}

	return &foundUsers, nil
}

// GetUsersForFeedback returns an array of users that
// passed the KYC but haven't provided feedback yet
func GetUsersForFeedback() ([]User, error) {
	var foundUsers []User
	err := MgoRequest("users", func(c *mgo.Collection) error {
		afterDays, _ := strconv.Atoi(os.Getenv("ASK_FOR_FEEDBACK_AFTER_DAYS"))
		daysAgo := time.Now().AddDate(0, 0, -afterDays)
		return c.Find(bson.M{
			"asked_for_feedback": false,
			"provided_feedback":  false,
			"passed_kyc_demo":    true,
			"date_of_contact":    bson.M{"$lte": daysAgo},
		}).All(&foundUsers)
	})
	if err != nil {
		return nil, err
	}

	return foundUsers, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(_existingUser User) error {
	where := bson.M{"twitter_id": _existingUser.TwitterID}
	change := bson.M{"$set": bson.M{
		"twitter_screen_name": _existingUser.TwitterScreenName,
		"eth_address":         _existingUser.ETHAddress,
		"passed_kyc_demo":     _existingUser.PassedKYCDemo,
		"passed_kyc_live":     _existingUser.PassedKYCLive,
		"skip_kyc":            _existingUser.SkipKYC,
		"date_of_contact":     _existingUser.DateOfContact,
		"used_client":         _existingUser.HasUsedClient,
		"asked_for_feedback":  _existingUser.HasBeenAskedForFeedback,
		"provided_feedback":   _existingUser.HasProvidedFeedback,
	}}
	err := MgoRequest("users", func(c *mgo.Collection) error {
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

	err = MgoRequest("users", func(c *mgo.Collection) error {
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
		mgoSession, mgoErr = mgo.DialWithInfo(
			&mgo.DialInfo{Addrs: []string{os.Getenv("MONGO_DB_ADDRESS")},
				Timeout:  10 * time.Second,
				Database: "twitterbot",
				Username: os.Getenv("MONGO_DB_USER"),
				Password: os.Getenv("MONGO_DB_PASSWORD")})
		if mgoErr != nil {
			return nil, mgoErr
		}
		mgoSession.SetMode(mgo.Monotonic, true)
	}
	return mgoSession.Clone(), nil
}

// MgoRequest will interact with the database
func MgoRequest(collection string, s func(*mgo.Collection) error) error {
	session, mgoErr := getSession()
	if mgoErr != nil {
		return mgoErr
	}
	defer session.Close()
	c := session.DB(databaseName).C(collection)
	return s(c)
}
