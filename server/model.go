package main

import "time"

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

type session struct {
	user
	//un           string
	LastActivity time.Time
	ViewingUser   string
	FollowingUser bool
	Following     []string
}

type pageVariables struct {
	Date         string
	Time         string
	UserName	 string
}
type followVariables struct {	
	UserName	 string
	UserNames	 []string  
}

//type tweet struct {
//	Msg      string
//	Time     time.Time
//	UserName string
//	Id string
//}

//type F struct {
//	Following string
//	Follower  string
//}

type myTalk struct {
	UserName	 string
	Talk         string
	Date         string
}

var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session
var dbSessionsCleaned time.Time
//var dbTweets = map[string]tweet{}
var dbMyTalk = map[int]myTalk{}
const sessionLength  = 300