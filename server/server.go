package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../proto"
	"google.golang.org/grpc/reflection"
	"golang.org/x/crypto/bcrypt"
	"time"
	"errors"
)

const (
	port = ":8080"
)

var count=0
var talks = make([]*pb.Talk,count)
var userlist = make([]string,count)
var dbUsers = map[string]pb.User{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session

type session struct {
	pb.User
	LastActivity time.Time
}

// server is used to implement server.
type server struct{}

// SendSignup implements signup request
func (s *server) SendSignup(ctx context.Context, in *pb.SignupRequest) (*pb.SignupReply, error) {
	log.Println(in.User.Email)

	bs, err := bcrypt.GenerateFromPassword([]byte(in.User.Password1), bcrypt.MinCost)
	if err != nil {
		log.Fatalf("failed to resolve: %v", err)
	}
	u := pb.User{
		Email:in.User.Email,
		Firstname:in.User.Firstname,
		Lastname:in.User.Lastname,
		Password1:string(bs[:]),
	}
	//log.Printf(" ",u.UserName,u.First)
	if _, ok := dbUsers[in.User.Email]; ok {
		return &pb.SignupReply{Message: "User Exists" + in.User.Email}, errors.New("user already exists")
	}
	//log.Printf("user does not exist")
	dbUsers[in.User.Email] = u
	//sID, _ := uuid.NewV4()
	dbSessions[in.User.Email]=session{u, time.Now()}
	userlist=append(userlist,u.Email)
	log.Println("user addition successful")
	return &pb.SignupReply{Message:in.User.Email, Sessionid:in.User.Email}, nil
}
// login request
func (s *server) SendLogin(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {

	u, ok := dbUsers[in.Email]
	if !ok {
		log.Println(u)
		log.Println(dbUsers)
		return &pb.LoginReply{}, errors.New("user does not exist")
	}
	// does the entered password match the stored password?
	err := bcrypt.CompareHashAndPassword([]byte(u.Password1), []byte(in.Password1))
	if err != nil {
		return &pb.LoginReply{Message: "username/password does not match"}, errors.New("username/password does not match")
	}
	dbSessions[in.Email]=session{u, time.Now()}
	log.Println("server index user:",in.Email)
	return &pb.LoginReply{Message: in.Email}, nil


}

// SendLogout implements logout request
func (s *server) SendLogout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutReply, error) {
	log.Println(in.Email)


	//// clean up dbSessions
	//if time.Now().Sub(dbSessionsCleaned) > (time.Second * 600) {
	//	go cleanSessions()
	//}

	delete(dbSessions,in.Email)
	return &pb.LogoutReply{}, nil
}

// cancel account request
func (s *server) SendCancel(ctx context.Context, in *pb.CancelRequest) (*pb.CancelReply, error) {
	log.Println("Email",in.Email)

	delete(dbUsers,in.Email)
	delete(dbSessions,in.Email)
	updateduserlist:=make([]string,len(userlist))
	for _,us:=range userlist{
		if us!=in.Email{
			updateduserlist=append(updateduserlist,us)
		}
	}
	userlist=updateduserlist
	log.Println(dbUsers)
	log.Println(userlist)
	talks=deleteTalk(in.Email,talks)
	return &pb.CancelReply{Talk:talks,Message: "SendCancel return:" + in.Email}, nil

}

// follow  request
func (s *server) SendFollow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowReply, error) {

	//return &pb.FollowReply{Userlist:userlist, Message: "SendFollow return:" + in.Email}, nil
	log.Println("recieve users:",in.Email)
	return &pb.FollowReply{Userlist:userlist, Message: "SendFollow returns" }, nil

}


// talk  request
func (s *server) SendTalk(ctx context.Context, in *pb.TalkRequest) (*pb.TalkReply, error) {
	log.Println(in.Talk.Talk)
	tal:=pb.Talk{
		Talk:in.Talk.Talk,
		Date:in.Talk.Date,
		Email:in.Talk.Email,
	}
	talks=append(talks,&tal)
	count= count + 1
	log.Println(talks)
	return &pb.TalkReply{Talk:talks,Message: "SendTalk return:"}, nil

}

func (s *server) FollowUsers(ctx context.Context, in *pb.FollowUserRequest) (*pb.FollowUserReply, error) {

	//var usertalks = make([]*pb.Talk,count)
	var dbusertalk = map[string][]*pb.Talk{}
	dbusertalk[in.Username]= updateTalk(in.Email,talks)
	return &pb.FollowUserReply{Talk:dbusertalk[in.Username], Username:in.Username }, nil

}

func updateTalk(userlist []string,talk []*pb.Talk) ([]*pb.Talk){
	var usertalks = make([]*pb.Talk,count)
	usertalks=nil
	for _,us:=range userlist{
		for _,tal:=range talk{
			if tal.Email==us{
				log.Println("in update talk",us)
				usertalks=append(usertalks,tal)
			}
		}
	}
	log.Println(usertalks)
	return usertalks
}

func deleteTalk(username string,talk []*pb.Talk) ([]*pb.Talk){
	var usertalks = make([]*pb.Talk,count)
	usertalks=nil
	for _,tal:=range talk{
		if tal.Email!=username{
			log.Println("cancel account:",username)
			usertalks=append(usertalks,tal)
		}
	}
	log.Println(" after delete tweets:",usertalks)
	return usertalks
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Printf("server created")
	pb.RegisterLetstalkServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


