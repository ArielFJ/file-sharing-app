package main

import (
	"fmt"
	"net"
)

const delimiter byte = 254 // ■

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		printErr(err)
	}
	defer conn.Close()

	closeChan := make(chan bool)

	client := NewClient(conn, closeChan)

	go client.handleSession()

	<- closeChan
}

func printErr(err error) {
	fmt.Printf("Error: %v", err)
}

func printErrPrefix(prefix string, err error) {
	fmt.Printf("Error %v: %v", prefix, err)
}
