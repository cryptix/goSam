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
	return func(c *Client) error {
		if u < 7 {
			c.inLength = u
			if c.inLength < c.inVariance {
				return fmt.Errorf("Inbound variance > Inbound length")
			}
			return nil
		} else {
			return fmt.Errorf("Invalid inbound tunnel length")
		}
	}
}

func SetOutLength(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 7 {
			c.outLength = u
			if c.outLength < c.outVariance {
				return fmt.Errorf("Outbound variance > Outbound length")
			}
			return nil
		} else {
			return fmt.Errorf("Invalid outbound tunnel length")
		}
	}
}

func SetInVariance(u uint) func(*Client) error {
	return func(c *Client) error {
		if u <= 3 {
			c.inVariance = u
			if c.inLength < c.inVariance {
				return fmt.Errorf("Inbound variance > Inbound length")
			}
			return nil
		} else {
			return fmt.Errorf("Invalid inbound tunnel length")
		}
	}
}

func SetOutVariance(u uint) func(*Client) error {
	return func(c *Client) error {
		if u <= 3 {
			c.outVariance = u
			if c.outLength < c.outVariance {
				return fmt.Errorf("Outbound variance > Outbound length")
			}
			return nil
		} else {
			return fmt.Errorf("Invalid outbound tunnel length")
		}
	}
}

func SetInQuantity(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 7 {
			c.inQuantity = u
			return nil
		} else {
			return fmt.Errorf("Invalid inbound tunnel length")
		}
	}
}

func SetOutQuantity(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 7 {
			c.outQuantity = u
			return nil
		} else {
			return fmt.Errorf("Invalid outbound tunnel length")
		}
	}
}

//return the inbound length as a string.
func (c *Client) inlength() string {
	return "inbound.length=" + fmt.Sprint(c.inLength)
}

//return the outbound length as a string.
func (c *Client) outlength() string {
	return "outbound.length=" + fmt.Sprint(c.outLength)
}

//return the inbound length variance as a string.
func (c *Client) invariance() string {
	return "inbound.lengthVariance=" + fmt.Sprint(c.inVariance)
}

//return the outbound length variance as a string.
func (c *Client) outvariance() string {
	return "outbound.lengthVariance=" + fmt.Sprint(c.outVariance)
}

//return the inbound tunnel quantity as a string.
func (c *Client) inquantity() string {
	return "inbound.quantity=" + fmt.Sprint(c.inQuantity)
}

//return the outbound tunnel quantity as a string.
func (c *Client) outquantity() string {
	return "outbound.quantity=" + fmt.Sprint(c.outQuantity)
}

//return all options as string array ready for passing to sendcmd
func (c *Client) allOptions() []string {
	var options []string
	options = append(options, c.inlength(), c.outlength(), c.invariance(), c.outvariance(), c.inquantity(), c.outquantity())
	return options
}
