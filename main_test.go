package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"golang.org/x/crypto/bcrypt"
)

func TestHealthCheckHandler(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello server is running")
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}
	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}

func TestSignupChcek(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello server is running")
	}))
	defer ts.Close()

	resp, err := http.Post(ts.URL + "/signup","application/x-www-form-urlencoded",nil)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be redirected, got %d", resp.StatusCode)
	}
	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}

func  TestUserSignupCheck(t *testing.T){
	p1:="help"
	bs, err := bcrypt.GenerateFromPassword([]byte(p1), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	var us1 =user{};
	var user= user{
		UserName:"jj@gmail.com",
		First:"JJ",
		Last:"Help",
		Password:bs,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createSession(w,r,user)
		us1=getUser(w,r)
		if len(us1.UserName)<1{
			t.Errorf("User addition failed")
		}
		dbUsers[user.UserName] = u
		cancel(w,r)
		if len(dbSessions)>0 ||len(dbUsers)>0{
			t.Errorf("User deletion failed")
		}
	}))
	defer ts.Close()

}