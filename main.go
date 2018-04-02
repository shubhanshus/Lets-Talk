package main

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"time"
	"log"
	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/home", home)

	http.HandleFunc("/showtweets", showtweets)

	//resource path
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}


func index(w http.ResponseWriter, r *http.Request){
	var u user
	var IndexPageVars pageVariables
	var uname string
	now := time.Now() // find the time right now
	if len(dbSessions)!=0{
		u = getUser(w,r)
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
		p2 := req.FormValue("password2")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// compare passwords
		if p1!=p2{
			http.Error(w, "Both the passwords don't match", http.StatusForbidden)
			return
		}
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

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "signup.html", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}
	var u user
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
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
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 300) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
