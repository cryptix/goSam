package goSam

import (
	"fmt"
	"strconv"
)

type Option func(*Client) error

func SetAddr(s string) func(*Client) error {
	return func(c *Client) error {
		c.addr = s
		return nil
	}
}

func SetPort(s string) func(*Client) error {
	return func(c *Client) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid port; non-number")
		}
		if port < 65536 && port > -1 {
			c.port = s
			return nil
		} else {
			return fmt.Errorf("Invalid port")
		}
	}
}

func SetDebug(b bool) func(*Client) error {
	return func(c *Client) error {
		c.debug = b
		return nil
	}
}

func SetInLength(u uint) func(*Client) error {
    return func(c *Client) error{
        if u < 7 {
            c.inLength = u
        }
    }
}

func SetOutLength(u uint) func(*Client) error {
    return func(c *Client) error{
        if u < 7 {
            c.outLength = u
        }
    }
}
