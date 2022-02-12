package models

import (
	"bufio"
	"encoding/json"
	"file-sharing-app/server/helpers"
	"fmt"
	"net"
	"strings"
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

		go s.handleConn(conn)
	}

	// NOTE: Cause the program to wait without closing
	closeChan <- true
}

func (s *server) handleConn(c net.Conn) {
	newClient := NewClient(c)
	s.clients[newClient] = true

	helpers.Notify(fmt.Sprintf("%v has connected",newClient.getIdentifier()))

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
	helpers.Notify("Disconnecting user " + client.username)

	res := NewResponse(OK, "You have been disconnected")
	client.send(res.ToBuffer())
	client.disconnect()
	delete(s.clients, *client)
}

func (s *server) handleRequest(c *Client, r request) {
	helpers.Notify(fmt.Sprintf("%v CMD: %v %v", c.getIdentifier(), r.Command, r.Payload))
	switch r.Command {
	case USERNAME:
		s.username(c, r)
	case CHANNEL:
		s.channel(c, r)
	case SEND:

	case MESSAGE:
		s.message(c, r)
	case LIST:
		s.list(c, r)
	case STOP:
		s.quitChannel(c, r)
	case EXIT:
		s.closeClientConn(c)
	default:
		c.conn.Write([]byte("Invalid Command\n"))
	}
}

func (s *server) username(c *Client, r request) {
	oldUsername := c.username
	response := NewResponse(OK, fmt.Sprintf("Current username: %v\n", c.username))
	if len(r.Payload) > 0 {
		c.username = r.Payload
		response.Result = fmt.Sprintf("New username: %v\n", c.username)
		helpers.Notify(fmt.Sprintf("%v change its username from %q", c.getIdentifier(),oldUsername))
	}
	c.send(response.ToBuffer())
}

func (s *server) channel(c *Client, r request) {
	chanName := r.Payload
	if len(chanName) < 1 {
		res := NewResponse(ERROR, "Channel must have an identifier")
		c.send(res.ToBuffer())
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

	helpers.Notify(fmt.Sprintf("%v connected to channel %q", c.getIdentifier(), c.currentChannel))

	res := NewResponse(OK, fmt.Sprintf("%v added to channel %q\n", c.username, chanName))
	c.send(res.ToBuffer())
}

func (s *server) quitChannel(c *Client, r request) {
	res := NewResponse(ERROR, fmt.Sprintf("%v removed from channel %q\n", c.username, c.currentChannel))
	currentClients, channelExists := s.channels[c.currentChannel]
	if channelExists {
		// Remove the client from the channel
		clients := []*Client{}
		for _, client := range currentClients {
			if client != c {
				clients = append(clients, client)
			}
		}
		s.channels[c.currentChannel] = clients
	}
	
	// Close the channel if it doesn't have users
	if len(s.channels[c.currentChannel]) == 0 {
		delete(s.channels, c.currentChannel)
	}

	helpers.Notify(fmt.Sprintf("%v left channel %q", c.getIdentifier(), c.currentChannel))
	c.currentChannel = ""

	c.send(res.ToBuffer())
}

func (s *server) list(c *Client, r request) {
	channelsText := "Available Channels:\n"
	for chanName, clients := range s.channels {
		channelsText += fmt.Sprintf("\t - %v (%v clients)\n", chanName, len(clients))
	}

	if len(s.channels) == 0 {
		channelsText = "No Available Channels. Run help to see how to create one."
	}

	res := NewResponse(OK, channelsText)
	c.send(res.ToBuffer())
}

func (s *server) message(c *Client, r request) {
	args := strings.Split(strings.TrimSpace(r.Payload), " ")
	chanName := strings.TrimSpace(args[0])
	realPayload := strings.TrimSpace(strings.Join(args[1:], " "))

	clients, channelExists := s.channels[chanName]
	if !channelExists {
		res := NewResponse(ERROR, fmt.Sprintf("Channel %v does not exists", chanName))
		c.send(res.ToBuffer())
		return
	}

	for _, client := range clients {
		if client != c {
			res := NewResponse(OK, fmt.Sprintf("MSG from %v: %v", c.username, realPayload))
			client.send(res.ToBuffer())
		}
	}

	senderResponse := NewResponse(OK, "Message sent to channel " + chanName)
	c.send(senderResponse.ToBuffer())
}