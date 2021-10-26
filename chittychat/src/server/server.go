package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/luks-itu/miniproject2/chittychat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChittyChatServer struct {
	pb.UnimplementedChittyChatServer
	// Data goes here
}

func (s *ChittyChatServer) Join(con context.Context, connection *pb.Server_Connection) (*pb.Server_ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}

func (s *ChittyChatServer) Leave(con context.Context, connection *pb.Server_Connection) (*pb.Server_ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}

func (s *ChittyChatServer) Publish(con context.Context, message *pb.Server_Message) (*pb.Server_ResponseCode, error) {
	// THIS IS PLACEHOLDER PLZ KILL IT
	text := message.Text;
	return &pb.Server_ResponseCode{Code: 0, Description: &text}, nil;
}

func newServer() *ChittyChatServer {
	s := ChittyChatServer { }
	return &s
}

func Start() {
	fmt.Println("STARTING CHAT SERVER")
	lis, err := net.Listen("tcp", getTarget())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterChittyChatServer(grpcServer, newServer())
	fmt.Println("CHAT SERVER STARTED")
	grpcServer.Serve(lis)
}

func getTarget() string {
	return "localhost:8080"
}
