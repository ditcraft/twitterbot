package twitter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ditcraft/twitterbot/database"
	"github.com/golang/glog"
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
	http.HandleFunc("/", handleTwitterWebhook)
	err := http.ListenAndServe(":"+os.Getenv("TWITTER_WEB_HOOK_PORT"), nil)
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
				if !strings.HasPrefix("RT", data.TweetCreateEvents[i].Text) {
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
	}

	if len(data.FollowEvents) > 0 {
		for _, newFollower := range data.FollowEvents {
			isFollower[newFollower.Source.ID] = true
		}
	}

	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

// InitializeWebhook will register and create the webhook through the twitter api
func InitializeWebhook() error {
	webhookID, err := database.GetKey("webhookID")
	if err != nil {
		return err
	}

	if webhookID == "" {
		id, err := createWebhook()
		if err != nil {
			return err
		}

		glog.Infof("Webhook %s created successfully", id)
		database.SetKey("webhookID", id)
	}

	if err := createSubscription(); err != nil {
		return err
	}

	fmt.Printf("Subscription created successfully")
	return nil
}

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
