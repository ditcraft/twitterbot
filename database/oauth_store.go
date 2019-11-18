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

// OAuthToken stores the data required to control the two bots which act as the
// frontend for our TCR. This table should only have two rows, one for the TCR
// Bot and one for the VIP bot.
type OAuthToken struct {
	TwitterHandle    string    `bson:"twitter_handle"`
	TwitterID        int64     `bson:"twitter_id"`
	OAuthToken       string    `bson:"oauth_token"`
	OAuthTokenSecret string    `bson:"oauth_token_secret"`
	CreatedAt        time.Time `bson:"created_at"`
}

// CreateOAuthToken inserts a new OAuth token into the database
func CreateOAuthToken() error {
	twitterID, _ := strconv.Atoi(os.Getenv("TWITTER_ID"))
	newToken := OAuthToken{
		TwitterHandle:    os.Getenv("TWITTER_HANDLE"),
		TwitterID:        int64(twitterID),
		OAuthToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		OAuthTokenSecret: os.Getenv("TWITTER_ACCESS_SECRET"),
		CreatedAt:        time.Now(),
	}

	err := MgoRequest("oauth_session", func(c *mgo.Collection) error {
		return c.Insert(newToken)
	})
	if err != nil {
		return err
	}

	return nil
}

// FindOAuthTokenByHandle returns an OAuth token given a twitter handle
func FindOAuthTokenByHandle(handle string) (*OAuthToken, error) {
	handle = strings.ToLower(handle)

	var oAuthTokenFromDB []OAuthToken
	err := MgoRequest("oauth_session", func(c *mgo.Collection) error {
		return c.Find(nil).Sort("-created_when").All(&oAuthTokenFromDB)
	})
	if err != nil {
		return nil, err
	}
	if len(oAuthTokenFromDB) == 0 {
		return nil, errors.New("No session found")
	}

	return &oAuthTokenFromDB[0], nil
}

// Save saves the OAuthToken to the database
func (token *OAuthToken) Save() error {
	err := MgoRequest("oauth_session", func(c *mgo.Collection) error {
		_, err := c.RemoveAll(bson.M{"twitter_handle": bson.M{"$eq": token.TwitterHandle}})
		return err
	})
	if err != nil {
		return err
	}

	err = MgoRequest("oauth_session", func(c *mgo.Collection) error {
		return c.Insert(token)
	})
	if err != nil {
		return err
	}

	return nil
}
