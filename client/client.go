package main

import (
	"bufio"
	"encoding/json"
	"file-sharing-app/client/helpers"
	"fmt"
	"io"
	"net"
)

var inputPromptText = ">> "

type Client struct {
	conn          net.Conn
	username      string
	closeChan     chan bool
	isChannelMode bool
}

func NewClient(c net.Conn, channel chan bool) Client {
	return Client{
		conn:          c,
		username:      "anonymous",
		closeChan:     channel,
		isChannelMode: false,
	}
}

// func (c *Client) receive() {

// }

func (c *Client) setUsername(name string) {
	c.username = name
}

func (c *Client) send(data []byte) {
	c.conn.Write(append(data, '\n'))
}

func (c *Client) disconnect() {
	// c.conn.Close()
	c.closeChan <- true
}

// func (c *Client) connect() {

// }

func (c *Client) handleSession() {
	fmt.Printf("\n\nType %v to see how to interact with the server.\n\n", HELP)
	for {
		input, err := helpers.TakeInput(inputPromptText)
		// c.isChannelMode = false
		if err != nil {
			printErrPrefix("INPUT", err)
			continue
		}

		ok, req := BuildRequest(input)
		if !ok {
			printErrPrefix("REQ", fmt.Errorf(""))
			continue
		}

		jsonBytes, err := json.Marshal(req)
		if err != nil {
			printErrPrefix("JSON", err)
			continue
		}

		if req.Command == HELP {
			showHelp()
			continue
		}

		if c.isChannelMode {
			if req.Command == STOP {
				c.disconnectFromChannel(jsonBytes)
			}
			continue
		}

		if ok, msg := validateCommand(req); !ok {
			fmt.Println("*", msg)
			continue
		}

		// Send request to SERVER
		c.send(jsonBytes)

		response, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			printErrPrefix("RESPONSE", err)
			if err == io.EOF {
				break
			}

			continue
		}

		fmt.Printf("-> %v", response)

		if req.Command == EXIT {
			break
		}

		go c.tryStartChannelMode(req)
	}

	c.disconnect()
}

func (c *Client) tryStartChannelMode(req request) {
	if req.Command != CHANNEL {
		return
	}
	c.isChannelMode = true
	inputPromptText = fmt.Sprintf("%v >> ", req.Payload)

	for c.isChannelMode {
		netData, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			printErrPrefix("CHANNEL", err)
			continue
		}
		var res response
		err = json.Unmarshal([]byte(netData), &res)
		if err != nil {
			printErrPrefix("UNMARSHAL", err)
			continue
		}

		if res.Code == ERROR {
			return
		}

		fmt.Println("\n", res.String())
		fmt.Print(inputPromptText)
	}
}

func (c *Client) disconnectFromChannel(jsonReq []byte) {
	inputPromptText = ">> "
	c.send(jsonReq)
	c.isChannelMode = false
}
