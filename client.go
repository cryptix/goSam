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

// create a new client, connecting to the default host:port at localhost:7656
func NewDefaultClient() (*Client, error) {
	return NewClient("localhost:7656")
}

// create a new client, connecting to a specified port
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
	return c, c.hello()
}

// send the initial handshake command and check that the reply is ok
func (c *Client) hello() (err error) {
	const hello = "HELLO VERSION MIN=3.0 MAX=3.0\n"
	var r *Reply

	r, err = c.sendCmd(hello)
	if err != nil {
		return err
	}

	if r.Topic != "HELLO" {
		return fmt.Errorf("Unknown Reply: %+v\n", r)
	}

	if r.Pairs["RESULT"] != "OK" || r.Pairs["VERSION"] != "3.0\n" {
		return fmt.Errorf("Handshake did not succeed\nReply:%+v\n", r)
	}

	return nil
}

// helper to send one command and parse the reply by sam
func (c *Client) sendCmd(cmd string) (r *Reply, err error) {
	if _, err = c.toSam.WriteString(cmd); err != nil {
		return
	}

	if err = c.toSam.Flush(); err != nil {
		return
	}

	line, err := c.fromSam.ReadString('\n')
	if err != nil {
		return
	}

	r, err = parseReply(line)
	return
}

func (c *Client) Close() error {
	return c.samConn.Close()
}
