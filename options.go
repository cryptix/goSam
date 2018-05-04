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
			return nil
		} else {
			return fmt.Errorf("Invalid outbound tunnel length")
		}
	}
}

func SetInVariance(i int) func(*Client) error {
	return func(c *Client) error {
		if i < 7 && i > -7 {
			c.inVariance = i
			return nil
		} else {
			return fmt.Errorf("Invalid inbound tunnel length")
		}
	}
}

func SetOutVariance(i int) func(*Client) error {
	return func(c *Client) error {
		if i < 7 && i > -7 {
			c.outVariance = i
			return nil
		} else {
			return fmt.Errorf("Invalid outbound tunnel variance")
		}
	}
}

func SetInQuantity(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 16 {
			c.inQuantity = u
			return nil
		} else {
			return fmt.Errorf("Invalid inbound tunnel quantity")
		}
	}
}

func SetOutQuantity(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 16 {
			c.outQuantity = u
			return nil
		} else {
			return fmt.Errorf("Invalid outbound tunnel quantity")
		}
	}
}

func SetInBackups(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 6 {
			c.inBackups = u
			return nil
		} else {
			return fmt.Errorf("Invalid inbound tunnel backup quantity")
		}
	}
}

func SetOutBackups(u uint) func(*Client) error {
	return func(c *Client) error {
		if u < 6 {
			c.outBackups = u
			return nil
		} else {
			return fmt.Errorf("Invalid outbound tunnel backup quantity")
		}
	}
}

func SetUnpublished(b bool) func(*Client) error {
	return func(c *Client) error {
		c.dontPublishLease = b
		return nil
	}
}

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
  }else{
      return "i2cp.encryptLeaseSet=false"
  }
}

func (c *Client) dontpublishlease() string {
  if c.dontPublishLease {
      return "i2cp.dontPublishLeaseSet=true"
  }else{
      return "i2cp.dontPublishLeaseSet=false"
  }
}



//return all options as string array ready for passing to sendcmd
func (c *Client) allOptions() []string {
	var options []string
	options = append(options, c.inlength(), c.outlength(), c.invariance(), c.outvariance(), c.inquantity(), c.outquantity(), c.inbackups(), c.outbackups(), c.encryptlease(), c.dontpublishlease())
	return options
}
