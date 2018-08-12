package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

// Request retreives the bearer token using the consumer key and secret
func Request(consumerKey string, consumerSecret string, apiURL string) ([]byte, error) {
	// Required Header and Body according to Twitter API documentation
	reqBody := bytes.NewBuffer([]byte(`grant_type=client_credentials`))
	req, err := http.NewRequest("POST", apiURL, reqBody)

	encodedKeySecret := encodedKeySecret(consumerKey, consumerSecret)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedKeySecret))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(reqBody.Len()))

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ParseBody(resp)
	if err != nil {
		return nil, err
	}

	return data, err
}

func ParseBody(resp *http.Response) ([]byte, error) {
	var buffer bytes.Buffer
	_, err := buffer.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
