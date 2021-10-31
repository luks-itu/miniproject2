package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/luks-itu/miniproject2/chittyclient/src/server"
	"google.golang.org/grpc"

	pbchat "chittychat"
)

var (
	port int
	username string
	scanner *bufio.Scanner

	stopAllSending = false
)

func main() {
	/// SETUP SCANNER ///
	scanner = bufio.NewScanner(os.Stdin)

	/// SETUP CLIENT SERVER ///
	fmt.Print("[Port]> ")
	if scanner.Scan() {
		var err error
		port, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Invalid port")
		}
	}
	fmt.Print("[Username]> ")
	if scanner.Scan() {
		username = strings.TrimSpace(scanner.Text())
		if username == "" {
			panic("Invalid username")
		}
	}
	fmt.Printf("Logged in as: Port=%v, Username=%v\n\r", port, username)

	go server.Start(port)

	/// SETUP CLIENT CLIENT ///
	// gRPC channel
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(getTarget(), opts...)
	if err != nil {
		panic("Could not connect!")
	}
	defer conn.Close()

	// client stub
	client := pbchat.NewChittyChatClient(conn)

	sendMessageLoop(client)
}

func getTarget() string {
	return "localhost:8080"
}

func sendMessageLoop(client pbchat.ChittyChatClient) {
	// Join server
	response, err := client.Join(context.Background(), &pbchat.Server_Connection{
		Port: int32(port),
		Name: &username,
	})
	if err != nil {
		panic("!!! Join went wrong !!!")
	}
	if (response.Code == 409) {
		panic("Port rejected")
	}

	// Defer leave server
	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        leaveServer(client)
        os.Exit(1)
    }()

	// Main loop
	fmt.Println("Client started. Ready to chat.")
	for {
		var userMessage string
		fmt.Print("[ToAll]> ")
		if scanner.Scan() {
			if stopAllSending {
				return
			}
			userMessage = scanner.Text()
			if strings.TrimSpace(userMessage) == "" {
				continue
			}
		}

		message := pbchat.Server_Message{
			Text: userMessage,
			Lamport: 0,
			Port: int32(port),
		}
		response, err := client.Publish(context.Background(), &message)
		if err != nil {
			fmt.Println("!!! Error sending the message !!!")
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Response: %v \n\r", response.Code)
		}
	}
}

func leaveServer(client pbchat.ChittyChatClient) {
	stopAllSending = true;
	fmt.Println("Leaving server...")
	response, err := client.Leave(context.Background(), &pbchat.Server_Connection{
		Port: int32(port),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(response.Code)
}
