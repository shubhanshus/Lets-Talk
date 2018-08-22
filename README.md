# distributed_systems_final
Team members:
Yingxi Cao(yc2539) and Shubhanshu Surana(ss11012)

This is the final project for distributed systems class. 

Separate front-end and backend
Front server port 8081
Back server port 8080
Communicate through RPC and channel
Data replication across servers is achived using ViewStamp Replication 

How to run on MacOS :
export GOPATH=/Users/XXX/Let-s-talk-master/
To start client server: go to directory /client
go run model.go client.go session.go main.go

To start backend server: go to directory /server
go run server.go

How to run on Window:
go build inside both the client and server folder
run the ./client.exe and ./server.exe files

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




