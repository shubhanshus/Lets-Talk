package main

import (
  "html/template"
  "log"
  "net/http"
  "time"
)

func test() {
  //html path
	http.HandleFunc("/", Index)
  http.HandleFunc("/login", Login)
  http.HandleFunc("/signup", Signup)
  //resource path
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
  http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
  http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Index(w http.ResponseWriter, r *http.Request){

    now := time.Now() // find the time right now
    IndexPageVars := pageVariables{ //store the date and time in a struct
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04PM"),
    }

    t, err := template.ParseFiles("templates/index.html") //parse the html file
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it on terminal
  	}
    err = t.Execute(w, IndexPageVars) //execute the template and pass it to index page
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it on terminal
  	}
}

func Login(w http.ResponseWriter, r *http.Request){

    LoginPageVars := pageVariables{ //store the date and time in a struct
      
    }
  
    t, err := template.ParseFiles("templates/login.html") //parse the html file 
    if err != nil { // if there is an error
      log.Print("template parsing error: ", err) // log it on terminal
    }
    err = t.Execute(w, LoginPageVars) //execute the template and pass it to index page
    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it on terminal
    }
    
}

func Signup(w http.ResponseWriter, r *http.Request){

    SignupPageVars := pageVariables{ //store the date and time in a struct
      
    }
  
    t, err := template.ParseFiles("templates/signup.html") //parse the html file 
    if err != nil { // if there is an error
      log.Print("template parsing error: ", err) // log it on terminal
    }
    err = t.Execute(w, SignupPageVars) //execute the template and pass it to index page
    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it on terminal
    }
    
}