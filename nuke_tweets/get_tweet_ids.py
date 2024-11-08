import json
import sys

# Helper functions
months_dict = {"Jan":1, "Feb":2, "Mar":3, "Apr":4,
               "May":5, "Jun":6, "Jul":7, "Aug":8,
               "Sep":9, "Oct":10, "Nov":11, "Dec":12}

def get_month_name(n):
     for x in months_dict.items():
          if x[1] == n:
               return x[0]


# Convert twitter-provided data to json
with open("tweet-headers.js", 'r') as f:
    data_js = f.read()
    try:
        data = json.loads( data_js.replace("window.YTD.tweet_headers.part0 =",
                                        "{ \"window.YTD.tweet_headers.part0\" :")
                                        + "}" )
    except:
        print("Missing valid tweet-headers.js file - copy from your downloaded Twitter archive.")
        print("Exiting...")
        sys.exit(0)


# Read exempt tweets (user-added and already-deleted)
with open("exempt_tweets.txt", 'r') as f:
    exempt_tweets = f.read().splitlines()


# ID-getter functions (just this one for now...)
def get_tweet_ids_before( DD, MM, YYYY ):
    tweets = []

    for obj in data["window.YTD.tweet_headers.part0"]:
        date = obj["tweet"]["created_at"].split(' ')

        month = months_dict[ date[1] ]
        day = int(date[2])
        year = int(date[-1])

        if obj["tweet"]["tweet_id"] not in exempt_tweets:
            if year < YYYY or (year == YYYY and (month < MM or (month == MM and day < DD))):
                tweets.append(obj["tweet"]["tweet_id"])

    print("Found {0} deletable tweets before {1} {2} {3}.".format(len(tweets), DD, get_month_name(MM), YYYY))
    return tweets
