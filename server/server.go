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

// server is used to implement server.
type server struct{}

// SendSignup implements signup request
func (s *server) SendSignup(ctx context.Context, in *pb.SignupRequest) (*pb.SignupReply, error) {
	log.Println(in.Email)

	bs, err := bcrypt.GenerateFromPassword([]byte(in.Password1), bcrypt.MinCost)
	if err != nil {
		log.Fatalf("failed to resolve: %v", err)
	}
	u := user{
		UserName:in.Email,
		First:in.Firstname,
		Last:in.Lastname,
		Password:bs,
	}
	log.Printf(" ",u.UserName,u.First)
	if _, ok := dbUsers[in.Email]; ok {
		return &pb.SignupReply{Message: "User Exists" + in.Email}, errors.New("user already exists")
	}
	log.Printf("user does not exist")
	dbUsers[in.Email] = u
	//sID, _ := uuid.NewV4()
	dbSessions[in.Email]=session{u, time.Now(),"",false,nil}
	log.Println("user addition successful")
	return &pb.SignupReply{Message: "Signup succeed:" + in.Email, Sessionid:in.Email}, nil
}
// login request
func (s *server) SendLogin(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
		

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
	
	return &pb.CancelReply{Message: "SendCancel return:" + in.Email}, nil

}

// follow  request
func (s *server) SendFollow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowReply, error) {
	
	return &pb.FollowReply{Message: "SendFollow return:" + in.Email}, nil

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


