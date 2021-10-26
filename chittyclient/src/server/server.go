package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/luks-itu/miniproject2/chittyclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChittyClientServer struct {
	pb.UnimplementedChittyClientServer
	// Data goes here
}

func (s *ChittyClientServer) Broadcast(con context.Context, message *pb.Client_Message) (*pb.Client_ResponseCode, error) {
	// THIS IS PLACEHOLDER PLZ KILL IT
	text := message.Text;
	return &pb.Client_ResponseCode{Code: 0, Description: &text}, nil;
}

func (s *ChittyClientServer) AnnounceJoin(con context.Context, userName *pb.Client_UserName) (*pb.Client_ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnounceJoin not implemented")
}

func (s *ChittyClientServer) AnnounceLeave(con context.Context, userName *pb.Client_UserName) (*pb.Client_ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnounceLeave not implemented")
}

func newServer() *ChittyClientServer {
	s := ChittyClientServer { }
	return &s
}

func Start() {
	fmt.Println("STARTING CLIENT SERVER")
	lis, err := net.Listen("tcp", getTarget())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterChittyClientServer(grpcServer, newServer())
	fmt.Println("CLIENT SERVER STARTED")
	grpcServer.Serve(lis)
}

func getTarget() string {
	return "localhost:8081"
}
