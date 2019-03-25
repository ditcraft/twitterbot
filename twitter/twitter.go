package twitter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/marvinkruse/dit-twitterbot/database"

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

	// if !strings.Contains(tweet.Text, "@"+os.Getenv("TWITTER_HANDLE")) || !isFollower(tweet.User.ScreenName) || !hasEnoughFollowers(tweet.User, followerThreshold) {
	// 	return
	// }
	// if strings.Contains(tweet.Text, "0x") {
	// 	ethAddress := tweet.Text[strings.Index(tweet.Text, "0x"):40]
	// 	if !common.IsHexAddress(ethAddress) {
	// 		return
	// 	}
	// 	// TODO add to DB
	// 	// TODO fund address
	// 	_ = ethAddress
	// }
}

func handleNewDM(_user string, _userID string, _followerCount int, _text string) {
	isFollower, err := isFollower(_user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("[DM] %s (Follower: %t, FollowerCount: %d): %s\n", _user, isFollower, _followerCount, _text)

	followerCount, _ := strconv.Atoi(os.Getenv("TWITTER_FOLLOWER_THRESHOLD"))
	if _followerCount < followerCount {
		// TODO respond
		return
	}
	if !isFollower {
		// TODO respond
		return
	}

	ethAddress, containsAddress := containsETHAddress(_text)
	if containsAddress {
		userObject, err := database.GetUser(_userID)
		if err != nil {
			fmt.Println(err)
			return
		}
		if userObject != nil && userObject.WasFunded {
			// TODO respond
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
				// TODO error handling
				return
			}
		}

		if !userObject.WasFunded {
			// TODO fund user
			// TODO set WasFunded to true
			// TODO respond
		}
	}
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
