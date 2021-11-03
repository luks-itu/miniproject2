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
	Mu   sync.Mutex
}

type ChittyChatServer struct {
	pbchat.UnimplementedChittyChatServer
	// Data goes here
}

type UserConnection struct {
	port   int32
	name   string
	client pbclient.ChittyClientClient
	conn   *grpc.ClientConn
}

var (
	userConnections = make(map[int32]*UserConnection)
	lamport         *Lamport
	logf            func(str string)
)

func IncrementLamport() {
	lamport.Mu.Lock()
	defer lamport.Mu.Unlock()
	lamport.Time++
	logf(fmt.Sprintf("<Lamport incremented to:%d>", lamport.Time))
}

func SetLamport(time int64) {
	lamport.Mu.Lock()
	defer lamport.Mu.Unlock()
	if time > lamport.Time {
		lamport.Time = time
		logf(fmt.Sprintf("<Lamport set to:%d>", lamport.Time))
	}
}

func (s *ChittyChatServer) Join(con context.Context, connection *pbchat.Server_Connection) (*pbchat.Server_ResponseCode, error) {
	logf(fmt.Sprintf(
		"Participant [Port:%d,Name:%s,Lamport:%d] joined Chitty-Chat at Lamport time %d",
		connection.Port,
		*connection.Name,
		connection.Lamport,
		lamport.Time,
	))
	SetLamport(connection.Lamport)
	IncrementLamport()
	if _, exists := userConnections[connection.Port]; exists {
		return &pbchat.Server_ResponseCode{Code: 409}, nil
	}
	/// SETUP Server CLIENT ///
	// gRPC channel
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(connection.Port)), opts...)
	if err != nil {
		panic("Could not connect!")
	}

	// server client stub
	newConnection := UserConnection{
		port:   connection.Port,
		name:   *connection.Name,
		client: pbclient.NewChittyClientClient(conn),
		conn:   conn,
	}
	userConnections[connection.Port] = &newConnection
	announceJoin(newConnection, connection.Lamport)
	return &pbchat.Server_ResponseCode{Code: 204}, nil
}

func (s *ChittyChatServer) Leave(con context.Context, connection *pbchat.Server_Connection) (*pbchat.Server_ResponseCode, error) {
	logf(fmt.Sprintf(
		// NOTE: This may be logged with the wrong lamport timestamp because of the
		// many empty messages sent by the client at the same time as the leave
		// message. This only affects this log line, not the actual lamport tracking,
		// so we do not lock it.
		"Participant [Port:%d,Lamport:%d] left Chitty-Chat at Lamport time %d",
		connection.Port,
		connection.Lamport,
		lamport.Time,
	))
	SetLamport(connection.Lamport)
	IncrementLamport()
	connectionToRemove := userConnections[connection.Port]
	if connectionToRemove != nil {
		removeUserConnection(connection.Port, connection.Lamport)
	}
	return &pbchat.Server_ResponseCode{Code: 204}, nil
}

func (s *ChittyChatServer) Publish(con context.Context, message *pbchat.Server_Message) (*pbchat.Server_ResponseCode, error) {
	logf(fmt.Sprintf(
		"Recieved from [Port:%d,Lamport:%d] \"%s\" at Lamport time %d",
		message.Port,
		message.Lamport,
		message.Text,
		lamport.Time,
	))
	SetLamport(message.Lamport)
	IncrementLamport()
	messageText := strings.TrimSpace(message.Text)
	if messageText == "" {
		return &pbchat.Server_ResponseCode{Code: 400}, nil
	} else if len(messageText) > 128 {
		des := "!!! Message too long !!!"
		return &pbchat.Server_ResponseCode{Code: 400, Description: &des}, nil
	} else {
		postMessage(fmt.Sprintf("[%v]: %v", userConnections[message.Port].name, message.Text), message.Lamport)
	}
	return &pbchat.Server_ResponseCode{Code: 204}, nil
}

func newServer() *ChittyChatServer {
	s := ChittyChatServer{}
	return &s
}

func Start(lp *Lamport, logger *log.Logger) {
	logf = func(str string) {
		logger.Output(2, str)
	}
	defer closeAllConnections()
	lamport = lp
	fmt.Println("STARTING CHAT SERVER")
	logf("STARTING CHAT SERVER")
	lis, err := net.Listen("tcp", getTarget())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pbchat.RegisterChittyChatServer(grpcServer, newServer())
	fmt.Println("CHAT SERVER STARTED")
	logf("CHAT SERVER STARTED")
	logf(fmt.Sprintf("<Initial Lamport:%d>", lp.Time))
	grpcServer.Serve(lis)
}

func getTarget() string {
	return "localhost:8080"
}

func postMessage(message string, lp int64) {
	fmt.Printf("%d:%s\n\r", lp, message)
	IncrementLamport()
	logf(fmt.Sprintf(
		"Broadcasting \"%s\" at Lamport time %d",
		message,
		lamport.Time,
	))
	for _, user := range userConnections {
		_, err := user.client.Broadcast(context.Background(), &pbclient.Client_Message{
			Text:    message,
			Lamport: lamport.Time,
		})
		if err != nil {
			fmt.Println(err.Error())
			removeUserConnection(user.port, lamport.Time)
		}
	}
}

func announceLeave(user UserConnection, lp int64) {
	message := fmt.Sprintf("%d:%s left the chat", lp, user.name)
	fmt.Println(message)
	IncrementLamport()
	logf(fmt.Sprintf(
		"Announcing that [Port:%d,Name:%s] has left the chat at Lamport time %d",
		user.port,
		user.name,
		lamport.Time,
	))
	for _, u := range userConnections {
		_, err := u.client.AnnounceLeave(context.Background(), &pbclient.Client_UserName{
			Name:    user.name,
			Lamport: lamport.Time,
		})
		if err != nil {
			fmt.Println(err.Error())
			removeUserConnection(u.port, lamport.Time)
		}
	}
}

func announceJoin(user UserConnection, lp int64) {
	message := fmt.Sprintf("%d:%s joined the chat", lp, user.name)
	fmt.Println(message)
	IncrementLamport()
	logf(fmt.Sprintf(
		"Announcing that [Port:%d,Name:%s] has joined the chat at Lamport time %d",
		user.port,
		user.name,
		lamport.Time,
	))
	for _, u := range userConnections {
		_, err := u.client.AnnounceJoin(context.Background(), &pbclient.Client_UserName{
			Name:    user.name,
			Lamport: lamport.Time,
		})
		if err != nil {
			fmt.Println(err.Error())
			removeUserConnection(u.port, lamport.Time)
		}
	}
}

func removeUserConnection(port int32, lp int64) {
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
