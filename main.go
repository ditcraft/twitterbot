package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/marvinkruse/dit-twitterbot/database"
	"github.com/marvinkruse/dit-twitterbot/twitter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("TWITTER_CONSUMER_KEY") == "" || os.Getenv("TWITTER_CONSUMER_SECRET") == "" || os.Getenv("TWITTER_ACCESS_TOKEN") == "" || os.Getenv("TWITTER_ACCESS_SECRET") == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	// Start API server for the Twitter Webhook
	go twitter.StartServer()

	// Wait 2 seconds in order to get the server up and running
	time.Sleep(2 * time.Second)

	// Get OAuth Token to interact with Twitter API
	_, err = database.FindOAuthTokenByHandle(os.Getenv("TWITTER_HANDLE"))
	if err != nil {
		err = database.CreateOAuthToken()
		if err != nil {
			log.Fatal("Creation of OAuth Token failed: " + err.Error())
		}

	}

	// Create Twitter Webhook Subscription
	err = twitter.InitializeWebhook()
	if err != nil {
		fmt.Println(err)
	}

	select {}
}
