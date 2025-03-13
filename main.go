package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Date struct {
	year  int
	month int
	day   int
}

func main() {
	fmt.Println("---")
	fmt.Println("Tweet Nuke v2 - rewritten in Go.\nMake sure you read the directions on Github and fill out SETUP.txt.")
	fmt.Println("---")

	tweets := getTweets()

	xcsrf_token, auth_bearer_token, cookie_auth_token, exemptions, deletions := parseSetup()

	scanner := bufio.NewScanner(os.Stdin)

	// Get keyword filter from user
	var keyword string
	fmt.Println("Enter a keyword, or hit Enter for none:")
	if scanner.Scan() {
		line := scanner.Text()
		keyword = line
		if keyword == "" {
			fmt.Println("No keyword selected.")
		}
	}

	// Get date filters from user
	var date_since Date
	err := checkDate(&date_since, "Select for all tweets SINCE (after) this date [YYYY-MM-DD]:")
	for err != nil {
		err = checkDate(&date_since, "Invalid date! Try again.\nSelect for all tweets SINCE (after) this date [YYYY-MM-DD]:")
	}

	var date_until Date
	err = checkDate(&date_until, "Select for all tweets UNTIL (before) this date [YYYY-MM-DD]:")
	for err != nil {
		err = checkDate(&date_until, "Invalid date! Try again.\nSelect for all tweets UNTIL (before) this date [YYYY-MM-DD]:")
	}

	// Select the tweets
	var tweet_ids []string
	count := 0
	for i := 0; i < len(tweets); i++ {
		// Tweet contents and id
		tweet_text := tweets[i]["tweet"]["full_text"].(string)
		tweet_id := tweets[i]["tweet"]["id"].(string)
		tweet_likes, _ := strconv.Atoi(tweets[i]["tweet"]["favorite_count"].(string))
		tweet_rts, _ := strconv.Atoi(tweets[i]["tweet"]["retweet_count"].(string))

		// Tweet date
		var tweet_date Date
		tweet_datestring_arr := strings.Split(tweets[i]["tweet"]["created_at"].(string), " ")
		tweet_date.year, _ = strconv.Atoi(tweet_datestring_arr[5])
		tweet_date.month = months_map[tweet_datestring_arr[1]]
		tweet_date.day, _ = strconv.Atoi(tweet_datestring_arr[2])

		// Filter...
		if strings.Contains(tweet_text, keyword) &&
			dateIsAfter(tweet_date, date_since) &&
			dateIsBefore(tweet_date, date_until) &&
			!tweetIsExempted(tweet_id, exemptions, deletions) {

			fmt.Println(tweet_text)
			fmt.Println("ID:", tweet_id)
			fmt.Println("Posted on:", dateToString(tweet_date))
			fmt.Println("Likes:", tweet_likes, "| Retweets: ", tweet_rts)
			fmt.Println()

			count++
			tweet_ids = append(tweet_ids, tweet_id)
		}
	}

	// Print filter selections back to user to review
	fmt.Print("\nFound " + strconv.Itoa(count) + " tweets ")
	if date_since.year != 0 {
		fmt.Print("since ", dateToString(date_since))
	}
	if date_until.year != 0 {
		fmt.Print(", until ", dateToString(date_until))
	}
	if keyword != "" {
		fmt.Print(" with keyword ", keyword)
	}
	fmt.Print(".\n")

	// Check that user wishes to continue
	var continueChoice string
	for strings.ToLower(continueChoice) != "y" {
		fmt.Println("Continue? [y/n]")
		fmt.Scanln(&continueChoice)
		if strings.ToLower(continueChoice) == "n" {
			return
		}
	}
	fmt.Println("Continuing to deletion.")

	// Send delete requests
	for _, id := range tweet_ids {
		fmt.Println("Deleting tweet with id", id)
		DeleteTweet(id, xcsrf_token, auth_bearer_token, cookie_auth_token)

		// TODO: more specific reporting to user here if the request fails
		// TODO: also check if any found tweet was a RT of another user (full text begin with 'RT @')
		// --> probably implement feature to un-retweet instead of delete
	}
}
