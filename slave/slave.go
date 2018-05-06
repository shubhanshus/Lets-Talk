package main

import (
	"flag"
	"math/rand"
	"time"
	"fmt"
	"google.golang.org/grpc"
	pb "../proto"
	"golang.org/x/net/context"

)

type Node struct {
	Id int32
	Port int32
}

var n1 Node

func main(){
	myport := flag.Int("myport", 8083, "ip address to run this node on. default is 8083.")
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	myid := rand.Int31n(99999999)
	fmt.Println(*myport)
	fmt.Println(myid)
	var porti32 int32
	porti32 = int32(*myport)
	n1.Port=porti32
	n1.Id=myid
	address := "localhost:8082"
	grpcReturn, err := callGrpcServer(address)
	if err != nil {
		fmt.Println("gRPC failed:", err)
	}
	fmt.Printf("grpc: %s", grpcReturn)
	//http.ListenAndServe(":"+*myport, nil)
	//http.HandleFunc("/login", login)
	//http.HandleFunc("/signup", signup)
}


func callGrpcServer(address string) (string, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cli := pb.NewLetstalkClient(conn)
	var node pb.Node
	node.Port=n1.Port
	node.Id=n1.Id
	r, err := cli.JoinSlave(context.Background(), &node)
	if err != nil {
		return "", err
	}
	return r.Message, nil
}
