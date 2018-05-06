package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../proto"
	"golang.org/x/crypto/bcrypt"
	"time"
	"errors"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"net/http"
)

const (
	port = ":8080"
)

var count=0
var talks = make([]*pb.Talk,count)
var userlist = make([]string,count)
var dbUsers = map[string]pb.User{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session
var dbusertalk = map[string][]*pb.Talk{}

type session struct {
	pb.User
	LastActivity time.Time
}

// server is used to implement server.
type server struct{}

var nodeList = map [int32]pb.Node{}


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
	dbusertalk[in.Talk.Email]=append(dbusertalk[in.Talk.Email],&tal)
	count= count + 1
	log.Println(talks)
	return &pb.TalkReply{Talk:dbusertalk[in.Talk.Email],Message: "Talk added successfully"}, nil

}

func (s *server) FollowUsers(ctx context.Context, in *pb.FollowUserRequest) (*pb.FollowUserReply, error) {

	//var usertalks = make([]*pb.Talk,count)
	log.Println("user list follow part:",in.Email)
	dbusertalk[in.Username]= updateTalk(in.Email,talks)
	log.Println("user talks updated:" ,dbusertalk[in.Username])
	return &pb.FollowUserReply{Talk:dbusertalk[in.Username], Username:in.Username }, nil

}

func (s *server) UnfollowUsers(ctx context.Context, in *pb.UnfollowUserRequest) (*pb.UnfollowUserReply, error) {

	//var usertalks = make([]*pb.Talk,count)

	return &pb.UnfollowUserReply{}, nil

}

func updateTalk(userlist []string,talk []*pb.Talk) ([]*pb.Talk){
	var usertalks = make([]*pb.Talk,count)
	usertalks=nil
	for _,us:=range userlist{
		for _,tal:=range talk{
			if tal.Email==us{
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
	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()
	//log.Printf("server created")
	//pb.RegisterLetstalkServer(s, &server{})
	//// Register reflection service on gRPC server.
	//reflection.Register(s)
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
	setupServer()
}


func setupServer() {

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server created")
	m := cmux.New(listener)
	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return grpcServe(grpcListener) })
	g.Go(func() error { return httpServe(httpListener) })
	g.Go(func() error { return m.Serve() })

	log.Println("run server:", g.Wait())
}


//func (*server) JoinSlave(ctx context.Context, in *pb.Node) (*pb.JoinReply, error) {
//
//	message:= fmt.Sprint(in.Id)
//
//	return &pb.JoinReply{Message: message}, nil
//}

func grpcServe(l net.Listener) error {
	s := grpc.NewServer()
	pb.RegisterLetstalkServer(s,&server{})
	return s.Serve(l)
}

func httpServe(l net.Listener) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/remove", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("remove"))
	})

	mux.HandleFunc("/join", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("join"))
	})

	s := &http.Server{Handler: mux}
	return s.Serve(l)
}





