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
	lastActivity time.Time
	Tweets        []Tweet
	ViewingUser   string
	FollowingUser bool
	Following     []F
}

type PageVariables struct {
	Date         string
	Time         string
}

type Tweet struct {
	Msg      string
	Time     time.Time
	UserName string
}

type F struct {
	Following string
	Follower  string
}


var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session
var dbSessionsCleaned time.Time

const sessionLength int = 300