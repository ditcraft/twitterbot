package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	// Flags for authentication to the Twitter API
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	followerThreshold := flags.Int("follower-threshold", 10, "Twitter Follower Threshold")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	go listenForRequestsInMentions(client, 5*60, *followerThreshold)

	select {}
}
func listenForRequestsInMentions(_client *twitter.Client, _intervalInSeconds int64, _followerThreshold int) error {
	lastTweetID := int64(0)
	fmt.Printf("Listening for requests in mentions with %d-second interval and %d follower threshold\n", _intervalInSeconds, _followerThreshold)

	for {
		tweets, _, err := _client.Timelines.MentionTimeline(&twitter.MentionTimelineParams{
			Count:   30,
			SinceID: lastTweetID,
		})
		if err != nil {
			// TODO handle error case
			fmt.Println(err)
		}
		for i := range tweets {
			if strings.Contains(tweets[i].Text, "0x") {
				ethAddress := tweets[i].Text[strings.Index(tweets[i].Text, "0x"):40]
				// TODO validate eth address
				if isFollower(tweets[i].User) == nil && hasEnoughFollowers(tweets[i].User, _followerThreshold) == nil {
					// TODO add to DB
					// TODO fund address
					_ = ethAddress
				}
			}
			if i == len(tweets)-1 {
				lastTweetID = tweets[i].ID
			}
		}
		time.Sleep(time.Duration(_intervalInSeconds) * time.Second)
	}
}

func isFollower(_user *twitter.User) error {
	if !_user.Following {
		return errors.New("User is not a follower")
	}
	return nil
}

func hasEnoughFollowers(_user *twitter.User, _amountOfFollowers int) error {
	if _user.FollowersCount < _amountOfFollowers {
		return errors.New("User doesn't have enough followers")
	}
	return nil
}
