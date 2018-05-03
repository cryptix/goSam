package goSam

import (
	"bufio"
	"fmt"
	"net"

	"github.com/cryptix/go/debug"
)

// ConnDebug if set to true, Sam connections are wrapped with logging
var ConnDebug = false

// SamHostAddress determines the address of the SAM bridge used in the Dial
//function
var SamHostAddress = "127.0.0.1"

// SamHostPort determines the port of the SAM bridge used in the Dial function.
var SamHostPort = "7656"

// A Client represents a single Connection to the SAM bridge
type Client struct {
	SamConn net.Conn
	rd      *bufio.Reader
}

// NewDefaultClient creates a new client, connecting to the default host:port at localhost:7656
func NewDefaultClient() (*Client, error) {
	return NewClient(SamHostAddress + ":" + SamHostPort)
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
func (c *Client) hello() error {
	r, err := c.sendCmd("HELLO VERSION MIN=3.0 MAX=3.0\n")
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
func (c *Client) sendCmd(str string, args ...interface{}) (*Reply, error) {
	if _, err := fmt.Fprintf(c.SamConn, str, args...); err != nil {
		return nil, err
	}

	line, err := c.rd.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return parseReply(line)
}

// Close the underlying socket to SAM
func (c *Client) Close() error {
	c.rd = nil
	return c.SamConn.Close()
}
