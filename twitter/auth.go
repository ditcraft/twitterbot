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

// // GetClientFromToken returns a twitter client given an OAuth token
// func GetClientFromToken(oauthToken *database.OAuthToken) *twitter.Client {
// 	conf := getOAuthConfiguration()
// 	token := oauth1.NewToken(oauthToken.OAuthToken, oauthToken.OAuthTokenSecret)
// 	httpClient := conf.Client(oauth1.NoContext, token)

// 	return twitter.NewClient(httpClient)
// }

// // GetOAuthURL returns a URL that can be used to authenticate a user via the
// // OAuth 1 API
// func (request *OAuthRequest) GetOAuthURL() (string, error) {
// 	conf := getOAuthConfiguration()

// 	requestToken, _, err := conf.RequestToken()
// 	if err != nil {
// 		return "", err
// 	}

// 	authorizationURL, err := conf.AuthorizationURL(requestToken)
// 	if err != nil {
// 		return "", err
// 	}

// 	request.RequestToken = requestToken

// 	return authorizationURL.String(), nil
// }

// // ReceivePIN completes the OAuth dance by retrieving user info. Note that it
// // requires the PIN to be set on the OAuthRequest struct.
// func (request *OAuthRequest) ReceivePIN() error {
// 	conf := getOAuthConfiguration()
// 	accessToken, accessSecret, err := conf.AccessToken(
// 		request.RequestToken,
// 		"",
// 		request.PIN,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	// Save the OAuth credentials in the database
// 	token := &database.OAuthToken{
// 		TwitterHandle:    request.Handle,
// 		OAuthToken:       accessToken,
// 		OAuthTokenSecret: accessSecret,
// 	}

// 	// Fetch the user ID
// 	client := GetClientFromToken(token)
// 	verifyParams := &twitter.AccountVerifyParams{
// 		SkipStatus:   twitter.Bool(true),
// 		IncludeEmail: twitter.Bool(false),
// 	}
// 	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
// 	if err != nil {
// 		log.Println("Could fetch user data")
// 		return err
// 	}

// 	token.TwitterID = user.ID

// 	err = token.Save()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
