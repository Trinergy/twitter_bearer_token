package main

import (
	"encoding/json"
	"os"
)

// Client represents the data required to get authentication from the Twitter API
type Client struct {
	apiURL         string
	consumerKey    string
	consumerSecret string
}

// NewClient creates and returns a bearer token
func NewClient() *Client {
	client := Client{
		apiURL:         "https://api.twitter.com/oauth2/token",
		consumerKey:    os.Getenv("TWITTER_CONSUMER_KEY"),
		consumerSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
	}

	return &client
}

/* -------------------- Public Functions -------------------- */

// BearerToken returns a bearer token using the consumer key and consumer secret
func (client *Client) BearerToken() Token {
	token, err := client.bearertoken()
	if err != nil {
		return Token{}
	}

	return token
}

/* -------------------- Private Functions -------------------- */

// bearertoken is the private interface for retrieving the bearer token based on the consumer key and secret
func (client *Client) bearertoken() (token Token, err error) {

	data, err := Request(client.consumerKey, client.consumerSecret, client.apiURL)
	if err != nil {
		return token, err
	}
	err = json.Unmarshal(data, &token)

	return
}
