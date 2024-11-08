#!/usr/bin/env python3
import requests
import sys
from nuke_tweets.get_tweet_ids import get_tweet_ids_before

# Deletion URL grabbed from a normal user delete request
# (So are the json_vars)
url = "https://x.com/i/api/graphql/VaenaVgh5q5ih7kvyVjgtg/DeleteTweet"

# Read and set headers
headers_dict = {}
with open("request_headers.txt", 'r') as f:
    for line in f.readlines():
        header = line.strip().split(": ")
        headers_dict[header[0]] = header[1]

s = requests.Session()
for header in headers_dict.keys():
    s.headers.update({header: headers_dict[header]})


# Get cutoff date
def validate_date(input_string):
    try:
        date = input_string.split('-')
        day, month, year = int(date[0]), int(date[1]), int(date[2])
        if 1 <= day <= 31 and 1 <= month <= 12:
            return day, month, year
        else:
            return False
    except:
        return False
    
cutoff_date = validate_date(input("Enter cutoff date in the format DAY-MONTH-YEAR: "))
while not cutoff_date:
    print("Bad date format, try again!")
    cutoff_date = validate_date(input("Enter cutoff date in the format DAY-MONTH-YEAR: "))


# Get tweet-IDs before cutoff date and confirm...
tweet_ids = get_tweet_ids_before(cutoff_date[0], cutoff_date[1], cutoff_date[2])

while 1:
    uin = input("Continue to deletion? [y/n] ")
    if uin.strip().lower() == 'y':
        print("Deleting tweets.")
        break
    elif uin.strip().lower() == 'n':
        print("Exiting.")
        sys.exit(0)


# Delete all tweets before cutoff
deleted_tweet_ids = []
N = len(tweet_ids)
count = 1

for tweet_id in tweet_ids:
    json_vars = {"variables" : {"tweet_id":tweet_id, "dark_request":False}, "queryId":"VaenaVgh5q5ih7kvyVjgtg"}

    res = s.post(url, json=json_vars)

    print("Deleting tweet with ID {0}... {1}/{2}".format(tweet_id, count, N))
    count += 1

    if (res.status_code == 200):
        print("Deleted successfully.")
        deleted_tweet_ids.append(tweet_id)
    else:
        print("Error: Server responded with status code " + res.status_code)

print("\nDeleted " + str(len(deleted_tweet_ids)) + " tweets successfully.")
print("Adding IDs of deleted tweets to exempt_tweets file for future...")

with open("exempt_tweets.txt", 'a') as f:
    for tweet_id in deleted_tweet_ids:
        f.write("\n")
        f.write(tweet_id)

print("Done.")
