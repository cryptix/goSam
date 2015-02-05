package goSam

import (
	"fmt"
	"net"

	"github.com/cryptix/go/debug"
)

// Accept creates a new Client and accepts a connection on it
func (c *Client) Accept() (net.Conn, error) {
	id, newAddr, err := c.CreateStreamSession("")
	if err != nil {
		return nil, err
	}

	fmt.Println("NewAddr:", newAddr)

	newC, err := NewDefaultClient()
	if err != nil {
		return nil, err
	}

	if c.debug {
		newC.SamConn = debug.WrapConn(newC.SamConn)
	}

	resp, err := newC.StreamAccept(id)
	if err != nil {
		return nil, err
	}

	fmt.Println("Accept Resp:", resp)

	return newC.SamConn, nil
}
