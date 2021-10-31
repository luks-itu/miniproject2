package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	pbclient "github.com/luks-itu/miniproject2/chittyclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//contains struct and methods for the clientserver

type ChittyClientServer struct {
	pbclient.UnimplementedChittyClientServer
	// Data goes here
}

func (s *ChittyClientServer) Broadcast(con context.Context, message *pbclient.Client_Message) (*pbclient.Client_ResponseCode, error) {
	// THIS IS PLACEHOLDER PLZ KILL IT
	fmt.Println(message.Text)
	return &pbclient.Client_ResponseCode{Code: 204}, nil;
}

func (s *ChittyClientServer) AnnounceJoin(con context.Context, userName *pbclient.Client_UserName) (*pbclient.Client_ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnounceJoin not implemented")
}

func (s *ChittyClientServer) AnnounceLeave(con context.Context, userName *pbclient.Client_UserName) (*pbclient.Client_ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnounceLeave not implemented")
}

func newServer() *ChittyClientServer {
	s := ChittyClientServer { }
	return &s
}

func Start(port int) {
	fmt.Println("STARTING CLIENT SERVER")
	lis, err := net.Listen("tcp", getTarget(port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pbclient.RegisterChittyClientServer(grpcServer, newServer())
	fmt.Println("CLIENT SERVER STARTED")
	grpcServer.Serve(lis)
}

func getTarget(port int) string {
	return "localhost:" + strconv.Itoa(port)
}
