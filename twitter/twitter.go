package twitter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/marvinkruse/dit-twitterbot/database"
	"github.com/marvinkruse/dit-twitterbot/ethereum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stevenleeg/go-twitter/twitter"
)

var followerThreshold int

func handleNewTweet(_tweetID string, _user string, _userID string, _followerCount int, _text string) {
	isFollower, err := isFollower(_user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("[Tweet] %s (Follower: %t, FollowerCount: %d): %s\n", _user, isFollower, _followerCount, _text)

	ethAddress, containsAddress := containsETHAddress(_text)

	if containsAddress {
		followerCount, _ := strconv.Atoi(os.Getenv("TWITTER_FOLLOWER_THRESHOLD"))
		if _followerCount < followerCount {
			err := sendTweet(_tweetID, _user, os.Getenv("TWITTER_RESPONSE_FOLLOWER_COUNT"))
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		if !isFollower {
			err := sendTweet(_tweetID, _user, os.Getenv("TWITTER_RESPONSE_IS_NO_FOLLOWER"))
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		answer := handleETHRequest(_userID, _user, ethAddress)
		err := sendTweet(_tweetID, _user, answer)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func handleNewDM(_user string, _userID string, _followerCount int, _text string) {
	isFollower, err := isFollower(_user)
	if err != nil {
		err := sendDM(_userID, os.Getenv("TWITTER_RESPONSE_ERROR"))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	fmt.Printf("[DM] %s (Follower: %t, FollowerCount: %d): %s\n", _user, isFollower, _followerCount, _text)

	followerCount, _ := strconv.Atoi(os.Getenv("TWITTER_FOLLOWER_THRESHOLD"))
	if _followerCount < followerCount {
		err := sendDM(_userID, os.Getenv("TWITTER_RESPONSE_FOLLOWER_COUNT"))
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if !isFollower {
		err := sendDM(_userID, os.Getenv("TWITTER_RESPONSE_IS_NO_FOLLOWER"))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	ethAddress, containsAddress := containsETHAddress(_text)

	if containsAddress {
		answer := handleETHRequest(_userID, _user, ethAddress)
		err := sendDM(_userID, answer)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func handleETHRequest(_userID string, _userName string, _ethAddress string) string {
	userObject, err := database.GetUser(_userID)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return os.Getenv("TWITTER_RESPONSE_ERROR")
	}
	if userObject != nil && userObject.WasFunded {
		return os.Getenv("TWITTER_RESPONSE_ALREADY_FUNDED")
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
			return os.Getenv("TWITTER_RESPONSE_ERROR")
			// TODO maybe worked nevertheless?
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

	_, _, err = client.Statuses.Update("@"+_username+" "+_text, &twitter.StatusUpdateParams{
		InReplyToStatusID: int64(tweetID),
	})

	if err != nil {
		return err
	}

	return nil
}

func sendDM(_userID string, _text string) error {
	client, _, err := GetClient()
	if err != nil {
		return err
	}

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

	fmt.Println("[DM] Responded with: " + _text)

	return nil
}

// isFollower indicates whether the user follows our account
func isFollower(_screenName string) (bool, error) {
	_client, _, err := GetClient()
	if err != nil {
		return false, err
	}

	user, _, err := _client.Users.Show(&twitter.UserShowParams{
		ScreenName: _screenName,
	})
	if err != nil {
		return false, err
	}

	if !user.Following {
		return false, nil
	}
	return true, nil
}

// hasEnoughFollowers indicates whether the user has enough followers
// according to our threshold
func hasEnoughFollowers(_user *twitter.User, _amountOfFollowers int) bool {
	if _user.FollowersCount < _amountOfFollowers {
		return false
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
