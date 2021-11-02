package main

import (
	"fmt"

	"github.com/luks-itu/miniproject2/chittychat/src/server"
)

var lamport server.Lamport

func main() {
	lamport = server.Lamport{Time: 0}

	/// SETUP ACTUAL SERVER ///
	go server.Start(&lamport)

	fmt.Println("Enter to exit")
	fmt.Scanln()
}
