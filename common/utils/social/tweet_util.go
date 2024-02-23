package social

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Mu-Exchange/Mu-End/common/errors"
)

const CALL_LIMIT = 5

type TweetRateLimiter struct {
	counters map[string]int
	mu       sync.Mutex
}

func NewRateLimiter() *TweetRateLimiter {
	return &TweetRateLimiter{
		counters: make(map[string]int),
	}
}

func (rl *TweetRateLimiter) CanCallMethod(userID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now().Format("2006-01-02")

	if _, ok := rl.counters[userID]; !ok {
		rl.counters[userID] = 0
	}

	if rl.counters[userID] < CALL_LIMIT {
		rl.counters[userID]++
		return true
	}

	if rl.counters[userID] == CALL_LIMIT && now != rl.getResetTime(userID) {
		rl.counters[userID] = 1
		return true
	}

	return false
}

func (rl *TweetRateLimiter) getResetTime(userID string) string {
	return time.Unix(0, int64(rl.counters[userID])).Format("2006-01-02")
}

func (rl *TweetRateLimiter) FindTweet(twitterId, userID, tweetId, token string) (*Tweet, error) {
	if !rl.CanCallMethod(userID) {
		return nil, errors.ErrTweetLimit
	}
	url := "https://api.twitter.com/2/tweets/" + tweetId + "?tweet.fields=attachments%2Cauthor_id%2Ccreated_at%2Centities%2Cgeo%2Cid%2Cin_reply_to_user_id%2Clang%2Cpossibly_sensitive%2Creferenced_tweets%2Csource%2Ctext%2Cwithheld"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	tweet := &Tweet{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, tweet); err != nil {
		return nil, err
	}
	if len(tweet.Data.Id) == 0 {
		return nil, errors.ErrTweetNotFound
	}
	if len(tweet.Data.AuthorId) == 0 || !strings.EqualFold(tweet.Data.AuthorId, twitterId) {
		return nil, errors.ErrTweetNotBelongUser
	}
	return tweet, nil
}

func (rl *TweetRateLimiter) CheckTweetDailyAndKeyword(tweet *Tweet, keyword string) error {
	if len(tweet.Data.Text) == 0 || len(keyword) == 0 || !strings.Contains(strings.ToLower(tweet.Data.Text), strings.ToLower(keyword)) {
		return errors.ErrTweetNotFoundKeyword
	}

	date1 := tweet.Data.CreatedAt.Truncate(24 * time.Hour)
	date2 := time.Now().Truncate(24 * time.Hour)
	if !date1.Equal(date2) {
		return errors.ErrTweetNotBelongToday
	}
	return nil
}
