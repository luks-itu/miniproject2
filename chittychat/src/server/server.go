package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	pbclient "chittyclient"

	pbchat "github.com/luks-itu/miniproject2/chittychat"
	"google.golang.org/grpc"
)

type ChittyChatServer struct {
	pbchat.UnimplementedChittyChatServer
	// Data goes here
}

type UserConnection struct {
	name string
	client pbclient.ChittyClientClient
	conn *grpc.ClientConn
}

var (
	userConnections = make(map[int32]*UserConnection)
)

func (s *ChittyChatServer) Join(con context.Context, connection *pbchat.Server_Connection) (*pbchat.Server_ResponseCode, error) {
	if _, exists := userConnections[connection.Port]; exists {
		return &pbchat.Server_ResponseCode{Code: 409}, nil
	}
	/// SETUP Server CLIENT ///
	// gRPC channel
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:" + strconv.Itoa(int(connection.Port)), opts...)
	if err != nil {
		panic("Could not connect!")
	}

	// server client stub
	newConnection := UserConnection{
		name: *connection.Name,
		client: pbclient.NewChittyClientClient(conn),
		conn: conn,
	}
	userConnections[connection.Port] = &newConnection
	postMessage(newConnection.name + " joined the chat.")
	return &pbchat.Server_ResponseCode{Code: 204}, nil
}

func (s *ChittyChatServer) Leave(con context.Context, connection *pbchat.Server_Connection) (*pbchat.Server_ResponseCode, error) {
	connectionToRemove := userConnections[connection.Port]
	if connectionToRemove != nil {
		postMessage(connectionToRemove.name + " left the chat.")
	}
	removeUserConnection(connection.Port);
	return &pbchat.Server_ResponseCode{Code: 204}, nil
}

func (s *ChittyChatServer) Publish(con context.Context, message *pbchat.Server_Message) (*pbchat.Server_ResponseCode, error) {
	if strings.TrimSpace(message.Text) == "" {
		return &pbchat.Server_ResponseCode{Code: 400}, nil;
	}
	postMessage(fmt.Sprintf("[%v]: %v", userConnections[message.Port].name, message.Text))
	return &pbchat.Server_ResponseCode{Code: 204}, nil;
}

func newServer() *ChittyChatServer {
	s := ChittyChatServer { }
	return &s
}

func Start() {
	defer closeAllConnections()
	fmt.Println("STARTING CHAT SERVER")
	lis, err := net.Listen("tcp", getTarget())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pbchat.RegisterChittyChatServer(grpcServer, newServer())
	fmt.Println("CHAT SERVER STARTED")
	grpcServer.Serve(lis)
}

func getTarget() string {
	return "localhost:8080"
}

func postMessage(message string) {
	fmt.Println(message)
}

func removeUserConnection(port int32){
	userConnections[port].conn.Close()
	delete(userConnections, port)
}

func closeAllConnections() {
	for port := range userConnections {
		removeUserConnection(port)
	}
}
