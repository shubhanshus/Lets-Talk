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
		uname = u.UserName
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
			errMsg:= err.Error()
			errMsg=errMsg[33:len(errMsg)]
			http.Error(w, errMsg , http.StatusForbidden)
		}

		userLoggedIn=true
		log.Println(r.Message)
		u.UserName = r.Message
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
			errMsg:= err.Error()
			errMsg=errMsg[33:len(errMsg)]
			http.Error(w, errMsg , http.StatusForbidden)
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
		if(r.Message == "User Not found"){
			http.Error(w, "User Not found", http.StatusForbidden)
			return
		}else if(r.Message == "username/password does not match"){
			http.Error(w, "username/password does not match", http.StatusForbidden)
			return
		}else{
			// redirect
			un = r.Message
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}
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
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLetstalkClient(conn)
	//log.Printf("connection established")
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendLogout(ctx, &pb.LogoutRequest{Email: un})
	if err != nil {
		//log.Println(r.Message,"  ",err)
		errMsg:= err.Error()
		errMsg=errMsg[33:len(errMsg)]
		http.Error(w, errMsg , http.StatusForbidden)
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
	
	
    count:=0
	log.Println("method:", req.Method) //get request method
	req.ParseForm()

	talkf := req.Form["mytalk"]
	talk := talkf[0]
	log.Println(u.UserName)
	if len(dbSessions)== 0{
		
	}else {
		talka := myTalk{
			UserName: u.UserName,
			Talk: talk,
			Date: time.Now().Format("02-01-2006")+" "+time.Now().Format("15:04PM"),
		}
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
		r, err := c.SendTalk(ctx, &pb.TalkRequest{Email: talka.UserName, Talk: talka.Talk, Date: talka.Date})
		if err != nil {
			errMsg:= err.Error()
			errMsg=errMsg[33:len(errMsg)]
			http.Error(w, errMsg , http.StatusForbidden)
		}
		log.Println(r.Talk)
		count=len(dbMyTalk)
		dbMyTalk[count]=talka
		
	}
    http.Redirect(w, req, "/", http.StatusSeeOther)
}

func showTalk(w http.ResponseWriter, req *http.Request) {
	//get json api
    json.NewEncoder(w).Encode(talks)


}

func cancel(w http.ResponseWriter, req *http.Request) {

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
	r, err := c.SendCancel(ctx, &pb.CancelRequest{Email: ""})
	if err != nil {
		http.Error(w, r.Message, http.StatusForbidden)
		return
	}

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
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	var users []user
	for _,us:=range dbUsers{
		users=append(users, us)
	}
	json.NewEncoder(w).Encode(users)
}

func followothers(w http.ResponseWriter, req *http.Request) {
	
	var FollowPageVars followVariables
	var uname string
	var users []string
	for _,us:=range dbUsers{
		users=append(users, us.UserName)
	}

	if len(dbSessions)!=0{
		u = getUser(w,req)
		log.Println("Hello ", u.UserName)
		uname = u.UserName
	}else {
		log.Println("Username Not found")
		uname = ""
	}

	FollowPageVars = followVariables{
		UserName: uname,
		UserNames: users,
	}
	//log.Println("users:", users)
	//log.Println("users map:", dbUsers)


	if req.Method == http.MethodPost {
		var ud []string
		req.ParseForm()
		log.Println(req.Form)
		c, _ := req.Cookie("session")
		log.Println("------------------------follow function------------")
		for key, values := range req.Form {   // range over map
			for _, value := range values {    // range over []string
				log.Println(key, value)
				log.Println(c.Value)
				ud=append(ud, value)
			}
		}
		var session = dbSessions[c.Value]
		session.Following=ud
		updateTweets(session)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}else{
		tpl, err := template.ParseFiles("templates/follow.html") //parse the html file
		if err != nil { // if there is an error
			log.Print("template parsing error: ", err) // log it on terminal
		}
		err = tpl.Execute(w, FollowPageVars) //execute the template and pass it to index page
		if err != nil { // if there is an error
			log.Print("template executing error: ", err) //log it on terminal
		}
	}


}

func updateTweets(session session) {
	
}

func deleteTweets(){
	
}