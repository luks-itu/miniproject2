package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"

	pbclient "github.com/luks-itu/miniproject2/chittyclient"
	"google.golang.org/grpc"
)

//contains struct and methods for the clientserver

type Lamport struct {
	Time int64
	Mu sync.Mutex
}

type ChittyClientServer struct {
	pbclient.UnimplementedChittyClientServer
	// Data goes here
}

var lamport *Lamport

func IncrementLamport() {
	lamport.Mu.Lock()
	defer lamport.Mu.Unlock()
	lamport.Time++
}

func SetLamport(time int64) {
	lamport.Mu.Lock()
	defer lamport.Mu.Unlock()
	if time > lamport.Time {
		lamport.Time = time
	}
}

func (s *ChittyClientServer) Broadcast(con context.Context, message *pbclient.Client_Message) (*pbclient.Client_ResponseCode, error) {
	SetLamport(message.Lamport)
	IncrementLamport()
	printMessageFromServer(fmt.Sprintf("%d:%s", message.Lamport, message.Text))
	return &pbclient.Client_ResponseCode{Code: 204}, nil;
}

func (s *ChittyClientServer) AnnounceJoin(con context.Context, userName *pbclient.Client_UserName) (*pbclient.Client_ResponseCode, error) {
	SetLamport(userName.Lamport)
	IncrementLamport()
	printMessageFromServer(fmt.Sprintf("%d:[%s joined the chat.]", userName.Lamport, userName.Name))
	return &pbclient.Client_ResponseCode{Code: 204}, nil;
}

func (s *ChittyClientServer) AnnounceLeave(con context.Context, userName *pbclient.Client_UserName) (*pbclient.Client_ResponseCode, error) {
	SetLamport(userName.Lamport)
	IncrementLamport()
	printMessageFromServer(fmt.Sprintf("%d:[%s left the chat.]", userName.Lamport, userName.Name))
	return &pbclient.Client_ResponseCode{Code: 204}, nil;
}

func newServer() *ChittyClientServer {
	s := ChittyClientServer { }
	return &s
}

func Start(port int, lp *Lamport) {
	lamport = lp
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

func printMessageFromServer(message string) {
	fmt.Print("\r")
	// Padded to make sure the "[ToAll]>" part is overwritten.
	fmt.Printf("%s           \n\r", message)
	fmt.Print("[ToAll]> ")
}
