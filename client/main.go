package main

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"time"
	"log"
	"encoding/json"
	"sort"
	"strings"
	//"os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../proto"
)

var tpl *template.Template
var u user
var talks []myTalk

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

/*func main() {
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
	http.HandleFunc("/followothers", followothers)
	//resource path
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

	//create json file
	//jsonFile, err := os.Create("json/talkList.json")
	//if err != nil{
    	//panic(err)
	//}
	//log.Println("creating file", jsonFile)
    


}*/

func index(w http.ResponseWriter, req *http.Request){
	var IndexPageVars pageVariables
	var uname string
	now := time.Now() // find the time right now

	if len(dbSessions)!=0{
		u = getUser(w,req)
		log.Println("Hello World", u.UserName)
		uname = u.UserName
	}else {
		log.Println("Username Not found")
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
/*
func showtweets(w http.ResponseWriter, req *http.Request) {
	var u user
	if alreadyLoggedIn(w, req) {
		u =getUser(w,req)
	}else{
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	var tweetUser tweet
	tweetUser=dbTweets[u.UserName]
	tpl, err := template.ParseFiles("templates/showtweets.html") //parse the html file
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it on terminal
	}
	showSessions() // for demonstration purposes
	tpl.Execute(w,tweetUser)
}

func home(w http.ResponseWriter, req *http.Request) {
	var u user
	if alreadyLoggedIn(w, req) {
		u =getUser(w,req)
	}else{
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	tpl, err := template.ParseFiles("templates/home.html") //parse the html file
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it on terminal
	}

	if req.Method == http.MethodPost {
		// get form values
		tweetMsg := req.FormValue("tweet")
		var tw tweet
		sID, _ := uuid.NewV4()
		tw = tweet{tweetMsg,time.Now(),u.UserName,sID.String()}
		// create session
		createSession(w,req,u)
		putTweet(req,&u,&tw)
		// redirect
		http.Redirect(w, req, "/showtweets", http.StatusSeeOther)
		return
	}

	showSessions() // for demonstration purposes
	tpl.Execute(w, "home.html")
}
*/
func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}
	var u user
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("email")
		p1 := req.FormValue("password1")
		//p2 := req.FormValue("password2")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		//// compare passwords
		//if p1!=p2{
		//	http.Error(w, "Both the passwords don't match", http.StatusForbidden)
		//	return
		//}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p1), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs, f, l}

		// create session
		createSession(w,req,u)

		dbUsers[un] = u

		//dial server
		address := "localhost:8080"

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewSignupClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SendSignup(ctx, &pb.SignupRequest{Email: un, Password1: p1, Firstname: f, Lastname: l})
		if err != nil {
			log.Fatalf("could not get: %v", err)
		}

		log.Println(r.Message)
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "signup.html", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/index", http.StatusSeeOther)
		return
	}
	var u user
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("name")
		p := req.FormValue("password")
		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			log.Println(un)
			log.Println(dbUsers)
			http.Error(w, "Username not found", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		createSession(w,req,u)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes

	tpl.ExecuteTemplate(w, "login.html", u)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 600) {
		go cleanSessions()
	}

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
		talks = append(talks, talka)

		log.Println(talks)
		count=len(dbMyTalk)
		dbMyTalk[count]=talka
		/*dont need to write file
		val, err := json.Marshal(talks)
		if err != nil {
	    	panic(err)
	    }
	    
	    log.Println(string(val))

	    //write json file
	    jsonFile, err := os.OpenFile("json/talkList.json",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		jsonFile.Write(val)
	    log.Println("list written to ", jsonFile.Name())*/
	}
    http.Redirect(w, req, "/", http.StatusSeeOther)
}

func showTalk(w http.ResponseWriter, req *http.Request) {
	//get json api
    json.NewEncoder(w).Encode(talks)


}

func cancel(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u = getUser(w,req)
	delete(dbUsers, u.UserName);
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	deleteTweets()
	http.SetCookie(w, c)
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
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
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
	var talks3 []myTalk
	log.Println("Followers:",session.Following)
	for _,uss:= range session.Following{
		for _,talk:=range dbMyTalk {
			if talk.UserName==uss{
				talks3=append(talks3, talk)
			}
		}
	}
	talks=talks3
	sort.Slice(talks, func(i, j int) bool {
		switch strings.Compare(talks[i].Date, talks[j].Date) {
		case -1:
			return true
		case 1:
			return false
		}
		return talks[i].Date > talks[j].Date
	})
	log.Println(talks)
}

func deleteTweets(){
	for i,talk:=range dbMyTalk {
		if talk.UserName==u.UserName{
			log.Println("inside the deletion loop")
			delete(dbMyTalk,i)
		}
	}
	var talks3 []myTalk
	talks=talks3
	for _,talk:=range dbMyTalk {
		talks=append(talks,talk)
		log.Println(talk)
	}
}