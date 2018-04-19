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

var talks[] pb.Talk

// server is used to implement server.
type server struct{}

// SendSignup implements signup request
func (s *server) SendSignup(ctx context.Context, in *pb.SignupRequest) (*pb.SignupReply, error) {
	log.Println(in.User.Email)

	bs, err := bcrypt.GenerateFromPassword([]byte(in.User.Password1), bcrypt.MinCost)
	if err != nil {
		log.Fatalf("failed to resolve: %v", err)
	}
	u := user{
		UserName:in.User.Email,
		First:in.User.Firstname,
		Last:in.User.Lastname,
		Password:bs,
	}
	//log.Printf(" ",u.UserName,u.First)
	if _, ok := dbUsers[in.User.Email]; ok {
		return &pb.SignupReply{Message: "User Exists" + in.User.Email}, errors.New("user already exists")
	}
	//log.Printf("user does not exist")
	dbUsers[in.User.Email] = u
	//sID, _ := uuid.NewV4()
	dbSessions[in.User.Email]=session{u, time.Now(),"",false,nil}
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
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(in.Password1))
	if err != nil {
		return &pb.LoginReply{Message: "username/password does not match"}, errors.New("username/password does not match")
	}
	dbSessions[in.Email]=session{u, time.Now(),"",false,nil}
	return &pb.LoginReply{Message: "SendLogin return:" + in.Email}, nil


}

// SendLogout implements logout request
func (s *server) SendLogout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutReply, error) {
	log.Println(in.Email)

	
	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 600) {
		go cleanSessions()
	}

	delete(dbSessions,in.Email)
	return &pb.LogoutReply{}, nil
}

// cancel account request
func (s *server) SendCancel(ctx context.Context, in *pb.CancelRequest) (*pb.CancelReply, error) {
	delete(dbUsers,in.Email)
	delete(dbSessions,in.Email)
	return &pb.CancelReply{Message: "SendCancel return:" + in.Email}, nil

}

// follow  request
func (s *server) SendFollow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowReply, error) {
	
	return &pb.FollowReply{Message: "SendFollow return:" + in.Email}, nil

}

// talk  request
func (s *server) SendTalk(ctx context.Context, in *pb.TalkRequest) (*pb.TalkReply, error) {
	talks=append(talks,*in.Talk)

	return &pb.TalkReply{Message: "SendTalk return:"}, nil

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


