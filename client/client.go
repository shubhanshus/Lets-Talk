package main

import (
	"net/http"	
)


func main() {

	//setup path
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	//http.HandleFunc("/home", home)
	http.HandleFunc("/cancel", cancel)
	http.HandleFunc("/cancelaccount", cancelaccount)

	http.HandleFunc("/talk", postTalk)
	http.HandleFunc("/list", showTalk)
	http.HandleFunc("/follow", follow)
	http.HandleFunc("/unfollow", unfollow)
	http.HandleFunc("/followothers", followothers)
	http.HandleFunc("/unfollowothers", unfollowothers)

	//resource path
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}
