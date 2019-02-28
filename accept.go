package goSam

import (
	"fmt"
	"net"

	"github.com/cryptix/go/debug"
)

// Accept creates a new Client and accepts a connection on it
func (c *Client) Accept() (net.Conn, error) {
	var err error
	var id int32
	id = c.NewID()
	c.destination, err = c.CreateStreamSession(id, "")
	if err != nil {
		return nil, err
	}

	fmt.Println("destination:", c.destination)

	c, err = c.NewClient()
	if err != nil {
		return nil, err
	}

	if c.debug {
		c.SamConn = debug.WrapConn(c.SamConn)
	}

	resp, err := c.StreamAccept(id)
	if err != nil {
		return nil, err
	}

	fmt.Println("Accept Resp:", resp)

	return c.SamConn, nil
}
