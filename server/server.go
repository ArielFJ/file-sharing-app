package main

import (
	"bufio"
	"encoding/json"
	"file-sharing-app/server/handlers"
	"file-sharing-app/server/helpers"
	"file-sharing-app/server/models"
	"fmt"
	"net"
)

type Server struct {
	clients  map[*models.Client]bool
	channels map[string]*models.Channel
	listener net.Listener
}

func NewServer(l net.Listener) Server {
	return Server{
		clients:  make(map[*models.Client]bool),
		channels: make(map[string]*models.Channel),
		listener: l,
	}
}

func (s *Server) RunServer(closeChan chan bool) {
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

func (s *Server) handleConn(c net.Conn) {
	newClient := models.NewClient(c)
	s.clients[&newClient] = true

	helpers.Notify(fmt.Sprintf("%v has connected", newClient.GetIdentifier()))

	for {
		netData, err := bufio.NewReader(newClient.Conn).ReadString('\n')
		if err != nil {
			helpers.PrintErr(fmt.Errorf("READ: %v", err))
			handlers.HandleExit(&newClient, s.clients)
			return
		}

		var req models.Request
		err = json.Unmarshal([]byte(netData), &req)
		if err != nil {
			helpers.PrintErr(fmt.Errorf("JSON: %v", err))
			return
		}

		s.handleRequest(&newClient, req)
		if req.Command == models.EXIT {
			return
		}
	}
}

func (s *Server) handleRequest(c *models.Client, r models.Request) {
	helpers.Notify(fmt.Sprintf("%v CMD: %v %v", c.GetIdentifier(), r.Command, r.Payload))
	switch r.Command {
	case models.USERNAME:
		handlers.HandleUsername(c, r)
	case models.CHANNEL:
		handlers.HandleJoinChannel(c, r, s.channels)
	case models.SEND:
		handlers.HandleSendFileToChannel(c, r, s.channels)
	case models.MESSAGE:
		handlers.HandleMessageToChannel(c, r, s.channels)
	case models.LIST:
		handlers.HandleListChannels(c, r, s.channels)
	case models.STOP:
		handlers.HandleQuitChannel(c, r, s.channels)
	case models.EXIT:
		handlers.HandleExit(c, s.clients)
	case models.DATA:
		handlers.HandleGetData(c, s.channels)
	case models.CLIENTS:
		handlers.HandleGetClients(c, s.clients)
	default:
		res := models.NewResponseWithDefaultCode(models.ERROR, "Invalid Command")
		c.Send(res.ToBuffer())
	}
}
