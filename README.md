# nuke_tweets
Tool for mass-deleting tweets.

## Disclaimer
Use at your own risk! Only tested on my own account in batches of around 200 tweets. More testing & expansion of the guide and functionality to come.

## Requirements
Latest version of Python3

If it can't find a module, install with pip via command-line, e.g.:
```
pip install requests
```

## Setup
You will need to provide your own ```tweet-headers.js``` from downloading your Twitter archive. Copy and paste it in (overwrite the blank placeholder file).

You will also need to set your own header tokens in ```headers.txt``` for the user session to work. These can be found in Chrome Devtools under Application > Cookies; the Authorization Bearer token can be found in any request on the Network tab. Copy and paste your values in. You only need the 4 values given (really only 3 - the ct0 and Xcsrf tokens will be the same value).

If you would like to preserve any specific tweets that fall within the cutoff date, enter these each on their own line in ```exempt_tweets.txt```.

## Usage
Run ```nuke_tweets.py``` after completing setup - it will prompt for a cutoff date, confirm, and then send requests for each tweet ID before the cutoff range.
