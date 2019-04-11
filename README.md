# ditCraft twitter bot

## In order to work this project needs an .env file in the same directory, containing the following things:
```
TWITTER_CONSUMER_KEY=
TWITTER_CONSUMER_SECRET=
TWITTER_ACCESS_TOKEN=
TWITTER_ACCESS_SECRET=
TWITTER_ENV=
TWITTER_HANDLE=
TWITTER_ID=
TWITTER_WEB_HOOK_PORT=443
TWITTER_FOLLOWER_THRESHOLD=10
TWITTER_TWEET_THRESHOLD=10

MONGO_DB_ADDRESS=127.0.0.1:27017

BASE_URL=<YOUR-SERVER-URL-HERE>

SERVER_SSL_CERT_PATH=/etc/letsencrypt/live/<YOUR-SERVER-URL-HERE>/fullchain.pem
SERVER_SLL_KEY_PATH=/etc/letsencrypt/live/<YOUR-SERVER-URL-HERE>/privkey.pem

ETHEREUM_PRIVATEKEY=0000000000000000000000000000000000000000000000000000000000000000
ETHEREUM_ADDRESS=0x0000000000000000000000000000000000000000
ETHEREUM_RPC=https://rinkeby.infura.io/v3/00000000000000000000000000000000

TWITTER_RESPONSE_KYC_FAIL_TWEET='Unfortunately you didn't pass our automated KYC process. Kindly slide into our DMs so that we can manually verify you.'
TWITTER_RESPONSE_KYC_FAIL_DM='Unfortunately you didn't pass our automated KYC process. We'll get back to you within 24 hours to manually verify you.'
TWITTER_RESPONSE_IS_NO_FOLLOWER='In order to interact with me, you need to be a follower!'
TWITTER_RESPONSE_ALREADY_FUNDED_DM='You are already verified! You aren't and still get this message? Please respond with "!problem".'
TWITTER_RESPONSE_ALREADY_FUNDED_TWEET='You are already verified! You aren't and still get this message? Kindly slide into our DMs and we'll handle this."'
TWITTER_RESPONSE_SUCCESS='Thanks, you are now verified and ready to go - let us know what you think!'
TWITTER_RESPONSE_ERROR='Unfortunately there was a problem, we are looking into it and we'll get back to you asap!'
TWITTER_RESPONSE_COMMAND_PROBLEM='Alright, I've notified a human to look into this. We'll get back to you within 24 hours.'
TWITTER_RESPONSE_COMMAND_NOT_FOUND='Sorry, I don't recognize that command. Please try again or respond with "!problem" to alert my owners.'
TWITTER_RESPONSE_COMMAND_LIST='Available commands: !problem'
```