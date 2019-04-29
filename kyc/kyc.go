package kyc

import (
	"os"
	"strconv"
	"time"

	"github.com/stevenleeg/go-twitter/twitter"
)

// hasEnoughFollowers indicates whether the user has enough followers
// according to our threshold
func hasEnoughFollowers(_user *twitter.User, _amountOfFollowers int) bool {
	if _user.FollowersCount < _amountOfFollowers {
		return false
	}
	return true
}

// Check will perform a KYC
func Check(_user *twitter.User) (bool, bool) {
	basicKYCpassed := false
	fullKYCpassed := false

	// If the user is verified on twitter (blue checkmark) he is
	// automatically fully verified
	if _user.Verified {
		return true, true
	}

	followerCountLow, _ := strconv.Atoi(os.Getenv("TWITTER_FOLLOWER_THRESHOLD_LOW"))
	statusesCountLow, _ := strconv.Atoi(os.Getenv("TWITTER_TWEET_THRESHOLD_LOW"))

	// If the user has less than the low-kyc amounts he'll not get through
	if _user.FollowersCount < followerCountLow || _user.StatusesCount < statusesCountLow {
		return false, false
	}

	followerCountHigh, _ := strconv.Atoi(os.Getenv("TWITTER_FOLLOWER_THRESHOLD_HIGH"))
	statusesCountHigh, _ := strconv.Atoi(os.Getenv("TWITTER_TWEET_THRESHOLD_HIGH"))

	// If the user has more then the low-kyc amounts but less than the high-kyc amounts
	// he'll be qualified to use the demo mode only
	if _user.FollowersCount < followerCountHigh || _user.StatusesCount < statusesCountHigh {
		basicKYCpassed = true
	}

	// If the user has more than the high-kyc amounts, he'll be qualified to use
	// the live mode as well
	if _user.FollowersCount >= followerCountHigh && _user.StatusesCount >= statusesCountHigh {
		fullKYCpassed = true
	}

	// Correction of the outcome based on the account age, to prevent new fake accounts
	twitterDateLayout := "Mon Jan 02 15:04:05 -0700 2006"
	createDate, _ := time.Parse(twitterDateLayout, _user.CreatedAt)
	lowDays, _ := strconv.Atoi(os.Getenv("TWITTER_ACCOUNT_AGE_LOW"))
	highDays, _ := strconv.Atoi(os.Getenv("TWITTER_ACCOUNT_AGE_HIGH"))
	lowTimeAgo := time.Now().AddDate(0, 0, -lowDays)
	highTimeAgo := time.Now().AddDate(0, 0, -highDays)

	if !createDate.Before(lowTimeAgo) {
		basicKYCpassed = false
		fullKYCpassed = false
	}
	if !createDate.Before(highTimeAgo) {
		fullKYCpassed = false
	}

	return basicKYCpassed, fullKYCpassed
}
