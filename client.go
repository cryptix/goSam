package goSam

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	samConn net.Conn

	fromSam *bufio.Reader
	toSam   *bufio.Writer
}

func NewDefaultClient() (*Client, error) {
	return NewClient("localhost:7656")
}

func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	c := &Client{
		samConn: conn,
		fromSam: bufio.NewReader(conn),
		toSam:   bufio.NewWriter(conn),
	}
	return c, nil
}

func (c *Client) Hello() (err error) {
	if _, err = c.toSam.WriteString("HELLO VERSION MIN=3.0 MAX=3.0\n"); err != nil {
		return err
	}

	if err = c.toSam.Flush(); err != nil {
		return err
	}

	for {
		line, err := c.fromSam.ReadString('\n')
		if err != nil {
			return err
		}

		reply, err := parseReply(line)
		if err != nil {
			return err
		}

		if reply.Topic != "HELLO" {
			return fmt.Errorf("Unknown Reply: %+v\n", reply)
		}

		if reply.Pairs["RESULT"] != "OK" {
			return fmt.Errorf("Handshake did not succeed")
		}

		break
	}
	return nil
}

func (c *Client) Close() error {
	return c.samConn.Close()
}
