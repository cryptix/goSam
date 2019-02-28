package goSam

import (
	"context"
	"fmt"
	"net"
	"strings"
)

// DialContext implements the net.DialContext function and can be used for http.Transport
func (c *Client) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	errCh := make(chan error, 1)
	connCh := make(chan net.Conn, 1)
	go func() {
		if conn, err := c.Dial(network, addr); err != nil {
			errCh <- err
		} else if ctx.Err() != nil {
			conn.Close()
		} else {
			connCh <- conn
		}
	}()
	select {
	case err := <-errCh:
		return nil, err
	case conn := <-connCh:
		return conn, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c Client) dialCheck(addr string) (int32, bool) {
	fmt.Println("DIAL CHECK")
	if c.lastaddr != addr {
		fmt.Println("Preparing to dial new address.")
		return c.NewID(), true
	}
	fmt.Println("No new address to dial.")
	return c.NewID(), false
}

// Dial implements the net.Dial function and can be used for http.Transport
func (c *Client) Dial(network, addr string) (net.Conn, error) {
	portIdx := strings.Index(addr, ":")
	if portIdx >= 0 {
		addr = addr[:portIdx]
	}
	addr, err := c.Lookup(addr)
	if err != nil {
		return nil, err
	}

	var id int32
	var test bool
	if id, test = c.dialCheck(addr); test == true {
		c.destination, err = c.CreateStreamSession(id, "")
		if err != nil {
			return nil, err
		}
		c.lastaddr = addr
		newC, err := c.NewClient()
		if err != nil {
			return nil, err
		}

		err = newC.StreamConnect(id, addr)
		if err != nil {
			return nil, err
		}
		return newC.SamConn, nil
	}
	c, err = c.NewClient()
	if err != nil {
		return nil, err
	}

	err = c.StreamConnect(c.id, addr)
	if err != nil {
		return nil, err
	}
	return c.SamConn, nil
}
