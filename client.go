package goSam

import (
	"bufio"
	"fmt"
	"io"
)

type Client struct {
	samConn io.ReadWriteCloser

	fromSam *bufio.Reader
	toSam   *bufio.Writer
}

func NewClient(samConn io.ReadWriteCloser) (*Client, error) {
	c := &Client{
		samConn: samConn,
		fromSam: bufio.NewReader(samConn),
		toSam:   bufio.NewWriter(samConn),
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
