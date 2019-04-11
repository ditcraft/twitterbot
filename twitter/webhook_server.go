package twitter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/golang/glog"
	"github.com/marvinkruse/dit-twitterbot/database"
	"github.com/stevenleeg/go-twitter/twitter"
)

type incomingDM struct {
	Type             string `json:"type"`
	ID               string `json:"id"`
	CreatedTimestamp string `json:"created_timestamp"`
	MessageCreated   struct {
		SenderID    string `json:"sender_id"`
		MessageData struct {
			Text string `json:"text"`
		} `json:"message_data"`
	} `json:"message_create"`
}

type incomingTweet struct {
	IDStr string    `json:"id_str"`
	Text  string    `json:"text"`
	User  tweetUser `json:"user"`
}

type incomingFollow struct {
	Source struct {
		ScreenName string `json:"screen_name"`
		ID         string `json:"id"`
	} `json:"source"`
}

type tweetUser struct {
	ID             string `json:"id_str"`
	ScreenName     string `json:"screen_name"`
	FollowersCount int64  `json:"followers_count"`
}

type dmUser struct {
	ScreenName     string `json:"screen_name"`
	FollowersCount int64  `json:"followers_count"`
}

type incomingWebhook struct {
	ForUserID           string            `json:"for_user_id"`
	DirectMessageEvents []incomingDM      `json:"direct_message_events"`
	TweetCreateEvents   []incomingTweet   `json:"tweet_create_events"`
	FollowEvents        []incomingFollow  `json:"follow_events"`
	User                map[string]dmUser `json:"users"`
}

// StartServer spins up a webserver for the API
func StartServer() {
	followerThreshold, _ = strconv.Atoi(os.Getenv("TWITTER_FOLLOWER_THRESHOLD"))

	http.HandleFunc("/webhook/twitter", handleTwitterWebhook)
	err := http.ListenAndServeTLS(":"+os.Getenv("TWITTER_WEB_HOOK_PORT"), os.Getenv("SERVER_SSL_CERT_PATH"), os.Getenv("SERVER_SLL_KEY_PATH"), nil)
	if err != nil {
		glog.Error(err)
	}
}

func handleTwitterWebhook(w http.ResponseWriter, r *http.Request) {
	// A GET request signals that Twitter is attempting a CRC request
	if r.Method == "GET" {
		keys, ok := r.URL.Query()["crc_token"]
		if !ok || len(keys) < 1 {
			w.WriteHeader(400)
			w.Write([]byte("Bad request"))
			return
		}

		mac := hmac.New(sha256.New, []byte(os.Getenv("TWITTER_CONSUMER_SECRET")))
		mac.Write([]byte(keys[0]))

		token := "sha256=" + base64.StdEncoding.EncodeToString(mac.Sum(nil))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("{\"response_token\": \"" + token + "\"}"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	data := &incomingWebhook{}
	err := decoder.Decode(data)
	if err != nil {
		glog.Error(err)
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}

	if len(data.DirectMessageEvents) > 0 {
		for i := range data.DirectMessageEvents {
			if data.DirectMessageEvents[i].MessageCreated.SenderID != os.Getenv("TWITTER_ID") {
				handleNewDM(
					data.User[data.DirectMessageEvents[i].MessageCreated.SenderID].ScreenName,
					data.DirectMessageEvents[i].MessageCreated.SenderID,
					int(data.User[data.DirectMessageEvents[i].MessageCreated.SenderID].FollowersCount),
					data.DirectMessageEvents[i].MessageCreated.MessageData.Text,
				)
			}
		}
	}

	if len(data.TweetCreateEvents) > 0 {
		for i := range data.TweetCreateEvents {
			if data.ForUserID == os.Getenv("TWITTER_ID") && data.TweetCreateEvents[i].User.ID != os.Getenv("TWITTER_ID") {
				handleNewTweet(
					data.TweetCreateEvents[i].IDStr,
					data.TweetCreateEvents[i].User.ScreenName,
					data.TweetCreateEvents[i].User.ID,
					int(data.TweetCreateEvents[i].User.FollowersCount),
					data.TweetCreateEvents[i].Text,
				)
			}
		}
	}

	// if len(data.FollowEvents) > 0 {
	// 	go server.processFollows(data.FollowEvents)
	// }

	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func InitializeWebhook() error {
	webhookID, err := database.GetKey("webhookID")
	if err != nil {
		return err
	}

	// If we don't already have a webhook ID we should create it
	if webhookID == "" {
		id, err := createWebhook()
		if err != nil {
			return err
		}

		glog.Infof("Webhook %s created successfully", id)
		database.SetKey("webhookID", id)
	}

	// And subscribe to TCRPartyVIP's DMs
	if err := createSubscription(); err != nil {
		return err
	}

	fmt.Printf("Subscription created successfully")
	return nil
}

// CreateWebhook creates a new webhook and subscribes it to the user, allowing
// us to receive notifications for new DMs. This should only be used on the
// TCRPartyVIP bot.
func createWebhook() (string, error) {
	client, _, err := GetClient()
	if err != nil {
		return "", err
	}

	webhookParams := &twitter.AccountActivityRegisterWebhookParams{
		EnvName: os.Getenv("TWITTER_ENV"),
		URL:     "https://" + os.Getenv("BASE_URL") + "/webhook/twitter",
	}
	webhook, _, err := client.AccountActivity.RegisterWebhook(webhookParams)

	if err != nil {
		return "", err
	}

	return webhook.ID, nil
}

// createSubscription subscribes the current webhook to the given user
func createSubscription() error {
	client, _, err := GetClient()
	if err != nil {
		return err
	}

	subParams := &twitter.AccountActivityCreateSubscriptionParams{
		EnvName: os.Getenv("TWITTER_ENV"),
	}
	_, err = client.AccountActivity.CreateSubscription(subParams)

	return err
}
