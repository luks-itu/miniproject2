package main

import (
	"google.golang.org/grpc"

	pb "github.com/luks-itu/miniproject2/chittychat"
)

func main() {
	// gRPC channel
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(getTarget(), opts...)
	if err != nil {
		panic("Could not connect!")
	}
	defer conn.Close()

	// client stub
	pb.NewBlabla
}

func getTarget() string {
	return "localhost:8080"
}