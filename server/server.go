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
	"sync"
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"os/signal"
	"syscall"
)

const (
	port = ":8082"
)

var peercount=0
var count=0
var talks = make([]*pb.Talk,count)
var userlist = make([]string,count)
var dbUsers = map[string]pb.User{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session
var dbusertalk = map[string][]*pb.Talk{}
var whoami=0

type session struct {
	pb.User
	LastActivity time.Time
}

// server is used to implement server.

var nodeList = map [int32]pb.Node{}
var userdata = make(map[string] User)

type server struct {
	mu             sync.Mutex // Lock to protect shared access to this peer's state
	peers          []string   // Ports of all peers
	peerRPC        [3]pb.LetstalkClient
	me             int      // this peer's index into peers[]
	currentView    int      // what this peer believes to be the current active view
	status         int      // the server's current status (NORMAL, VIEWCHANGE or RECOVERING)
	lastNormalView int      // the latest view which had a NORMAL status
	log            []string // the log of "commands"
	commitIndex    int      // all log entries <= commitIndex are considered to have been committed.
	opNo           int
	port 			int
}

const (
	NORMAL = iota
	VIEWCHANGE
	RECOVERING
)

type User struct {
	username string
	password string
	tweets   []pb.Talk
	follows  map[string]bool
}

// SendSignup implements signup request
func (s *server) SendSignup(ctx context.Context, in *pb.SignupRequest) (*pb.SignupReply, error) {
	log.Println(in.User.Email)
	if(peercount<=1){
		return &pb.SignupReply{Message: "User Exists" + in.User.Email}, errors.New("only one server is running. Please try again in sometime")
		os.Exit(3)
	}
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
	if(peercount<=1){
		return &pb.LoginReply{Message: ""}, errors.New("only one server is running. Please try again in sometime")
		os.Exit(3)
	}
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
	if(peercount<=1){
		return &pb.LogoutReply{Message: ""}, errors.New("only one server is running. Please try again in sometime")
		os.Exit(3)
	}
	delete(dbSessions,in.Email)
	return &pb.LogoutReply{}, nil
}

// cancel account request
func (s *server) SendCancel(ctx context.Context, in *pb.CancelRequest) (*pb.CancelReply, error) {
	log.Println("Email",in.Email)
	if(peercount<=1){
		return &pb.CancelReply{Message: ""}, errors.New("only one server is running. Please try again in sometime")
		os.Exit(3)
	}

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
	if(peercount<=1){
		return &pb.FollowReply{Message: ""}, errors.New("only one server is running. Please try again in sometime")
		os.Exit(3)
	}
	//return &pb.FollowReply{Userlist:userlist, Message: "SendFollow return:" + in.Email}, nil
	log.Println("recieve users:",in.Email)
	return &pb.FollowReply{Userlist:userlist, Message: "SendFollow returns" }, nil

}


// talk  request
func (s *server) SendTalk(ctx context.Context, in *pb.TalkRequest) (*pb.TalkReply, error) {
	if(peercount<=1){
		return &pb.TalkReply{Message: ""}, errors.New("only one server is running. Please try again in sometime")
		os.Exit(3)
	}
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
	if(peercount<=1){
		return &pb.FollowUserReply{Username: ""}, errors.New("only one server is running. Please try again in sometime")
		os.Exit(3)
	}
	//var usertalks = make([]*pb.Talk,count)
	log.Println("user list follow part:",in.Email)

	dbusertalk[in.Username]= updateTalk(in.Email,talks)
	log.Println("user talks updated:" ,dbusertalk[in.Username])
	return &pb.FollowUserReply{Talk:dbusertalk[in.Username], Username:in.Username }, nil

}

func (s *server) UnfollowUsers(ctx context.Context, in *pb.UnfollowUserRequest) (*pb.UnfollowUserReply, error) {

	//var usertalks = make([]*pb.Talk,count)
	if(peercount<=1){
		return &pb.UnfollowUserReply{Username: ""}, errors.New("only one server is running. Please try again in sometime")
		os.Exit(3)
	}
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
	//var srv2 http.Server
	srv :=getServerinfo()
	setupServer(srv)
	//os.Exit(exitPeer())
}

//var srv *server

func getServerinfo() (server){
	myid:=flag.Int("myid", 12345, "ip address to run this node on. default is 8083.")
	myport := flag.Int("myport", 10000, "ip address to run this node on. default is 8083.")
	flag.Parse()

	if(*myid==12345){
		fmt.Println("ServerID not entered exiting")
		os.Exit(2)
	}

	if(*myport==10000){
		fmt.Println("Port not entered exiting")
		os.Exit(2)
	}
	fmt.Println(*myport)
	fmt.Println(*myid)
	srv := &server{
		me:             *myid,
		currentView:    0,
		lastNormalView: 0,
		status:         NORMAL,
		opNo:           0,
	}
	srv.log = append(srv.log, "")
	srv.peers = append(srv.peers, ":8082")
	srv.peers = append(srv.peers, ":8083")
	srv.peers = append(srv.peers, ":8084")

	if (*myid!=0){
		joinPeer()
	}else {
		peercount++
	}
	whoami=*myid
	return *srv
}

func checkOtherServers(srv server){
		for index, port := range srv.peers {
			conn, err := grpc.Dial(port, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v to port %s", err, port)
			}
			defer conn.Close()
			//c := pb.NewGreeterClient(conn)
			srv.peerRPC[index] = pb.NewLetstalkClient(conn)
		}

		for index, rpccaller := range srv.peerRPC {
			if index != srv.me {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				reply, err := rpccaller.WhoIsPrimary(ctx, &pb.WhoisPrimaryRequest{})
				if err != nil {
					fmt.Printf("Server %d returned an error %v \n", index, err)
				} else {
					fmt.Printf("Server %d replied that the primary is %d \n", index, reply.Index)
				}
			}
		}

}

func setupServer(srv server) {
	listener, err := net.Listen("tcp", srv.peers[srv.me])
	log.Println("myport:",srv.peers[srv.me])
	if err != nil {
		log.Fatal(err)
	}
	//checkOtherServers(srv)
	log.Printf("server created")
	m := cmux.New(listener)
	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return grpcServe(grpcListener) })
	g.Go(func() error { return httpServe(httpListener) })
	g.Go(func() error { return gracefulShutdown()	})
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
		peercount--
		fmt.Println("Peer Count:",peercount)
		w.Write([]byte("remove"))
	})

	mux.HandleFunc("/join", func(w http.ResponseWriter, _ *http.Request) {
		peercount++
		fmt.Println("Peer Count:",peercount)
		w.Write([]byte("joined"))
	})

	s := &http.Server{Handler: mux}
	return s.Serve(l)
}


func (s *server) WhoIsPrimary(ctx context.Context, in *pb.WhoisPrimaryRequest) (*pb.WhoIsPrimaryResponse, error) {

	primaryIndex := GetPrimary(s.currentView, len(s.peers))
	if primaryIndex > -1 && primaryIndex < len(s.peers) {
		return &pb.WhoIsPrimaryResponse{Index: int32(primaryIndex)}, nil
	}
	return &pb.WhoIsPrimaryResponse{Index: -1}, errors.New("Debug: Index of primary out of bounds")
}


//used to rpc and check if connection is alive
func (s *server) HeartBeat(ctx context.Context, in *pb.HeartBeatRequest) (*pb.HeartBeatResponse, error) {
	return &pb.HeartBeatResponse{IsAlive: true, CurrentView: int32(s.currentView)}, nil
}

//internal function call
func GetPrimary(view int, nservers int) int {
	return view % nservers
}

//prepare is used to synchronize servers
func (srv *server) Prepare(ctx context.Context, args *pb.PrepareArgs) (reply *pb.PrepareReply, err error) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	reply = &pb.PrepareReply{}
	reply.View = int32(srv.currentView)
	reply.Success = false
	if int(args.View) < srv.currentView {
		return
	}

	if int(args.Index) <= srv.commitIndex {
		return
	}
	if int(args.PrimaryCommit) > srv.commitIndex {
		srv.commitIndex = int(args.PrimaryCommit)
	}

	if int(args.Index) != srv.opNo+1 || int(args.View) > srv.currentView {
		fmt.Println("Debug: Server needs to recover")
		//log.Fatal("Debug: Server needs to recover")
		srv.status = RECOVERING
		PrimaryIndex := GetPrimary(int(args.View), len(srv.peers))
		RecoveryInArgs := pb.RecoveryArgs{
			View:   args.View,
			Server: int32(srv.me),
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		RecoveryOutArgs, err := srv.peerRPC[PrimaryIndex].Recovery(ctx, &RecoveryInArgs)

		if err == nil {
			if RecoveryOutArgs.Success {
				srv.log = RecoveryOutArgs.Entries
				srv.commitIndex = int(RecoveryOutArgs.PrimaryCommit)
				srv.currentView = int(RecoveryOutArgs.View)

				//StartRestoring Data
				userdata = make(map[string]User)
				for _, recoveredUser := range RecoveryOutArgs.Data {
					//recover user credentials
					userToRecover := User{username: recoveredUser.Username, password: recoveredUser.Password}
					userToRecover.follows = make(map[string]bool)
					//recover tweets for user
					for _, tweetToRecover := range recoveredUser.TweetList {
						recreatedTweet := pb.Talk{Talk: tweetToRecover.Talk}
						userToRecover.tweets = append(userToRecover.tweets, recreatedTweet)
					}
					//recover users followlist
					for _, followerToRecover := range recoveredUser.Follows {
						userToRecover.follows[followerToRecover] = true
					}
					//add user to user data
					userdata[userToRecover.username] = userToRecover
				}

				srv.status = NORMAL
				srv.opNo = len(srv.log) - 1
				//srv.commitIndex=int(args.PrimaryCommit)
				reply.Success = true
				fmt.Println(userdata)
				return reply, nil
			} else {
				return reply, errors.New("Debug: Error while recovering")

			}
		} else {
			return reply, errors.New("Debug: Error while recovering")
		}
	}
	//srv.commitIndex=args.PrimaryCommit
	if int(args.Index) == len(srv.log) {
		srv.log = append(srv.log, args.Entry)
		srv.opNo = srv.opNo + 1
		srv.commitIndex = int(args.PrimaryCommit)
		reply.Success = true
		return
	}
	return

}

//Start calls prepare and returns index to commit on. In this case with >1/2 prepare's start does not immediately write the commit index.
//The commit index is updated after > 1/2 Prepare+RPC
func (srv *server) Start(command string) (index int, view int, ok bool) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	// do not process command if status is not NORMAL
	// and if i am not the primary in the current view
	if srv.status != NORMAL {
		return -1, srv.currentView, false
	} else if GetPrimary(srv.currentView, len(srv.peers)) != srv.me {
		//Check if you're the Primary
		return -1, srv.currentView, false
	}

	//In case of failure, the command is still added to the log so we tell backup the new index
	srv.log = append(srv.log, command)
	srv.opNo = srv.opNo + 1
	count := 0

	//Calling all backups
	for i, rpcEndPoint := range srv.peerRPC {
		if i != srv.me {
			pointer := i
			inArgs := &pb.PrepareArgs{
				View:          int32(srv.currentView),
				PrimaryCommit: int32(srv.commitIndex),
				Index:         int32(srv.opNo),
				Entry:         command,
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			outArgs, err := rpcEndPoint.Prepare(ctx, inArgs)
			if err == nil {
				if outArgs.Success == true {
					count = count + 1
				} else {
					fmt.Printf("Debug: Prepare rpc to Server %d failed, error is : %s \n", pointer, err)
				}
			} else {
				fmt.Printf("Debug: Prepare rpc to Server %d failed, error is : %s \n", pointer, err)
			}

		}

	}

	//Determine the number of severs for majority
	length := len(srv.peers)

	//Check if majority calls have returned, consider Primary as committed
	if count >= length/2 {
		//	srv.log = append(srv.log,command)
		//srv.commitIndex=srv.opNo
		////	index=srv.commitIndex
		//ok=true
		ok = true
		index = srv.opNo
	} else {
		index = -1
		ok = false
	}
	return index, view, ok
}

func (srv *server) Recovery(ctx context.Context, args *pb.RecoveryArgs) (reply *pb.RecoveryReply, err error) {

	reply = &pb.RecoveryReply{}
	reply.View = int32(srv.currentView)
	reply.Entries = srv.log
	reply.PrimaryCommit = int32(srv.commitIndex)
	reply.Success = true

	//Start Initializing data
	for _, value := range userdata {
		//add users credentials to userobject
		userToAdd := &pb.UserData{Username: value.username, Password: value.password}

		//add users tweets to userobject
		for _, userTweet := range value.tweets {
			tweetToAdd := &pb.Talk{Talk: userTweet.Talk}
			userToAdd.TweetList = append(userToAdd.TweetList, tweetToAdd)
		}

		//add users followlist to userobject
		for userFollows := range value.follows {
			userToAdd.Follows = append(userToAdd.Follows, userFollows)
		}

		//Now finally after building the userobject append this to the recoveryReply
		reply.Data = append(reply.Data, userToAdd)

	}
	return reply, nil
	//return
}


var node pb.Node

func (*server) JoinSlave(ctx context.Context, in *pb.Node) (*pb.JoinReply, error) {
	message:= fmt.Sprint(in.Id)
	node.Id=in.Id
	node.Port=in.Port
	nodeList[in.Id]=node
	log.Println(nodeList)
	return &pb.JoinReply{Message: message}, nil
}

func joinPeer(){
	resp, err := http.Get("http://127.0.0.1:8082/join")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

}

func exitPeer() int{
	resp, err := http.Get("http://127.0.0.1:8082/remove")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	return 3
}

func gracefulShutdown() error{
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		if(whoami!=0){
			exitPeer()
		}
		time.Sleep(1*time.Second)
		fmt.Println("Sutting down gracefully.")
		// clean up here
		os.Exit(0)
	}()
	return errors.New("Func Error")
}