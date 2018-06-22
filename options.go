package goSam

import (
	"fmt"
	"strconv"
	"strings"
)

//Option is a client Option
type Option func(*Client) error

//SetAddr sets a clients's address in the form host:port or host, port
func SetAddr(s ...string) func(*Client) error {
	return func(c *Client) error {
		if len(s) == 1 {
			split := strings.SplitN(s[0], ":", 2)
			if len(split) == 2 {
				if i, err := strconv.Atoi(split[1]); err == nil {
					if i < 65536 {
						c.addr = split[0]
						c.port = split[1]
						return nil
					}
					return fmt.Errorf("Invalid port")
				}
				return fmt.Errorf("Invalid port; non-number")
			}
			return fmt.Errorf("Invalid address; use host:port %s", split)
		} else if len(s) == 2 {
			if i, err := strconv.Atoi(s[1]); err == nil {
				if i < 65536 {
					c.addr = s[0]
					c.port = s[1]
					return nil
				}
				return fmt.Errorf("Invalid port")
			}
			return fmt.Errorf("Invalid port; non-number")
		} else {
			return fmt.Errorf("Invalid address")
		}
	}
}

//SetAddrMixed sets a clients's address in the form host:port or host, port
func SetAddrMixed(s string, i int) func(*Client) error {
	return func(c *Client) error {
		if i, err := strconv.Atoi(s); err == nil {
			if i < 65536 {
				c.addr = s
				c.port = strconv.Itoa(i)
				return nil
			}
			return fmt.Errorf("Invalid port")
		}
		return fmt.Errorf("Invalid port; non-number")
	}
}

//SetHost sets the host of the client's SAM bridge
func SetHost(s string) func(*Client) error {
	return func(c *Client) error {
		c.addr = s
		return nil
	}
}

//SetPort sets the port of the client's SAM bridge using a string
func SetPort(s string) func(*Client) error {
	return func(c *Client) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid port; non-number")
		}
		if port < 65536 && port > -1 {
			c.port = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetPortInt sets the port of the client's SAM bridge using a string
func SetPortInt(i int) func(*Client) error {
	return func(c *Client) error {
		if i < 65536 && i > -1 {
			c.port = strconv.Itoa(i)
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetDebug enables debugging messages
func SetDebug(b bool) func(*Client) error {
	return func(c *Client) error {
		c.debug = b
		return nil
	}
}

//SetInLength sets the number of hops inbound
func SetInLength(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 7 {
			c.inLength = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutLength sets the number of hops outbound
func SetOutLength(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 7 {
			c.outLength = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetInVariance sets the variance of a number of hops inbound
func SetInVariance(i int) func(*Client) error {
	return func(c *Client) error {
		if i < 7 && i > -7 {
			c.inVariance = i
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutVariance sets the variance of a number of hops outbound
func SetOutVariance(i int) func(*Client) error {
	return func(c *Client) error {
		if i < 7 && i > -7 {
			c.outVariance = i
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetInQuantity sets the inbound tunnel quantity
func SetInQuantity(u uint) func(*Client) error {
	return func(c *Client) error {
		if u <= 16 {
			c.inQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetOutQuantity sets the outbound tunnel quantity
func SetOutQuantity(u uint) func(*Client) error {
	return func(c *Client) error {
		if u <= 16 {
			c.outQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetInBackups sets the inbound tunnel backups
func SetInBackups(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 6 {
			c.inBackups = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetOutBackups sets the inbound tunnel backups
func SetOutBackups(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 6 {
			c.outBackups = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetUnpublished tells the router to not publish the client leaseset
func SetUnpublished(b bool) func(*Client) error {
	return func(c *Client) error {
		c.dontPublishLease = b
		return nil
	}
}

//SetEncrypt tells the router to use an encrypted leaseset
func SetEncrypt(b bool) func(*Client) error {
	return func(c *Client) error {
		c.encryptLease = b
		return nil
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

//return the inbound tunnel quantity as a string.
func (c *Client) inbackups() string {
	return "inbound.backupQuantity=" + fmt.Sprint(c.inQuantity)
}

//return the outbound tunnel quantity as a string.
func (c *Client) outbackups() string {
	return "outbound.backupQuantity=" + fmt.Sprint(c.outQuantity)
}

func (c *Client) encryptlease() string {
	if c.encryptLease {
		return "i2cp.encryptLeaseSet=true"
	}
	return "i2cp.encryptLeaseSet=false"
}

func (c *Client) dontpublishlease() string {
	if c.dontPublishLease {
		return "i2cp.dontPublishLeaseSet=true"
	}
	return "i2cp.dontPublishLeaseSet=false"
}

//return all options as string array ready for passing to sendcmd
func (c *Client) allOptions() []string {
	var options []string
	options = append(options, c.inlength())
	options = append(options, c.outlength())
	options = append(options, c.invariance())
	options = append(options, c.outvariance())
	options = append(options, c.inquantity())
	options = append(options, c.outquantity())
	options = append(options, c.inbackups())
	options = append(options, c.outbackups())
	options = append(options, c.dontpublishlease())
	options = append(options, c.encryptlease())
	return options
}
