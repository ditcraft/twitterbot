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

func handleNewTweet(_user string, _followerCount int, _text string) {
	isFollower, err := isFollower(_user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[Tweet] %s (Follower: %t, FollowerCount: %d): %s\n", _user, isFollower, _followerCount, _text)
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
		userObject, err := database.GetUser(_userID)
		if err != nil && !strings.Contains(err.Error(), "not found") {
			err := sendDM(_userID, os.Getenv("TWITTER_RESPONSE_ERROR"))
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		if userObject != nil && userObject.WasFunded {
			err := sendDM(_userID, os.Getenv("TWITTER_RESPONSE_ALREADY_FUNDED"))
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		if userObject == nil {
			userObject = &database.User{
				TwitterID:         _userID,
				TwitterScreenName: _user,
				ETHAddress:        ethAddress,
				DateOfContact:     time.Now(),
				WasFunded:         false,
			}
			err := database.CreateUser(*userObject)
			if err != nil {
				err := sendDM(_userID, os.Getenv("TWITTER_RESPONSE_ERROR"))
				if err != nil {
					fmt.Println(err)
				}
				return
			}
		}

		if !userObject.WasFunded {
			err := ethereum.SendEther(ethAddress, 1)
			if err != nil {
				err := sendDM(_userID, os.Getenv("TWITTER_RESPONSE_ERROR"))
				if err != nil {
					fmt.Println(err)
				}
				return
			}

			userObject.WasFunded = true

			if ethAddress != userObject.ETHAddress {
				userObject.ETHAddress = ethAddress
			}

			err = database.UpdateUser(*userObject)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = sendDM(_userID, os.Getenv("TWITTER_RESPONSE_SUCCESS"))
			if err != nil {
				fmt.Println(err)
			}

		}
	}
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
