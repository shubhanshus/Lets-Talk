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

func TestPostTalk(t *testing.T){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLetstalkClient(conn)
	var talk1 pb.Talk
	talk1.Email="test@test.com"
	talk1.Talk="Test message"
	talk1.Date=time.Now().Format("02-01-2006")+" "+time.Now().Format("15:04PM")
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SendTalk(ctx, &pb.TalkRequest{Talk: &talk1})
	if err != nil {
		t.Errorf("Post tweet Failed")
	}
	log.Println(resp.Message)
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
	if resp.Message=="Talk added successfully" {
		t.Errorf("User addtion failed")
	}
}

