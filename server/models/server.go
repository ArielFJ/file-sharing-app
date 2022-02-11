package models

import (
	"encoding/json"
	"file-sharing-app/server/helpers"
	"fmt"
	"io/ioutil"
	"net"
)

type server struct {
	clients map[Client]bool
	listener net.Listener
}

func NewServer(l net.Listener) server {
	return server{
		clients: make(map[Client]bool),
		listener: l,
	}
}

func (s *server) Broadcast() {
	
}

func (s *server) RunServer(closeChan chan bool) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			helpers.PrintErr(err)
			continue
		}

		go s.handleConn(conn)
	}

	// NOTE: Cause the program to wait without closing
	closeChan <- true
}

func (s *server) handleConn(c net.Conn) {
	newClient := NewClient(c)
	s.clients[newClient] = true
	for {
		netData, err := ioutil.ReadAll(c)
		if err != nil {
			// printErr(err)
			closeClientConn(newClient, s)
			return
		}

		var myMessage message
		err = json.Unmarshal(netData, &myMessage)
		if err != nil {
			// helpers.PrintErr(fmt.Errorf("JSON: %v", err))
			return
		}

		// if strings.ToUpper(msg) == "EXIT" {
		// 	closeClientConn(client, s)
		// 	return
		// }
		
		// toWrite := msg[: len(msg) - 1]
		// toWrite = append(toWrite, '\n')
		// fmt.Println(toWrite)
		fmt.Println(myMessage.String())
		// os.WriteFile("../file/"+myMessage.Filename, myMessage.Data, 0644)
		// return
	}
}

func closeClientConn(client Client, server *server) {
	client.disconnect()
	delete(server.clients, client)
}