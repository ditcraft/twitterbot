package twitter

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/marvinkruse/dit-twitterbot/database"
	"github.com/marvinkruse/dit-twitterbot/ethereum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stevenleeg/go-twitter/twitter"
)

var followerThreshold int

func handleNewTweet(_tweetID string, _user string, _userID string, _followerCount int, _text string) {
	twitterUser, err := getUser(_user)
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Infof("[Tweet] %s (Follower: %t, FollowerCount: %d): %s\n", _user, twitterUser.Following, _followerCount, _text)

	if !twitterUser.Following {
		err := sendTweet(_tweetID, _user, os.Getenv("TWITTER_RESPONSE_IS_NO_FOLLOWER"))
		if err != nil {
			glog.Error(err)
		}
		return
	}

	ethAddress, containsAddress := containsETHAddress(_text)

	if containsAddress {
		user, err := database.GetUser(_userID)
		if err != nil && !strings.Contains(err.Error(), "not found") {
			glog.Error(err)
		}
		if user == nil {
			passedKYC := doKYC(twitterUser)
			if !passedKYC {
				err := sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_KYC_FAIL_TWEET"))
				if err != nil {
					glog.Error(err)
				}
				return
			}
		}

		answer := handleETHRequest(_userID, _user, ethAddress, false)
		err = sendTweet(_tweetID, _user, answer)
		if err != nil {
			glog.Error(err)
		}
	}
}

func handleNewDM(_user string, _userID string, _followerCount int, _text string) {
	twitterUser, err := getUser(_user)
	if err != nil {
		err := sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_ERROR"))
		if err != nil {
			glog.Error(err)
		}
		return
	}

	glog.Infof("[DM] %s (Follower: %t, FollowerCount: %d): %s\n", _user, twitterUser.Following, _followerCount, _text)

	if !twitterUser.Following {
		err := sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_IS_NO_FOLLOWER"))
		if err != nil {
			glog.Error(err)
		}
		return
	}

	dbUser, err := database.GetUser(_userID)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		glog.Error(err)
	}
	if dbUser == nil {
		passedKYC := doKYC(twitterUser)
		if !passedKYC {
			err := sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_KYC_FAIL_DM"))
			if err != nil {
				glog.Error(err)
			}
			return
		}
	}

	wasCommand, err := handleCommand(_user, _userID, _text)
	if err != nil {
		glog.Error(err)
		err := sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_ERROR"))
		if err != nil {
			glog.Error(err)
		}
		return
	}

	if !wasCommand {
		ethAddress, containsAddress := containsETHAddress(_text)

		if containsAddress {
			answer := handleETHRequest(_userID, _user, ethAddress, true)
			err := sendDM(_user, _userID, answer)
			if err != nil {
				glog.Error(err)
			}
		}
	}
}

func handleETHRequest(_userID string, _userName string, _ethAddress string, _viaDM bool) string {
	userObject, err := database.GetUser(_userID)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return os.Getenv("TWITTER_RESPONSE_ERROR")
	}
	if userObject != nil && userObject.WasFunded {
		if _viaDM {
			return os.Getenv("TWITTER_RESPONSE_ALREADY_FUNDED_DM")
		}
		return os.Getenv("TWITTER_RESPONSE_ALREADY_FUNDED_TWEET")
	}

	if userObject == nil {
		userObject = &database.User{
			TwitterID:         _userID,
			TwitterScreenName: _userName,
			ETHAddress:        _ethAddress,
			DateOfContact:     time.Now(),
			WasFunded:         false,
		}
		err := database.CreateUser(*userObject)
		if err != nil {
			return os.Getenv("TWITTER_RESPONSE_ERROR")
		}
	}

	if !userObject.WasFunded {
		err := ethereum.SendEther(_ethAddress, 1)
		if err != nil {
			return os.Getenv("TWITTER_RESPONSE_ERROR")
		}

		userObject.WasFunded = true

		if _ethAddress != userObject.ETHAddress {
			userObject.ETHAddress = _ethAddress
		}

		err = database.UpdateUser(*userObject)
		if err != nil {
			glog.Error(err)
		}
	}

	return os.Getenv("TWITTER_RESPONSE_SUCCESS")
}

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

	return user, nil
}

// hasEnoughFollowers indicates whether the user has enough followers
// according to our threshold
func hasEnoughFollowers(_user *twitter.User, _amountOfFollowers int) bool {
	if _user.FollowersCount < _amountOfFollowers {
		return false
	}
	return true
}

func doKYC(_user *twitter.User) bool {
	if _user.Verified {
		return true
	}

	followerCount, _ := strconv.Atoi(os.Getenv("TWITTER_FOLLOWER_THRESHOLD"))
	if _user.FollowersCount < followerCount {
		return false
	}

	statusesCount, _ := strconv.Atoi(os.Getenv("TWITTER_FOLLOWER_THRESHOLD"))
	if _user.StatusesCount < statusesCount {
		return true
	}

	return true
}

func containsETHAddress(_text string) (string, bool) {
	if strings.Contains(_text, "0x") {
		start := strings.Index(_text, "0x")
		if len(_text[start:]) >= 42 {
			ethAddress := _text[start : start+42]
			if common.IsHexAddress(ethAddress) {
				return ethAddress, true
			}
		}
	}
	return "", false
}

func handleCommand(_user string, _userID string, _text string) (bool, error) {
	if strings.HasPrefix(_text, "!") {
		var err error
		switch {
		case strings.HasPrefix(_text, "!commands"):
			err = sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_COMMAND_LIST"))
		case strings.HasPrefix(_text, "!problem"):
			err = sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_COMMAND_PROBLEM"))
		default:
			err = sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_COMMAND_NOT_FOUND"))
		}
		if err != nil {
			return true, err
		}
		return true, nil
	}
	return false, nil
}
