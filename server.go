package main

import (
  "html/template"
  "log"
  "net/http"
  "time"
)
//Pass variables to html
type PageVariables struct {
	Date         string
	Time         string
}

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Index(w http.ResponseWriter, r *http.Request){

    now := time.Now() // find the time right now
    IndexPageVars := PageVariables{ //store the date and time in a struct
      Date: now.Format(),
      Time: now.Format(),
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