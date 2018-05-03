package goSam

import (
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
			return err
		}
		if port < 65536 && port > -1 {
			c.port = s
		} else {
			c.port = "7656"
		}
		return nil
	}
}
