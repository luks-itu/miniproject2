package main

import (
	"fmt"
	//"sync"

	"github.com/luks-itu/miniproject2/chittychat/src/server"
)

func main() {
	/// SETUP ACTUAL SERVER ///
	go server.Start()

	fmt.Println("Enter to exit")
	fmt.Scanln()
}

func getTarget() string {
	return "localhost:8081"
}
