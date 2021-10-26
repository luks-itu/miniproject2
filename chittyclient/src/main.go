package main

import (
	"context"
	"fmt"

	"github.com/luks-itu/miniproject2/chittyclient/src/server"
	"google.golang.org/grpc"

	pb "chittychat"
)

func main() {
	/// SETUP SERVER ///
	go server.Start()
	
	fmt.Print("Enter to continue")
	fmt.Scanln()

	/// SETUP CLIENT ///
	// gRPC channel
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(getTarget(), opts...)
	if err != nil {
		panic("Could not connect!")
	}
	defer conn.Close()

	// client stub
	client := pb.NewChittyChatClient(conn)

	myMessage := pb.Server_Message{Text: "HelloWorld :D", Lamport: 0}
	response, _ := client.Publish(context.Background(), &myMessage)
	fmt.Println(*response.Description)

	fmt.Print("Enter to stop")
	fmt.Scanln()
}

func getTarget() string {
	return "localhost:8080"
}
