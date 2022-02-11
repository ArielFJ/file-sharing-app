package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	conn      net.Conn
	username  string
	closeChan chan bool
}

func NewClient(c net.Conn, channel chan bool) Client {
	return Client{
		conn:      c,
		username:  "anonymous",
		closeChan: channel,
	}
}

// func (c *Client) receive() {

// }

func (c *Client) setUsername(name string) {
	c.username = name
}

func (c *Client) send() {

}

func (c *Client) disconnect() {
	// c.conn.Close()
	c.closeChan <- true
}

// func (c *Client) connect() {

// }

func (c *Client) handleSession() {
	for {
		input, err := takeInput()
		if err != nil {
			continue
		}

		req := buildRequest(input)

		jsonBytes, err := json.Marshal(req)
		if err != nil {
			printErr(err)
			continue
		}

		if req.Command == HELP {
			showHelp()
			continue
		}
		
		c.conn.Write(append(jsonBytes, '\n'))

		// if !expectResponse(msg.Command) {
		// 	continue
		// }

		response, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			printErr(err)
			continue
		}

		fmt.Printf("-> %v", response)

		if req.Command == EXIT {
			break
		}
	}

	c.disconnect()
}

func takeInput() (input string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ")
	input, err = reader.ReadString('\n')
	return
}

func normalizeInput(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

func buildRequest(text string) request {
	cleanText := strings.ReplaceAll(text, "\r\n", "") // Take just the input without the return
	words := strings.Split(cleanText, " ")
	cmd := normalizeInput(words[0])
	payload := strings.Join(words[1:], " ")

	return NewRequest(cmd, payload)
}
