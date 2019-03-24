package twitter

import (
	"fmt"
	"os"
	"strconv"

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

func handleNewDM(_user string, _followerCount int, _text string) {
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
	// TODO check in DB
	// TODO store in DB
	// TODO check ETH stuff
	// TODO respond
	return
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
