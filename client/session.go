package main

import (
"net/http"
)
const sessionLength = 300
func createCookie(sID string,w http.ResponseWriter) {

	c := &http.Cookie{
		Name:  "session",
		Value: sID,
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	//dbSessions[c.Value] = session{user, time.Now(),"",false,nil}

}