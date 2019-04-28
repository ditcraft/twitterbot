# ditCraft Twitter Bot
The ditCraft Twitter bot is an automated piece of software that interacts with followers. It's main purpose for now is to perform a KYC - or know-your-coder - to ensure that users of the dit client are legitimate ones instead of being fake. Knowledge tokens can only work when it's impossible (or at least very hard) to perform sybil attacks on the system that is making use of them.

In the future, the Twitter bot will receive updates including more features.

## Running this bot yourself
You like the idea and want to use a Twitter bot in your own project? Go ahead - we embrace open source! In order to work, this bot needs an .env file in the same directory, containing the following things:

```
TWITTER_CONSUMER_KEY=
TWITTER_CONSUMER_SECRET=
TWITTER_ACCESS_TOKEN=
TWITTER_ACCESS_SECRET=
TWITTER_ENV=

TWITTER_HANDLE=<TWITTER-ACCOUNT-USERNAME-HERE>
TWITTER_ID=<ID-OF-THE-TWITTER-ACCOUNT>
TWITTER_WEB_HOOK_PORT=443

TWITTER_FOLLOWER_THRESHOLD_LOW=20
TWITTER_TWEET_THRESHOLD_LOW=15

TWITTER_FOLLOWER_THRESHOLD_HIGH=100
TWITTER_TWEET_THRESHOLD_HIGH=50

TWITTER_ACCOUNT_AGE_LOW=30
TWITTER_ACCOUNT_AGE_HIGH=90

TWITTER_ADMIN_USER_ID=<TWITTER-ACCOUNT-ID-OF-ADMIN>
TWITTER_ADMIN_USER_NAME=<TWITTER-ACCOUNT-NAME-OF-ADMIN>

MONGO_DB_ADDRESS=127.0.0.1:27017

BASE_URL=<BASE-URL-OF-SERVER-WITHOUT-HTTP(S)>

SERVER_SSL_CERT_PATH=/etc/letsencrypt/live/<YOUR-TLD>/fullchain.pem
SERVER_SLL_KEY_PATH=/etc/letsencrypt/live/<YOUR-TLD>/privkey.pem

ETHEREUM_PRIVATEKEY=0000000000000000000000000000000000000000000000000000000000000000
ETHEREUM_ADDRESS=0x0000000000000000000000000000000000000000
ETHEREUM_RPC=https://my.fancy.rpc.url

CONTRACT_DIT_TOKEN=0x0000000000000000000000000000000000000000
CONTRACT_DIT_COORDINATOR_LIVE=0x0000000000000000000000000000000000000000
CONTRACT_DIT_COORDINATOR_DEMO=0x0000000000000000000000000000000000000000
DIT_TOKEN_AMOUNT=50
```