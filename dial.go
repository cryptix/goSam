package goSam

import (
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

	id, _, err := c.createStreamSession("")
	if err != nil {
		return nil, err
	}

	newC, err := NewDefaultClient()
	if err != nil {
		return nil, err
	}

	if c.verbose {
		newC.ToggleVerbose()
	}

	err = newC.StreamConnect(id, addr)
	if err != nil {
		return nil, err
	}

	return newC.SamConn, nil
}
