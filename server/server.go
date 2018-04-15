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
	if _, ok := dbUsers[in.Email]; ok {
		return &pb.SignupReply{Message: "User Exists" + in.Email}, errors.New("user already exists")
	}

	dbUsers[in.Email] = u
	//sID, _ := uuid.NewV4()
	dbSessions[in.Email]=session{u, time.Now(),"",false,nil}
	return &pb.SignupReply{Message: "Signup succeed:" + in.Email, Sessionid:in.Email}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Printf("server created")
	pb.RegisterSignupServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


