package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestVars struct {
	TweetID     string `json:"tweet_id"`
	DarkRequest bool   `json:"dark_request"`
}

type RequestParams struct {
	QueryID   string      `json:"queryId"`
	Variables RequestVars `json:"variables"`
}

func DeleteTweet(tweetID string, xcsrf_token string, auth_bearer_token string, cookie_auth_token string) {
	jsonData, err := json.Marshal(RequestParams{QueryID: "VaenaVgh5q5ih7kvyVjgtg", Variables: RequestVars{TweetID: tweetID, DarkRequest: false}})
	check(err)

	// Create and populate the request
	req, err := http.NewRequest("POST", "https://x.com/i/api/graphql/VaenaVgh5q5ih7kvyVjgtg/DeleteTweet", bytes.NewBuffer(jsonData))
	check(err)

	req.Header.Add("Authorization", auth_bearer_token)
	req.Header.Add("Cookie", "auth_token="+cookie_auth_token+"; ct0="+xcsrf_token)
	req.Header.Add("X-Csrf-Token", xcsrf_token)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	// Initiate client and perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	if res.StatusCode == 200 {
		fmt.Println("Successfully deleted tweet with ID", tweetID)
		writeToDeletionLog(tweetID)
	}
}
