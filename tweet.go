package main


import (
"net/http"
)

func putTweet(req *http.Request, user *user, tweet *tweet) {
	dbTweets[user.UserName] = *tweet
}

func getTweets(req *http.Request, user *user) ([]tweet, error) {
	//ctx := appengine.NewContext(req)
	//
	var tweets []tweet
	//q := datastore.NewQuery("Tweets")
	//
	//if user != nil {
	//	// show tweets of a specific user
	//	userKey := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	//	q = q.Ancestor(userKey)
	//}
	//
	//q = q.Order("-Time").Limit(20)
	//_, err := q.GetAll(ctx, &tweets)
	return tweets, nil
}
