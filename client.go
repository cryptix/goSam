package goSam

import (
	"bufio"
	"fmt"
	"net"

	"github.com/cryptix/go/debug"
)

var ConnDebug = false

// A Client represents a single Connection to the SAM bridge
type Client struct {
	SamConn net.Conn
	rd      *bufio.Reader
}

// NewDefaultClient creates a new client, connecting to the default host:port at localhost:7656
func NewDefaultClient() (*Client, error) {
	return NewClient("localhost:7656")
}

// NewClient creates a new client, connecting to a specified port
func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	if ConnDebug {
		conn = debug.WrapConn(conn)
	}
	c := &Client{
		SamConn: conn,
		rd:      bufio.NewReader(conn),
	}
	return c, c.hello()
}

// send the initial handshake command and check that the reply is ok
func (c *Client) hello() (err error) {
	var r *Reply

	r, err = c.sendCmd("HELLO VERSION MIN=3.0 MAX=3.0")
	if err != nil {
		return err
	}

	if r.Topic != "HELLO" {
		return fmt.Errorf("Unknown Reply: %+v\n", r)
	}

	if r.Pairs["RESULT"] != "OK" || r.Pairs["VERSION"] != "3.0" {
		return fmt.Errorf("Handshake did not succeed\nReply:%+v\n", r)
	}

	return nil
}

// helper to send one command and parse the reply by sam
func (c *Client) sendCmd(cmd string) (r *Reply, err error) {
	if _, err = fmt.Fprintln(c.SamConn, cmd); err != nil {
		return
	}

	line, err := c.rd.ReadString('\n')
	if err != nil {
		return
	}

	return parseReply(line)
}

// Close the underlying socket to SAM
func (c *Client) Close() error {
	return c.SamConn.Close()
}
