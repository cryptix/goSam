package goSam

import (
	"fmt"
	"net"
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

	if c.verbose {
		newC.ToggleVerbose()
	}

	resp, err := newC.StreamAccept(id)
	if err != nil {
		return nil, err
	}

	fmt.Println("Accept Resp:", resp)

	return newC.SamConn, nil
}
