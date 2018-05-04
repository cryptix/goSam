package goSam

import (
	"fmt"
	"strconv"
	"strings"
)

type Option func(*Client) error

func SetAddr(s ...interface{}) func(*Client) error {
	return func(c *Client) error {
		if len(s) == 1 {
			switch v := s[0].(type) {
			case string:
				split := strings.SplitN(v, ":", 2)
				if len(split) == 2 {
					if i, err := strconv.Atoi(split[1]); err == nil {
                        if i < 65536 {
    						c.addr = split[0]
                            c.port = split[1]
                        }else{
                            return fmt.Errorf("Invalid port")
                        }
					} else {
						return fmt.Errorf("Invalid port; non-number")
					}
				} else {
					return fmt.Errorf("Invalid address; use host:port", split)
				}
			default:
				return fmt.Errorf("Invalid address; address must be string")
			}
		} else if len(s) == 2 {
			switch v := s[1].(type) {
			case int:
                if v < 65536 {
                    c.addr = s[0].(string)
                    c.port = strconv.Itoa(v)
                }else{
                    return fmt.Errorf("Invalid port")
                }
			case string:
				if i, err := strconv.Atoi(s[1].(string)); err == nil {
                    if i < 65536 {
                        c.addr = s[0].(string)
                        c.port = s[1].(string)
                    }else{
                        return fmt.Errorf("Invalid port")
                    }
				} else {
					return fmt.Errorf("Invalid port; non-number")
				}
			default:
				return fmt.Errorf("Invalid port; non-number")
			}
		} else {
			return fmt.Errorf("Invalid address")
		}
		return nil
	}
}

func SetHost(s string) func(*Client) error {
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
