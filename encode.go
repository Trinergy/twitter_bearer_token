package main

import (
	b64 "encoding/base64"
	"fmt"
	"net/url"
)

// encodedKeySecret encodes the consumer key and secret according for Twitter API OAuth
func encodedKeySecret(consumerKey string, consumerSecret string) string {
	return b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
		url.QueryEscape(consumerKey),
		url.QueryEscape(consumerSecret))))
}
