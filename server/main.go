package main

import (
	"file-sharing-app/server/helpers"
	"file-sharing-app/server/models"
	"net"
)

const delimiter byte = 254 // ■

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		helpers.PrintErr(err)
		return
	}
	defer listener.Close()

	s := models.NewServer(listener)
	closeChan := make(chan bool)
	go s.RunServer(closeChan)

	// NOTE: Avoid program to finish, close with CTRL+C
	<- closeChan
}