package main

import (
	"fmt"
	"golang.org/x/net/context"
	pb "../proto"
	"log"
)

var node pb.Node

func (*server) JoinSlave(ctx context.Context, in *pb.Node) (*pb.JoinReply, error) {
	message:= fmt.Sprint(in.Id)
	node.Id=in.Id
	node.Port=in.Port
	nodeList[in.Id]=node
	log.Println(nodeList)
	return &pb.JoinReply{Message: message}, nil
}

/*
//prepare is used to synchronize servers
func (srv *server) Prepare(ctx context.Context, args *pb.PrepareArgs) (reply *pb.PrepareReply, err error) {

	reply = &pb.PrepareReply{}
	reply.View=int32(srv.currentView)
	reply.Success=false;
	if(int(args.View) < srv.currentView){
		return
	}
	if(int(args.Index)<=srv.commitIndex){
		return
	}
	if(int(args.PrimaryCommit)>srv.commitIndex){
		srv.commitIndex=int(args.PrimaryCommit)
	}
	if(int(args.Index)!=srv.opNo+1||int(args.View)>srv.currentView){
		fmt.Println("Debug: Server needs to recover")
		//log.Fatal("Debug: Server needs to recover")
		srv.status = RECOVERING
		PrimaryIndex:=GetPrimary(int(args.View), len(srv.peers))
		RecoveryInArgs:=pb.RecoveryArgs{
			View:args.View,
			Server:int32(srv.me),
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		RecoveryoutArgs, err := srv.peerRPC[PrimaryIndex].Recovery(ctx,&RecoveryInArgs)

		if(err==nil) {
			if(RecoveryoutArgs.Success){
				srv.log=RecoveryoutArgs.Entries
				srv.commitIndex=int(RecoveryoutArgs.PrimaryCommit)
				srv.currentView=int(RecoveryoutArgs.View)

				//StartRestoring Data
				userdata = make(map[string]User)
				for _,recoveredUser := range RecoveryoutArgs.Data{
					//recover user credentials
					userToRecover := User{username:recoveredUser.Username,password:recoveredUser.Password}
					//recover tweets for user
					for _,tweetToRecover := range recoveredUser.TweetList{
						recreatedTweet := tweet{text:tweetToRecover.Text}
						userToRecover.tweets = append(userToRecover.tweets,recreatedTweet)
					}
					//recover users followlist
					for _, followerToRecover := range  recoveredUser.Follows{
						userToRecover.follows[followerToRecover] = true
					}
					//add user to user data
					userdata[userToRecover.username]=userToRecover
				}

				srv.status = NORMAL
				srv.opNo=len(srv.log)-1
				//srv.commitIndex=int(args.PrimaryCommit)
				reply.Success=true
				fmt.Println(userdata)
				return reply, nil
			}else{
				return reply, errors.New("Debug: Error while recovering")

			}
		}else{
			return reply, errors.New("Debug: Error while recovering")
		}
	}
	//srv.commitIndex=args.PrimaryCommit
	if(int(args.Index)==len(srv.log)){
		srv.log = append(srv.log,args.Entry)
		srv.opNo=srv.opNo+1
		srv.commitIndex=int(args.PrimaryCommit)
		reply.Success=true
		return
	}
	return

}
*/