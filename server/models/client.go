package models

import (
	"fmt"
	"net"
)

type Client struct {
	Conn           net.Conn
	Username       string
	CurrentChannel string
}

func NewClient(c net.Conn) Client {
	return Client{
		Conn:     c,
		Username: "anonymous",
	}
}

func (c *Client) Send(data []byte) {
	c.Conn.Write(append(data, '\n'))
}

func (c *Client) Disconnect() {
	res := NewResponse(OK, EXIT, "You have been disconnected")
	c.Send(res.ToBuffer())
	c.Conn.Close()
}

func (c *Client) GetIdentifier() string {
	return fmt.Sprintf("%v[%v]", c.Username, c.Conn.RemoteAddr())
}
