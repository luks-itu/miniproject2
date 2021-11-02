package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	pbclient "chittyclient"

	pbchat "github.com/luks-itu/miniproject2/chittychat"
	"google.golang.org/grpc"
)

type Lamport struct {
	Time int64
	Mu sync.Mutex
}

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
	lamport *Lamport
)

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

func (s *ChittyChatServer) Join(con context.Context, connection *pbchat.Server_Connection) (*pbchat.Server_ResponseCode, error) {
	SetLamport(connection.Lamport)
	IncrementLamport()
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
	announceJoin(newConnection, connection.Lamport)
	return &pbchat.Server_ResponseCode{Code: 204}, nil
}

func (s *ChittyChatServer) Leave(con context.Context, connection *pbchat.Server_Connection) (*pbchat.Server_ResponseCode, error) {
	SetLamport(connection.Lamport)
	IncrementLamport()
	connectionToRemove := userConnections[connection.Port]
	if connectionToRemove != nil {
		removeUserConnection(connection.Port, connection.Lamport);
	}
	return &pbchat.Server_ResponseCode{Code: 204}, nil
}

func (s *ChittyChatServer) Publish(con context.Context, message *pbchat.Server_Message) (*pbchat.Server_ResponseCode, error) {
	SetLamport(message.Lamport)
	IncrementLamport()
	messageText := strings.TrimSpace(message.Text)
	if messageText == "" {
		return &pbchat.Server_ResponseCode{Code: 400}, nil;
	}
	if len(messageText) > 128 {
		des := "!!! Message too long !!!"
		return &pbchat.Server_ResponseCode{Code: 400, Description: &des}, nil;
	}
	postMessage(fmt.Sprintf("[%v]: %v", userConnections[message.Port].name, message.Text), message.Lamport)
	return &pbchat.Server_ResponseCode{Code: 204}, nil;
}

func newServer() *ChittyChatServer {
	s := ChittyChatServer { }
	return &s
}

func Start(lp *Lamport) {
	defer closeAllConnections()
	lamport = lp
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

func postMessage(message string, lp int64) {
	fmt.Printf("%d:%s\n\r", lp, message)
	IncrementLamport()
	for _, user := range userConnections {
		_, err := user.client.Broadcast(context.Background(), &pbclient.Client_Message{
			Text: message,
			Lamport: lamport.Time,
		})
		if err != nil {
			fmt.Println(err.Error())
			removeUserConnection(user.port, lamport.Time)
		}
	}
}

func announceLeave(user UserConnection, lp int64) {
	fmt.Printf("%d:%s left the chat\n\r", lp, user.name)
	IncrementLamport()
	for _, u := range userConnections {
		_, err := u.client.AnnounceLeave(context.Background(), &pbclient.Client_UserName{
			Name: user.name,
			Lamport: lamport.Time,
		})
		if err != nil {
			fmt.Println(err.Error())
			removeUserConnection(u.port, lamport.Time)
		}
	}
}

func announceJoin(user UserConnection, lp int64) {
	fmt.Printf("%d:%s joined the chat\n\r", lp, user.name)
	IncrementLamport()
	for _, u := range userConnections {
		_, err := u.client.AnnounceJoin(context.Background(), &pbclient.Client_UserName{
			Name: user.name,
			Lamport: lamport.Time,
		})
		if err != nil {
			fmt.Println(err.Error())
			removeUserConnection(u.port, lamport.Time)
		}
	}
}

func removeUserConnection(port int32, lp int64){
	userToRemove := userConnections[port]
	userToRemove.conn.Close()
	delete(userConnections, port)
	announceLeave(*userToRemove, lp)
}

func closeAllConnections() {
	for port := range userConnections {
		removeUserConnection(port, lamport.Time)
	}
}
