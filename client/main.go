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
var u pb.User
var talks []*pb.Talk
var address = "localhost:8080"
var userLoggedIn =false
var un string
var uname string
var ud []string //follow list

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}



func index(w http.ResponseWriter, req *http.Request){
	var IndexPageVars pageVariables
	
	now := time.Now() // find the time right now
	if userLoggedIn{
		uname = u.Email
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

	if userLoggedIn{
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("email")
		p1 := req.FormValue("password1")
		p2 := req.FormValue("password2")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		if(un == "" || p1 == "" || p2 == "" || f == "" || l == ""){
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}
		if(p1 != p2){
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
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
		var user pb.User;
		user.Email= un
		user.Password1=p1
		user.Lastname=l
		user.Firstname=f
		r, err := c.SendSignup(ctx, &pb.SignupRequest{User:&user})
		if err != nil {
			errMsg:= err.Error()
			errMsg=errMsg[33:len(errMsg)]
			http.Error(w, errMsg , http.StatusForbidden)
			return
		}

		userLoggedIn=true
		log.Println(r.Message)
		u.Email = r.Message
		un = u.Email
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
		
	}
	tpl.ExecuteTemplate(w, "signup.html", u)
}

func login(w http.ResponseWriter, req *http.Request) {

	if userLoggedIn{
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("name")
		p := req.FormValue("password")
		
		if un == "" || p == ""{
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}


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
		u.Email=un
		log.Println(u.Email)
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
		return
	}
	log.Println(u.Email, r.Message)
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
	if !userLoggedIn{
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}
	//count:=0
	var talk1 pb.Talk
	log.Println("method:", req.Method) //get request method
	req.ParseForm()
	talkf := req.Form["mytalk"]
	talk := talkf[0]
	talk1.Talk=talk
	talk1.Email=u.Email
	talk1.Date=time.Now().Format("02-01-2006")+" "+time.Now().Format("15:04PM")
	log.Println(u.Email)
	if req.Method == http.MethodPost {

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
		r, err := c.SendTalk(ctx, &pb.TalkRequest{Talk: &talk1})
		if err != nil {
			errMsg:= err.Error()
			errMsg=errMsg[33:len(errMsg)]
			http.Error(w, errMsg , http.StatusForbidden)
			return
		}
		talks = r.Talk
		log.Println(talks)
		log.Println(r.Message)
		//log.Println(r.Talk)
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
	log.Println("UserName:",u.Email)
	defer conn.Close()
	c := pb.NewLetstalkClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendCancel(ctx, &pb.CancelRequest{Email: u.Email})
	if err != nil {
		http.Error(w, r.Message, http.StatusForbidden)
		return
	}
	cookie, _ := req.Cookie("session")
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	log.Println("Follow Tweets: ",r.Talk)
	talks=r.Talk
	http.SetCookie(w, cookie)
	userLoggedIn=false
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func cancelaccount(w http.ResponseWriter, req *http.Request){
	if !userLoggedIn{
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}
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
	if !userLoggedIn{
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}


	if req.Method == http.MethodPost {
		
		req.ParseForm()
		log.Println(req.Form)
		for _, values := range req.Form {   // range over map
			for _, value := range values {    // range over []string
				ud=append(ud, value)
			}
		}
		log.Println("follow list",ud)
		conn, err := grpc.Dial(address, grpc.WithInsecure())

		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewLetstalkClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.FollowUsers(ctx, &pb.FollowUserRequest{Username:u.Email,Email: ud})
		if err != nil {
			errMsg:= err.Error()
			errMsg=errMsg[33:len(errMsg)]
			http.Error(w, errMsg , http.StatusForbidden)
			return
		}
		log.Println("Follow Return:",r.Talk)
		log.Println(r.Username)
		talks=r.Talk
		
	}
	http.Redirect(w, req, "/", http.StatusSeeOther)
	
}

func followothers(w http.ResponseWriter, req *http.Request) {
	if !userLoggedIn{
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLetstalkClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendFollow(ctx, &pb.FollowRequest{})
	if err != nil {
		errMsg:= err.Error()
		errMsg=errMsg[33:len(errMsg)]
		http.Error(w, errMsg , http.StatusForbidden)
		return
	}
	log.Println(r.Userlist)
	log.Println(r.Message)
	log.Println("followed users",ud)

    var ulist []string
    for _,us:=range r.Userlist{
		if us!=""{
			ulist=append(ulist,us)
		}
	}
	check := make([]string, len(ulist))
	for i := 0; i < len(ulist); i++ {
      check[i] = "unchecked"
	}
	for i := 0; i < len(r.Userlist); i++{
		for _,fl:=range ud{
			if fl == r.Userlist[i]{
				check[i] = "checked"
				continue
			}else{
				check[i] = "unchecked"
				continue
			}
		}
	}
	log.Println("checked:", check)
	tpl, err := template.ParseFiles("templates/follow.html") //parse the html file
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it on terminal
	}
	var FollowPageVars followVariables
	FollowPageVars = followVariables{
		UserName: uname,
		UserNames: ulist,
		Checks: check,
	}
	log.Println("ulist",ulist)
	err = tpl.Execute(w, FollowPageVars) //execute the template and pass it to index page
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it on terminal
	}
	
}