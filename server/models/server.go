package models

import (
	"bufio"
	"encoding/json"
	"file-sharing-app/server/helpers"
	"fmt"
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

		fmt.Printf("%v has connected\n", conn.RemoteAddr())

		go s.handleConn(conn)
	}

	// NOTE: Cause the program to wait without closing
	closeChan <- true
}

func (s *server) handleConn(c net.Conn) {
	newClient := NewClient(c)
	s.clients[newClient] = true
	for {

		// netData, err := ioutil.ReadAll(newClient.conn) For reading files
		netData, err := bufio.NewReader(newClient.conn).ReadString('\n')
		if err != nil {
			helpers.PrintErr(fmt.Errorf("READ: %v", err))
			closeClientConn(newClient, s)
			return
		}

		var myMessage message
		err = json.Unmarshal([]byte(netData), &myMessage)
		if err != nil {
			helpers.PrintErr(fmt.Errorf("JSON: %v", err))
			return
		}

		if myMessage.Command == EXIT {
			closeClientConn(newClient, s)
			return
		}
		
		fmt.Println("MSG", myMessage.String())
		// os.WriteFile("../file/"+myMessage.Filename, myMessage.Data, 0644)
	}
}

func closeClientConn(client Client, server *server) {
	fmt.Println("Disconnecting user", client.username)
	client.send([]byte("You have been disconnected\n"))
	client.disconnect()
	delete(server.clients, client)
}