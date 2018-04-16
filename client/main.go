package main

import (
	"html/template"
	"net/http"
	"time"
	"log"
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../proto"
)

var tpl *template.Template
var u user
var talks []myTalk
var address = "localhost:8080"
var userLoggedIn =false
var un string


func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}



func index(w http.ResponseWriter, req *http.Request){
	var IndexPageVars pageVariables
	var uname string
	now := time.Now() // find the time right now

	if userLoggedIn{
		uname = "test"
		log.Println("Hello user:", uname)
	}else {
		log.Println("User not logged in")
		uname = ""
	}
    IndexPageVars = pageVariables{ //store the date and time in a struct
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04PM"),
      UserName: uname,
    }
    //log.Println("uname", uname)
	tpl, err := template.ParseFiles("templates/index.html") //parse the html file
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it on terminal
	}
	err = tpl.Execute(w, IndexPageVars) //execute the template and pass it to index page
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it on terminal
	}


}

func signup(w http.ResponseWriter, req *http.Request) {
	//if alreadyLoggedIn(w, req) {
	//	http.Redirect(w, req, "/home", http.StatusSeeOther)
	//	return
	//}
	if userLoggedIn{
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("email")
		p1 := req.FormValue("password1")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		//// username taken?
		//if _, ok := dbUsers[un]; ok {
		//	http.Error(w, "Username already taken", http.StatusForbidden)
		//	return
		//}
		//// store user in dbUsers
		//bs, err := bcrypt.GenerateFromPassword([]byte(p1), bcrypt.MinCost)
		//if err != nil {
		//	http.Error(w, "Internal server error", http.StatusInternalServerError)
		//	return
		//}
		//u = user{un, bs, f, l}
		//
		//dbUsers[un] = u
		////dial server
		////address := "localhost:8080"

		//dial server
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewLetstalkClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SendSignup(ctx, &pb.SignupRequest{Email: un, Password1: p1, Firstname: f, Lastname: l})
		if err != nil {
			http.Error(w, r.Message, http.StatusForbidden)
			return
		}
		un = r.Message
		userLoggedIn=true
		log.Println(r.Message)
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	/*if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/index", http.StatusSeeOther)
		return
	}*/
	if userLoggedIn{
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}

	var u user
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("name")
		p := req.FormValue("password")
		// is there a username?

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewLetstalkClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SendLogin(ctx, &pb.LoginRequest{Email:un,Password1:p})
		if err != nil {
			http.Error(w, r.Message, http.StatusForbidden)
			return
		}
		createCookie(r.SessionId,w)
		userLoggedIn=true

		//u, ok := dbUsers[un]
		//if !ok {
		//	log.Println(un)
		//	log.Println(dbUsers)
		//	http.Error(w, "Username not found", http.StatusForbidden)
		//	return
		//}
		//// does the entered password match the stored password?
		//err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		//if err != nil {
		//	http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		//	return
		//}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.html", u)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !userLoggedIn{
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}
	cookie, _ := req.Cookie("session")
	//dial server
	conn2, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	c2 := pb.NewLetstalkClient(conn2)
	//log.Printf("connection established")
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c2.SendLogout(ctx, &pb.LogoutRequest{Email: un})
	if err != nil {
		//log.Println(r.Message,"  ",err)
		http.Error(w, r.Message, http.StatusForbidden)
		return
	}
	log.Println(u.UserName, r.Message)
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	userLoggedIn=false
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func postTalk(w http.ResponseWriter, req *http.Request) {
	
	
    http.Redirect(w, req, "/", http.StatusSeeOther)
}

func showTalk(w http.ResponseWriter, req *http.Request) {
	//get json api
    json.NewEncoder(w).Encode(talks)


}

func cancel(w http.ResponseWriter, req *http.Request) {
	
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func cancelaccount(w http.ResponseWriter, req *http.Request){
	
	tpl, err := template.ParseFiles("templates/cancel.html") //parse the html file
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it on terminal
	}
	err = tpl.Execute(w, "") //execute the template and pass it to index page
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it on terminal
	}

}

func follow(w http.ResponseWriter, req *http.Request) {

}

func followothers(w http.ResponseWriter, req *http.Request) {
	

}

func updateTweets(session session) {
	
}

func deleteTweets(){
	
}