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
	lamport server.Lamport
)

func main() {
	lamport = server.Lamport{Time: 0}

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

	go server.Start(port, &lamport)

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
	server.IncrementLamport()
	response, err := client.Join(context.Background(), &pbchat.Server_Connection{
		Port: int32(port),
		Name: &username,
		Lamport: lamport.Time,
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
	fmt.Println("\rClient started. Ready to chat.")
	fmt.Print("[ToAll]> ")
	for {
		var userMessage string
		if scanner.Scan() {
			if scanner.Text() != "" {
				fmt.Print("\033[F")
			}
			if stopAllSending {
				return
			}
			userMessage = scanner.Text()
			if strings.TrimSpace(userMessage) == "" {
				continue
			}
		}

		server.IncrementLamport()
		message := pbchat.Server_Message{
			Text: userMessage,
			Lamport: lamport.Time,
			Port: int32(port),
		}
		_, err := client.Publish(context.Background(), &message)
		if err != nil {
			fmt.Println("!!! Error sending the message !!!")
			fmt.Println(err.Error())
		} else {
			if response.Code == 400 {
				fmt.Println(response.Description)
			}
		}
	}
}

func leaveServer(client pbchat.ChittyChatClient) {
	stopAllSending = true;
	fmt.Println("Leaving server...")
	server.IncrementLamport()
	response, err := client.Leave(context.Background(), &pbchat.Server_Connection{
		Port: int32(port),
		Lamport: lamport.Time,
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(response.Code)
}
