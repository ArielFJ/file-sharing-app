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
	channels map[string][]*Client
	listener net.Listener
}

func NewServer(l net.Listener) server {
	return server{
		clients:  make(map[Client]bool),
		channels: make(map[string][]*Client),
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
			s.closeClientConn(&newClient)
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

func (s *server) closeClientConn(client *Client) {
	fmt.Println("Disconnecting user", client.username)
	client.send([]byte("You have been disconnected\n"))
	client.disconnect()
	delete(s.clients, *client)
}

func (s *server) handleRequest(c *Client, r request) {
	switch r.Command {
	case USERNAME:
		s.username(c, r)
	case CHANNEL:
		s.channel(c, r)
	case SEND:

	case MESSAGE:

	case LIST:

	case STOP:
		s.quitChannel(c, r)
	case EXIT:
		s.closeClientConn(c)
	default:
		c.conn.Write([]byte("Invalid Command\n"))
	}
}

func (s *server) username(c *Client, r request) {
	response := fmt.Sprintf("Current username: %v\n", c.username)
	if len(r.Payload) > 0 {
		c.username = r.Payload
		response = fmt.Sprintf("New username: %v\n", c.username)
	}
	c.conn.Write([]byte(response))
}

func (s *server) channel(c *Client, r request) {
	chanName := r.Payload
	if len(chanName) < 1 {
		c.conn.Write([]byte("Channel must have an identifier"))
		return
	}
	clientsInChannel, exists := s.channels[chanName]
	var clients []*Client
	if !exists {
		clients = []*Client{c}
	} else {
		clients = append(clientsInChannel, c)
	}

	s.channels[chanName] = clients
	c.currentChannel = chanName

	res := NewResponse(OK, fmt.Sprintf("%v added to channel %v\n", c.username, chanName))
	c.send(res.ToBuffer())
}

func (s *server) quitChannel(c *Client, r request) {
	res := NewResponse(ERROR, fmt.Sprintf("%v removed from channel %v\n", c.username, c.currentChannel))
	_, exists := s.channels[c.currentChannel]
	if exists {
		delete(s.channels, c.currentChannel)
	}
	c.currentChannel = ""
	c.send(res.ToBuffer())
}
