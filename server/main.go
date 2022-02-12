package main

import (
	"file-sharing-app/server/helpers"
	"file-sharing-app/server/models"
	"fmt"
	"net"
)

const delimiter byte = 254 // ■
const defaultPort = 8000

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", defaultPort))
	if err != nil {
		helpers.PrintErr(err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server ready at 127.0.0.1:%v\n", defaultPort)

	s := models.NewServer(listener)
	closeChan := make(chan bool)
	go s.RunServer(closeChan)

	// NOTE: Avoid program to finish, close with CTRL+C
	<- closeChan
}