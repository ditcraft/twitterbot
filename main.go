package main

import (
	"flag"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/joho/godotenv"
	"github.com/marvinkruse/dit-twitterbot/database"
	"github.com/marvinkruse/dit-twitterbot/twitter"
)

func main() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "./log")
	flag.Set("v", "0")
	flag.Parse()
	glog.Info("Starting ditCraft twitter bot...")

	err := godotenv.Load()
	if err != nil {
		glog.Fatal("Error loading .env file")
	}

	if os.Getenv("TWITTER_CONSUMER_KEY") == "" || os.Getenv("TWITTER_CONSUMER_SECRET") == "" || os.Getenv("TWITTER_ACCESS_TOKEN") == "" || os.Getenv("TWITTER_ACCESS_SECRET") == "" {
		glog.Fatal("Consumer key/secret and Access token/secret required")
	}

	// Start API server for the Twitter Webhook
	go twitter.StartServer()

	// Start buffered channel to process twitter api calls
	go twitter.MonitorRatelimit()

	// Wait 2 seconds in order to get the server up and running
	time.Sleep(2 * time.Second)

	// Get OAuth Token to interact with Twitter API
	_, err = database.FindOAuthTokenByHandle(os.Getenv("TWITTER_HANDLE"))
	if err != nil {
		err = database.CreateOAuthToken()
		if err != nil {
			glog.Fatal("Creation of OAuth Token failed: " + err.Error())
		}

	}

	// Create Twitter Webhook Subscription
	err = twitter.InitializeWebhook()
	if err != nil {
		glog.Error(err)
	}

	select {}
}
