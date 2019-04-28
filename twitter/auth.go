package twitter

import (
	"os"
	"time"

	"github.com/dghubble/oauth1"
	twitterOAuth "github.com/dghubble/oauth1/twitter"
	"github.com/golang/glog"
	"github.com/marvinkruse/dit-twitterbot/database"
	"github.com/stevenleeg/go-twitter/twitter"
)

// OAuthRequest collects all required information for completing the
// OAuth1 authentication flow
type OAuthRequest struct {
	Handle       string
	PIN          string
	RequestToken string
}

// Stores all threads waiting to be unlocked
var requestQueue chan chan bool

// MonitorRatelimit will ensure that all twitter calls are executed at most
// once per two seconds
func MonitorRatelimit() {
	requestQueue = make(chan chan bool, 500)

	for {
		time.Sleep(2 * time.Second)
		request := <-requestQueue
		close(request)
	}
}

// awaitRatelimit will add the current thread's execution to a queue and will
// block until it is released by the ratelimiting thread
func awaitRatelimit() {
	await := make(chan bool)
	requestQueue <- await
	<-await
}

// GetClient returns a twitter client
func GetClient() (*twitter.Client, *database.OAuthToken, error) {
	oauthToken, err := database.FindOAuthTokenByHandle(os.Getenv("TWITTER_HANDLE"))
	if err != nil {
		glog.Errorf("Could not find OAuth token for %s", os.Getenv("TWITTER_HANDLE"))
		return nil, nil, err
	}

	conf := getOAuthConfiguration()
	token := oauth1.NewToken(oauthToken.OAuthToken, oauthToken.OAuthTokenSecret)
	httpClient := conf.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient), oauthToken, nil
}

func getOAuthConfiguration() *oauth1.Config {
	return &oauth1.Config{
		ConsumerKey:    os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		Endpoint:       twitterOAuth.AuthorizeEndpoint,
		CallbackURL:    "oob",
	}
}
