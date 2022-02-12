package main

import (
	"file-sharing-app/client/helpers"
	"file-sharing-app/client/models"
	"net"
)

const delimiter byte = 254 // ■

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		helpers.PrintErr(err)
	}
	defer conn.Close()

	closeChan := make(chan bool)

	client := models.NewClient(conn, closeChan)

	go client.HandleSession()

	<- closeChan
}