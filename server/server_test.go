package main

import ("testing"
pb "../proto"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)
const address = "localhost:8080"

func TestSingup(t *testing.T){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	c := pb.NewLetstalkClient(conn)
	defer conn.Close()

	var u pb.User
	u.Email="test@test.com"
	u.Password1="test"
	u.Lastname="test"
	u.Firstname="test"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.SendSignup(ctx, &pb.SignupRequest{User:&u})
	if err != nil {
		t.Errorf("Signup Failed")
	}
	if resp.Message != u.Email {
		t.Errorf("User addtion failed")
	}

}


func TestLogin(t *testing.T){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	c := pb.NewLetstalkClient(conn)
	defer conn.Close()

	var u pb.User
	u.Email="test@test.com"
	u.Password1="test"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err:=c.SendLogin(ctx,&pb.LoginRequest{Email:u.Email,Password1:u.Password1})
	if err != nil {
		t.Errorf("Login Failed")
	}
	log.Println(resp.Message)
	log.Println(u.Email)
	if resp.Message=="" {
		t.Errorf("User addtion failed")
	}
}



func TestAccountCancel(t *testing.T){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	c := pb.NewLetstalkClient(conn)
	defer conn.Close()

	var u pb.User
	u.Email="test@test.com"
	u.Password1="test"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err:=c.SendCancel(ctx,&pb.CancelRequest{Email:u.Email})
	if err != nil {
		t.Errorf("Account Cancel Failed")
	}
	log.Println(resp.Message)
	if resp.Message=="" {
		t.Errorf("User addtion failed")
	}
}

