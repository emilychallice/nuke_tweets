package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func checkDate(date *Date, promptMessage string) error {
	fmt.Println(promptMessage)

	var dateEntered string
	fmt.Scanln(&dateEntered)

	// User may enter nothing to skip selection.
	if strings.TrimSpace(dateEntered) == "" {
		fmt.Println("No date entered.")
		return nil
	}

	// Validate YYYY-MM-DD format
	date_arr := strings.Split(dateEntered, "-")
	if len(date_arr) != 3 {
		return errors.New("invalid format")
	}

	// Validate that YYYY, MM, DD are integers
	year, err := strconv.Atoi(date_arr[0])
	if err != nil {
		return err
	}
	month, err := strconv.Atoi(date_arr[1])
	if err != nil {
		return err
	}
	day, err := strconv.Atoi(date_arr[2])
	if err != nil {
		return err
	}

	// Validate date selection
	if month < 1 || month > 12 || day < 0 || day > 31 || year < 1000 {
		return errors.New("invalid date")
	}

	date.year = year
	date.month = month
	date.day = day

	return nil
}

func dateToString(date Date) string {
	var month string
	for k, v := range months_map {
		if v == date.month {
			month = k
		}
	}
	return strconv.Itoa(date.day) + " " + month + " " + strconv.Itoa(date.year)
}

func dateIsAfter(date Date, dateSince Date) bool {
	// Blank date means user skipped this select, so return true
	if dateSince.year == 0 {
		return true
	}
	return date.year > dateSince.year || (date.year == dateSince.year &&
		(date.month > dateSince.month || (date.month > dateSince.month && date.day > dateSince.day)))
}

func dateIsBefore(date Date, dateUntil Date) bool {
	// Blank date means user skipped this select, so return true
	if dateUntil.year == 0 {
		return true
	}
	return date.year < dateUntil.year || (date.year == dateUntil.year &&
		(date.month < dateUntil.month || (date.month < dateUntil.month && date.day < dateUntil.day)))
}

func tweetIsExempted(id string, exemptions []string, deletions []string) bool {
	for i := 0; i < len(exemptions); i++ {
		if id == strings.TrimSpace(exemptions[i]) {
			return true
		}
	}
	for i := 0; i < len(deletions); i++ {
		if id == strings.TrimSpace(deletions[i]) {
			return true
		}
	}
	return false
}
