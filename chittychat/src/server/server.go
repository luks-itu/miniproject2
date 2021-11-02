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
	port int32
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
		port: connection.Port,
		name: *connection.Name,
		client: pbclient.NewChittyClientClient(conn),
		conn: conn,
	}
	userConnections[connection.Port] = &newConnection
	announceJoin(newConnection)
	return &pbchat.Server_ResponseCode{Code: 204}, nil
}

func (s *ChittyChatServer) Leave(con context.Context, connection *pbchat.Server_Connection) (*pbchat.Server_ResponseCode, error) {
	connectionToRemove := userConnections[connection.Port]
	if connectionToRemove != nil {
		removeUserConnection(connection.Port);
	}
	return &pbchat.Server_ResponseCode{Code: 204}, nil
}

func (s *ChittyChatServer) Publish(con context.Context, message *pbchat.Server_Message) (*pbchat.Server_ResponseCode, error) {
	messageText := strings.TrimSpace(message.Text)
	if messageText == "" {
		return &pbchat.Server_ResponseCode{Code: 400}, nil;
	}
	if len(messageText) > 128 {
		des := "!!! Message too long !!!"
		return &pbchat.Server_ResponseCode{Code: 400, Description: &des}, nil;
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
	for _, user := range userConnections {
		_, err := user.client.Broadcast(context.Background(), &pbclient.Client_Message{
			Text: message,
			Lamport: 0,
		})
		if err != nil {
			fmt.Println(err.Error())
			removeUserConnection(user.port)
		}
	}
}

func announceLeave(user UserConnection) {
	fmt.Println(user.name + " left the chat.")
	for _, u := range userConnections {
		_, err := u.client.AnnounceLeave(context.Background(), &pbclient.Client_UserName{
			Name: user.name,
		})
		if err != nil {
			fmt.Println(err.Error())
			removeUserConnection(u.port)
		}
	}
}

func announceJoin(user UserConnection) {
	fmt.Println(user.name + " joined the chat.")
	for _, u := range userConnections {
		_, err := u.client.AnnounceJoin(context.Background(), &pbclient.Client_UserName{
			Name: user.name,
		})
		if err != nil {
			fmt.Println(err.Error())
			removeUserConnection(u.port)
		}
	}
}

func removeUserConnection(port int32){
	userToRemove := userConnections[port]
	userToRemove.conn.Close()
	delete(userConnections, port)
	announceLeave(*userToRemove)
}

func closeAllConnections() {
	for port := range userConnections {
		removeUserConnection(port)
	}
}
