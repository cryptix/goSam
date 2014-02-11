package goSam

import (
	"fmt"
	"net"
	"strings"
)

// implements the net.Dial function to be used as http.Transport
func (c *Client) Dial(network, addr string) (net.Conn, error) {
	portIdx := strings.Index(addr, ":")
	if portIdx >= 0 {
		addr = addr[:portIdx]
	}
	addr, err := c.Lookup(addr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Dial Lookup:", addr)

	id, _, err := c.createStreamSession("")
	if err != nil {
		return nil, err
	}

	newC, err := NewDefaultClient()
	if err != nil {
		return nil, err
	}

	fmt.Println("newC Hello OK")

	if newC.StreamConnect(id, addr) != nil {
		return nil, err
	}

	fmt.Println("StreamConnect OK")

	return newC.samConn, nil
}
