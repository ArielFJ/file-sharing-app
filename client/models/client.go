package models

import (
	"bufio"
	"encoding/json"
	"file-sharing-app/client/helpers"
	"fmt"
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

func (c *Client) send(data []byte) {
	c.conn.Write(append(data, '\n'))
}

func (c *Client) disconnect() {
	// c.conn.Close()
	c.closeChan <- true
}


func (c *Client) HandleSession() {
	fmt.Printf("\n\nType %v to see how to interact with the server.\n\n", HELP)
	for {
		input, err := helpers.TakeInput(inputPromptText)
		if err != nil {
			helpers.PrintErrPrefix("INPUT", err)
			continue
		}

		ok, req := BuildRequest(input)
		if !ok {
			helpers.PrintErrPrefix("REQ", fmt.Errorf(""))
			continue
		}

		jsonBytes, err := json.Marshal(req)
		if err != nil {
			helpers.PrintErrPrefix("JSON", err)
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

		res, mildError, fatalError := ReadFromConn(c.conn)
		if fatalError {
			break
		}

		if mildError {
			continue
		}

		fmt.Printf("-> %v\n", res.Result)

		if req.Command == EXIT {
			break
		}

		if req.Command == CHANNEL {
			inputPromptText = fmt.Sprintf("%v >> ", req.Payload)
			go c.tryStartChannelMode(req)
		}

	}

	c.disconnect()
}


func (c *Client) tryStartChannelMode(req request) {
	if req.Command != CHANNEL {
		return
	}
	c.isChannelMode = true

	for c.isChannelMode {
		netData, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			helpers.PrintErrPrefix("CHANNEL", err)
			continue
		}

		var res Response
		err = json.Unmarshal([]byte(netData), &res)
		if err != nil {
			helpers.PrintErrPrefix("UNMARSHAL", err)
			continue
		}

		if res.Code == ERROR {
			return
		}

		fmt.Println("\n", res.Result)
		fmt.Print(inputPromptText)
	}
}

func (c *Client) disconnectFromChannel(jsonReq []byte) {
	inputPromptText = ">> "
	c.send(jsonReq)
	c.isChannelMode = false
}
