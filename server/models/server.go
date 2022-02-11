package models

import (
	"bufio"
	"encoding/json"
	"file-sharing-app/server/helpers"
	"fmt"
	"net"
)

type server struct {
	clients  map[Client]bool
	channels map[string][]Client
	listener net.Listener
}

func NewServer(l net.Listener) server {
	return server{
		clients:  make(map[Client]bool),
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
			s.closeClientConn(newClient)
			return
		}

		var req request
		err = json.Unmarshal([]byte(netData), &req)
		if err != nil {
			helpers.PrintErr(fmt.Errorf("JSON: %v", err))
			return
		}

		s.handleRequest(&newClient, req)
		if req.Command == EXIT {
			return
		}

		// os.WriteFile("../file/"+myMessage.Filename, myMessage.Data, 0644)
	}
}

func (s *server) closeClientConn(client Client) {
	fmt.Println("Disconnecting user", client.username)
	client.send([]byte("You have been disconnected\n"))
	client.disconnect()
	delete(s.clients, client)
}

func (s *server) handleRequest(c *Client, r request) {
	switch r.Command {
	case USERNAME:
		s.username(c, r)
	case CHANNEL:

	case SEND:

	case MESSAGE:

	case LIST:

	case EXIT:
		s.closeClientConn(*c)
	default:
		c.conn.Write([]byte("Invalid Command\n"))
	}
}

func (s *server) username(c *Client, r request) {
	c.username = r.Payload
	response := fmt.Sprintf("Your new username is %v\n", c.username)
	c.conn.Write([]byte(response))
}
