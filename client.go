package goSam

import (
	"bufio"
	"fmt"
	"net"

	"github.com/cryptix/go/debug"
)

// A Client represents a single Connection to the SAM bridge
type Client struct {
	addr string
	port string

	SamConn net.Conn
	rd      *bufio.Reader

	inLength uint
    inVariance int

	outLength uint
    outVariance int

	debug bool
}

// NewDefaultClient creates a new client, connecting to the default host:port at localhost:7656
func NewDefaultClient() (*Client, error) {
	return NewClient("localhost:7656")
}

// NewClient creates a new client, connecting to a specified port
func NewClient(addr string) (*Client, error) {
	return NewClientFromOptions(SetAddr(addr))
}

// NewClientFromOptionss creates a new client, connecting to a specified port
func NewClientFromOptions(opts ...func(*Client) error) (*Client, error) {
	var c Client
	c.addr = "127.0.0.1"
	c.port = "7656"
    c.inLength = 3
    c.inVariance = 0
	c.outLength = 3
	c.outVariance = 0
	c.debug = false
	for _, o := range opts {
		if err := o(&c); err != nil {
			return nil, err
		}
	}
	conn, err := net.Dial("tcp", c.samaddr())
	if err != nil {
		return nil, err
	}
	if c.debug {
		conn = debug.WrapConn(conn)
	}
	c.SamConn = conn
	c.rd = bufio.NewReader(conn)
	return &c, c.hello()
}

//return the combined addr:port of the SAM bridge
func (c *Client) samaddr() string {
	return fmt.Sprintf("%s:%s", c.addr, c.port)
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
