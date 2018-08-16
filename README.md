# Twitter Bearer Token
A tool for application-only authentication that helps you fetch your bearer token.

Consumer api keys are encoded according to the [Twitter API Documentation](https://developer.twitter.com/en/docs/basics/authentication/overview/application-only).

## Quick Start
```
// Setup your Consumer API Keys
export TWITTER_CONSUMER_KEY=<your consumer key>
export TWITTER_CONSUMER_SECRET=<your consumer secret>

// Execute App
make run

/*
* =======
*
*
* This is your BearerToken: <your bearer token>
*
*
* ======
*/

```

## Acknowledgments
This was result of building the twitter module for Chris Cummer's [wtf](https://github.com/senorprogrammer/wtf) personal terminal-based dashboard

