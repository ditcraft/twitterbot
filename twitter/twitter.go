package twitter

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/EagleChen/mapmutex"
	"github.com/golang/glog"
	"github.com/marvinkruse/dit-twitterbot/database"
	"github.com/marvinkruse/dit-twitterbot/ethereum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stevenleeg/go-twitter/twitter"
)

var followerThreshold int

// PerUserMutex will block on a userid basis to prevent spamming from one user
var PerUserMutex *mapmutex.Mutex

func handleNewTweet(_tweetID string, _user string, _userID string, _followerCount int, _text string) {
	gotLock := false
	for !gotLock {
		gotLock = PerUserMutex.TryLock(_userID)
		defer PerUserMutex.Unlock(_userID)
	}

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

		passedFullKYC := false

		if user != nil && user.SkipKYC == true {
			passedFullKYC = true
		} else if user == nil || !user.PassedKYCDemo || !user.PassedKYCLive {
			var passedKYC bool
			passedKYC, passedFullKYC = doKYC(twitterUser)
			if !passedKYC {
				err := sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_KYC_FAIL_TWEET"))
				if err != nil {
					glog.Error(err)
				}
				alertAdmin(_user + os.Getenv("TWITTER_ADMIN_NOTIFY_NOKYC"))
				return
			} else if !passedFullKYC {
				alertAdmin(_user + os.Getenv("TWITTER_ADMIN_NOTIFY_HALFKYC"))
			}
		}

		answer := handleKYCApprove(_userID, _user, ethAddress, false, passedFullKYC)
		err = sendTweet(_tweetID, _user, answer)
		if err != nil {
			glog.Error(err)
		}
	}
}

func handleNewDM(_user string, _userID string, _followerCount int, _text string) {
	gotLock := false
	for !gotLock {
		gotLock = PerUserMutex.TryLock(_userID)
		defer PerUserMutex.Unlock(_userID)
	}

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
			dbUser, err := database.GetUser(_userID)
			if err != nil && !strings.Contains(err.Error(), "not found") {
				glog.Error(err)
			}

			passedFullKYC := false

			if dbUser != nil && dbUser.SkipKYC == true {
				passedFullKYC = true
			} else if dbUser == nil || !dbUser.PassedKYCDemo || !dbUser.PassedKYCLive {
				var passedKYC bool
				passedKYC, passedFullKYC = doKYC(twitterUser)
				if !passedKYC {
					err := sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_KYC_FAIL_DM"))
					if err != nil {
						glog.Error(err)
					}
					alertAdmin(_user + os.Getenv("TWITTER_ADMIN_NOTIFY_NOKYC"))
					return
				} else if !passedFullKYC {
					alertAdmin(_user + os.Getenv("TWITTER_ADMIN_NOTIFY_HALFKYC"))
				}
			}

			answer := handleKYCApprove(_userID, _user, ethAddress, true, passedFullKYC)
			err = sendDM(_user, _userID, answer)
			if err != nil {
				glog.Error(err)
			}
		}
	}
}

func handleKYCApprove(_userID string, _userName string, _ethAddress string, _viaDM bool, _liveKYC bool) string {
	userObject, err := database.GetUser(_userID)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return os.Getenv("TWITTER_RESPONSE_ERROR")
	}
	if userObject != nil && (userObject.PassedKYCLive || (!_liveKYC && !userObject.PassedKYCLive)) && userObject.PassedKYCDemo {
		if _viaDM {
			return os.Getenv("TWITTER_RESPONSE_ALREADY_KYCED_DM")
		}
		return os.Getenv("TWITTER_RESPONSE_ALREADY_KYCED_TWEET")
	}

	if userObject == nil {
		userObject = &database.User{
			TwitterID:         _userID,
			TwitterScreenName: _userName,
			ETHAddress:        _ethAddress,
			DateOfContact:     time.Now(),
			PassedKYCDemo:     false,
			PassedKYCLive:     false,
		}
		err := database.CreateUser(*userObject)
		if err != nil {
			glog.Error(err)
			return os.Getenv("TWITTER_RESPONSE_ERROR")
		}
	}

	if (!userObject.PassedKYCDemo || (!userObject.PassedKYCLive && _liveKYC)) && _viaDM {
		err := sendDM(_userName, _userID, os.Getenv("TWITTER_RESPONSE_STARTING_ETHEREUM_TXS"))
		if err != nil {
			glog.Error(err)
		}
	}

	if !userObject.PassedKYCDemo {
		ethereum.Mutex.Lock()
		defer ethereum.Mutex.Unlock()

		err := ethereum.SendXDaiCent(_ethAddress)
		if err != nil {
			glog.Error(err)
			return os.Getenv("TWITTER_RESPONSE_ERROR")
		}

		err = ethereum.SendDitTokens(_ethAddress)
		if err != nil {
			glog.Error(err)
			return os.Getenv("TWITTER_RESPONSE_ERROR")
		}

		err = ethereum.KYCPassed(_ethAddress, false)
		if err != nil {
			glog.Error(err)
			return os.Getenv("TWITTER_RESPONSE_ERROR")
		}

		userObject.PassedKYCDemo = true
	}

	if !userObject.PassedKYCLive && _liveKYC {
		err = ethereum.KYCPassed(_ethAddress, true)
		if err != nil {
			glog.Error(err)
			return os.Getenv("TWITTER_RESPONSE_ERROR")
		}

		userObject.PassedKYCLive = true
	}

	if _ethAddress != userObject.ETHAddress {
		userObject.ETHAddress = _ethAddress
	}

	err = database.UpdateUser(*userObject)
	if err != nil {
		glog.Error(err)
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

func alertAdmin(_text string) {
	err := sendDM(os.Getenv("TWITTER_ADMIN_USER_NAME"), os.Getenv("TWITTER_ADMIN_USER_ID"), _text)
	if err != nil {
		glog.Error(err)
	}
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

func doKYC(_user *twitter.User) (bool, bool) {
	if _user.Verified {
		return true, true
	}

	followerCountLow, _ := strconv.Atoi(os.Getenv("TWITTER_FOLLOWER_THRESHOLD_LOW"))
	statusesCountLow, _ := strconv.Atoi(os.Getenv("TWITTER_TWEET_THRESHOLD_LOW"))

	if _user.FollowersCount < followerCountLow || _user.StatusesCount < statusesCountLow {
		return false, false
	}

	followerCountHigh, _ := strconv.Atoi(os.Getenv("TWITTER_FOLLOWER_THRESHOLD_HIGH"))
	statusesCountHigh, _ := strconv.Atoi(os.Getenv("TWITTER_TWEET_THRESHOLD_HIGH"))

	if _user.FollowersCount < followerCountHigh || _user.StatusesCount < statusesCountHigh {
		return true, false
	}

	return true, true
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

func skipKYC(_username string) string {
	if strings.HasPrefix(_username, "@") {
		_username = _username[1:]
	}

	twitterUser, err := getUser(_username)
	if err != nil {
		return os.Getenv("TWITTER_ADMIN_ERROR")
	}

	userObject, err := database.GetUser(twitterUser.IDStr)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return os.Getenv("TWITTER_ADMIN_ERROR")
	}

	if userObject == nil {
		userObject = &database.User{
			TwitterID:         twitterUser.IDStr,
			TwitterScreenName: twitterUser.ScreenName,
			DateOfContact:     time.Now(),
			SkipKYC:           true,
		}
		err := database.CreateUser(*userObject)
		if err != nil {
			glog.Error(err)
			return os.Getenv("TWITTER_ADMIN_ERROR")
		}
	} else {
		userObject.SkipKYC = true
		err = database.UpdateUser(*userObject)
		if err != nil {
			glog.Error(err)
		}
	}

	return "Successfully KYCed user " + _username
}

func handleCommand(_user string, _userID string, _text string) (bool, error) {
	if strings.HasPrefix(_text, "!") {
		var err error
		switch {
		case strings.HasPrefix(_text, "!command"), strings.HasPrefix(_text, "!help"):
			if _userID == os.Getenv("TWITTER_ADMIN_USER_ID") {
				err = sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_COMMAND_LIST_ADMIN"))
			} else {
				err = sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_COMMAND_LIST"))
			}
		case strings.HasPrefix(_text, "!kyc"):
			if _userID == os.Getenv("TWITTER_ADMIN_USER_ID") {
				answer := skipKYC(_text[5:])
				err = sendDM(_user, _userID, answer)
			}
		case strings.HasPrefix(_text, "!problem"):
			err = sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_COMMAND_PROBLEM"))
			alertAdmin(_user + os.Getenv("TWITTER_ADMIN_NOTIFY_PROBLEM"))
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
