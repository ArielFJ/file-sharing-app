package models

import "net"

type Client struct {
	conn           net.Conn
	username       string
	currentChannel string
}

func NewClient(c net.Conn) Client {
	return Client{
		conn:     c,
		username: "anonymous",
	}
}

func (c *Client) send(data []byte) {
	c.conn.Write(append(data, '\n'))
}

func (c *Client) disconnect() {
	c.conn.Close()
}
