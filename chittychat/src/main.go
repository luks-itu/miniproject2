package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/luks-itu/miniproject2/chittychat/src/server"
)

var (
	lamport server.Lamport
	logger  *log.Logger
)

func main() {
	lamport = server.Lamport{Time: 0}
	var buf bytes.Buffer
	logger = log.New(&buf, "LOG: ", log.Lshortfile)

	logf := func(str string) {
		logger.Output(2, str)
	}

	logf(fmt.Sprintf("=== LOG START: %v ===", time.Now().Format(time.RFC1123)))

	defer fmt.Print(&buf)

	/// SETUP ACTUAL SERVER ///
	go server.Start(&lamport, logger)

	fmt.Println("Enter to exit")
	fmt.Scanln()
	os.WriteFile("log.txt", buf.Bytes(), 0644)
}
