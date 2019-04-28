package twitter

import (
	"os"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/stevenleeg/go-twitter/twitter"
)

var isFollower map[string]bool

func sendTweet(_tweetID string, _username string, _text string) error {
	client, _, err := GetClient()
	if err != nil {
		return err
	}

	tweetID, err := strconv.ParseInt(_tweetID, 10, 64)
	if err != nil {
		return err
	}

	awaitRatelimit()

	_, _, err = client.Statuses.Update("@"+_username+" "+_text, &twitter.StatusUpdateParams{
		InReplyToStatusID: int64(tweetID),
	})

	if err != nil {
		return err
	}

	glog.Info("[Tweet] Responded to " + _username + " with: " + _text)

	if _text == os.Getenv("TWITTER_RESPONSE_ERROR") {
		alertAdmin(_username + os.Getenv("TWITTER_ADMIN_NOTIFY_PROBLEM"))
	}

	return nil
}

func sendDM(_user string, _userID string, _text string) error {
	client, _, err := GetClient()
	if err != nil {
		return err
	}

	awaitRatelimit()

	_, _, err = client.DirectMessages.EventsNew(&twitter.DirectMessageEventsNewParams{
		Event: &twitter.DirectMessageEvent{
			Type: "message_create",
			Message: &twitter.DirectMessageEventMessage{
				Target: &twitter.DirectMessageTarget{
					RecipientID: _userID,
				},
				Data: &twitter.DirectMessageData{
					Text: _text,
				},
			},
		},
	})
	if err != nil {
		return err
	}

	glog.Info("[DM] Responded to " + _user + " with: " + _text)

	if _text == os.Getenv("TWITTER_RESPONSE_ERROR") {
		alertAdmin(_user + os.Getenv("TWITTER_ADMIN_NOTIFY_PROBLEM"))
	}

	return nil
}

// getUser retrieves a twitter user object
func getUser(_screenName string) (*twitter.User, error) {
	_client, _, err := GetClient()
	if err != nil {
		return nil, err
	}

	awaitRatelimit()

	user, _, err := _client.Users.Show(&twitter.UserShowParams{
		ScreenName: _screenName,
	})
	if err != nil {
		return nil, err
	}

	user.Following = isFollower[user.IDStr]

	return user, nil
}

// GetFollowers will start a periodic watcher retrieving the
// followers of our account every 6 hours
func GetFollowers() {
	for {
		client, _, err := GetClient()
		if err != nil {
			glog.Error(err)
		}

		isFollower = make(map[string]bool)
		stillChecking := true
		amountOfFollowers := 0
		currentCursor := int64(-1)

		for stillChecking {
			awaitRatelimit()
			followers, _, err := client.Followers.List(&twitter.FollowerListParams{
				IncludeUserEntities: &[]bool{false}[0],
				Count:               200,
				Cursor:              currentCursor,
			})
			if err != nil {
				glog.Error(err)
			}
			for _, user := range followers.Users {
				isFollower[user.IDStr] = true
				amountOfFollowers++
			}
			if followers.NextCursor != 0 {
				currentCursor = followers.NextCursor
			} else {
				stillChecking = false
			}
		}
		glog.Info("Refreshed " + strconv.Itoa(amountOfFollowers) + " followers")
		time.Sleep(6 * 60 * time.Minute)
	}
}
