package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getTweets() []map[string]map[string]interface{} {
	data, err := os.ReadFile("./tweets.js")
	check(err)

	// Reformat the JS data a bit
	data = bytes.Replace(data, []byte("window.YTD.tweets.part0 = "), []byte(""), 1)

	var tweets []map[string]map[string]interface{}
	err = json.Unmarshal(data, &tweets)
	if err != nil {
		fmt.Println("Error parsing tweets.js. Make sure to complete the setup instructions and overwrite with your own file.")
		os.Exit(1)
	}

	return tweets
}

func writeToDeletionLog(id string) {
	file, err := os.OpenFile("deletion_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	file.Write([]byte("\n"))
	file.Write([]byte(id))
	// defer file.Close()
}

func parseSetup() (string, string, string, []string, []string) {
	data_setup, err := os.ReadFile("SETUP.txt")
	check(err)

	data_deletelog, err := os.ReadFile("deletion_log.txt")
	check(err)

	deletionsSplit := strings.Split(string(data_deletelog), "\n")

	splitData := strings.Split(strings.Split(string(data_setup), "[HEADERS]")[1], "[EXEMPTIONS]")
	headers := strings.TrimSpace(splitData[0])
	exempt_ids := strings.TrimSpace(splitData[1])

	headersSplit := strings.Split(headers, "\n")
	xcsrf_token := strings.TrimSpace(strings.Split(headersSplit[0], ":")[1])
	auth_bearer_token := strings.TrimSpace(strings.Split(headersSplit[1], ":")[1])
	cookie_auth_token := strings.TrimSpace(strings.Split(headersSplit[2], ":")[1])

	exemptionsSplit := strings.Split(exempt_ids, "\n")

	return xcsrf_token, auth_bearer_token, cookie_auth_token, exemptionsSplit, deletionsSplit
}
