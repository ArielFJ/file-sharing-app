package models

import "net"

type Client struct {
	conn net.Conn
	username string
}

func NewClient(c net.Conn) Client {
	return Client{
		conn: c,
		username: "anonymous",
	};
}

// func (c *Client) receive() {

// }

func (c *Client) send() {
	
}

func (c *Client) disconnect() {
	c.conn.Close()
}

// func (c *Client) connect() {
	
// }