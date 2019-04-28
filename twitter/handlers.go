package twitter

import (
	"os"
	"strings"
	"time"

	"github.com/EagleChen/mapmutex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/glog"
	"github.com/marvinkruse/dit-twitterbot/database"
	"github.com/marvinkruse/dit-twitterbot/ethereum"
	"github.com/marvinkruse/dit-twitterbot/kyc"
)

// PerUserMutex will block on a userid basis to prevent spamming from one user
var PerUserMutex *mapmutex.Mutex

func handleNewTweet(_tweetID string, _user string, _userID string, _followerCount int, _text string) {
	gotLock := false
	for !gotLock {
		gotLock = PerUserMutex.TryLock(_userID)
		defer PerUserMutex.Unlock(_userID)
	}

	if strings.Contains(_text, "verif") || strings.Contains(_text, "KYC") {
		ethAddress, containsAddress := containsETHAddress(_text)
		if containsAddress {
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

			user, err := database.GetUser(_userID)
			if err != nil && !strings.Contains(err.Error(), "not found") {
				glog.Error(err)
			}

			passedFullKYC := false

			if user != nil && user.SkipKYC == true {
				passedFullKYC = true
			} else if user == nil || !user.PassedKYCDemo || !user.PassedKYCLive {
				var passedKYC bool
				passedKYC, passedFullKYC = kyc.Check(twitterUser)
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

	if !wasCommand && strings.Contains(_text, "verif") || strings.Contains(_text, "KYC") {
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
				passedKYC, passedFullKYC = kyc.Check(twitterUser)
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

func alertAdmin(_text string) {
	err := sendDM(os.Getenv("TWITTER_ALERT_ADMIN_USER_NAME"), os.Getenv("TWITTER_ALERT_ADMIN_USER_ID"), _text)
	if err != nil {
		glog.Error(err)
	}
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

func resetKYC(_username string) string {
	if strings.HasPrefix(_username, "@") {
		_username = _username[1:]
	}

	twitterUser, err := getUser(_username)
	if err != nil {
		glog.Error(err)
		return os.Getenv("TWITTER_ADMIN_ERROR")
	}

	userObject, err := database.GetUser(twitterUser.IDStr)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		glog.Error(err)
		return os.Getenv("TWITTER_ADMIN_ERROR")
	}

	if userObject == nil {
		return "Couldn't find user " + _username
	}

	userObject.PassedKYCDemo = false
	userObject.PassedKYCLive = false

	err = database.UpdateUser(*userObject)
	if err != nil {
		glog.Error(err)
		return os.Getenv("TWITTER_ADMIN_ERROR")
	}

	return "Successfully resetted KYC for user " + _username
}

func gaveFeedback(_username string) string {
	if strings.HasPrefix(_username, "@") {
		_username = _username[1:]
	}

	twitterUser, err := getUser(_username)
	if err != nil {
		glog.Error(err)
		return os.Getenv("TWITTER_ADMIN_ERROR")
	}

	userObject, err := database.GetUser(twitterUser.IDStr)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		glog.Error(err)
		return os.Getenv("TWITTER_ADMIN_ERROR")
	}

	if userObject == nil {
		return "Couldn't find user " + _username
	}

	userObject.HasProvidedFeedback = true

	err = database.UpdateUser(*userObject)
	if err != nil {
		glog.Error(err)
		return os.Getenv("TWITTER_ADMIN_ERROR")
	}

	return "Successfully saved that user " + _username + " has provided feedback"
}

func retryKYC(_username string) string {
	if strings.HasPrefix(_username, "@") {
		_username = _username[1:]
	}

	twitterUser, err := getUser(_username)
	if err != nil {
		glog.Error(err)
		return os.Getenv("TWITTER_ADMIN_ERROR")
	}

	userObject, err := database.GetUser(twitterUser.IDStr)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		glog.Error(err)
		return os.Getenv("TWITTER_ADMIN_ERROR")
	}

	if userObject == nil {
		return "Couldn't find user " + _username
	}

	passedKYC, passedFullKYC := kyc.Check(twitterUser)
	if !passedKYC {
		return "User hasn't passed the KYC, you can pre-approve him with '!kyc " + twitterUser.ScreenName + "' and then try again"
	}

	answer := handleKYCApprove(userObject.TwitterID, userObject.TwitterScreenName, userObject.ETHAddress, false, passedFullKYC)

	if answer != os.Getenv("TWITTER_RESPONSE_SUCCESS") {
		return "There was an error while executing the KYC for user " + _username
	}

	return "Successfully resetted KYC for user " + _username
}

func handleCommand(_user string, _userID string, _text string) (bool, error) {
	if strings.HasPrefix(_text, "!") {
		var err error
		switch {
		case strings.HasPrefix(_text, "!command"), strings.HasPrefix(_text, "!help"):
			if strings.Contains(os.Getenv("TWITTER_ADMIN_USER_IDS"), _userID) {
				err = sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_COMMAND_LIST_ADMIN"))
			} else {
				err = sendDM(_user, _userID, os.Getenv("TWITTER_RESPONSE_COMMAND_LIST"))
			}
		case strings.HasPrefix(_text, "!kyc"):
			if strings.Contains(os.Getenv("TWITTER_ADMIN_USER_IDS"), _userID) {
				answer := skipKYC(_text[5:])
				err = sendDM(_user, _userID, answer)
			}
		case strings.HasPrefix(_text, "!reset"):
			if strings.Contains(os.Getenv("TWITTER_ADMIN_USER_IDS"), _userID) {
				answer := resetKYC(_text[7:])
				err = sendDM(_user, _userID, answer)
			}
		case strings.HasPrefix(_text, "!feedback"):
			if strings.Contains(os.Getenv("TWITTER_ADMIN_USER_IDS"), _userID) {
				answer := gaveFeedback(_text[10:])
				err = sendDM(_user, _userID, answer)
			}
		case strings.HasPrefix(_text, "!retry"):
			if strings.Contains(os.Getenv("TWITTER_ADMIN_USER_IDS"), _userID) {
				answer := retryKYC(_text[7:])
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
