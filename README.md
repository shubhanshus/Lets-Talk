# distributed_systems_final
Project Url: https://github.com/caocasey/Let-s-talk

Team members:
Yingxi Cao(yc2539) and Shubhanshu Surana(ss11012)

This is the final project for distributed systems class. 

How to run on MacOS:
export GOPATH=/Users/XXX/Let-s-talk-master/
To start run model.go main.go session.go

To run test cases:
go test -v

How to run on Window:
go build
./distributed_systems_final.exe

Directory
css:
Store all the css or bootstrap file 
js:
Store all the javascript files
img:
Store all the images
pkg:
External packages

templates:
Store all the html pages

Install package:
go get -u golang.org/x/crypto/bcrypt


Router:
RESTFUL Api for post list:
http://localhost:8080/list
Ajax get json and parse data to html
http://localhost:8080/
http://localhost:8080/login
http://localhost:8080/signup
http://localhost:8080/logout
http://localhost:8080/cancelaccount
http://localhost:8080/postTalk
http://localhost:8080/showTalk
http://localhost:8080/followothers


Features:
Login
Signup (password is encrypted)
Cancel your account
Post talks
View others talks
Like others talks

Additonal features:
Share to twitter and facebook
Follow other users and see corresponding posts
Encryption password


Part2- Separate front-end and backend
Front server port 8081
Back server port 8080
Communicate through RPC and channel

