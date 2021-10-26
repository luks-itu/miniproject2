package server

import (
	"context"

	pb "github.com/luks-itu/miniproject2/chittyclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChittyClientServer struct {
	pb.UnimplementedChittyClientServer
	// Data goes here
}

func (s *ChittyClientServer) Broadcast(context.Context, *pb.Message) (*pb.ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}

func (s *ChittyClientServer) AnnounceJoin(context.Context, *pb.UserName) (*pb.ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnounceJoin not implemented")
}

func (s *ChittyClientServer) AnnounceLeave(context.Context, *pb.UserName) (*pb.ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnounceLeave not implemented")
}
