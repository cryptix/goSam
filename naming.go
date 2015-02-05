package goSam

import (
	"fmt"
	"os"
)

// Lookup askes SAM for the internal i2p address from name
func (c *Client) Lookup(name string) (addr string, err error) {
	var r *Reply

	r, err = c.sendCmd(fmt.Sprintf("NAMING LOOKUP NAME=%s", name))
	if err != nil {
		return
	}

	if r.Topic != "NAMING" || r.Type != "REPLY" {
		err = fmt.Errorf("Unknown Reply: %+v\n", r)
		return
	}

	result := r.Pairs["RESULT"]
	if result != "OK" {
		err = ReplyError{result, r}
		return
	}

	if r.Pairs["NAME"] != name {
		// somehow different on i2pd
		if r.Pairs["NAME"] != "ME" {
			err = fmt.Errorf("Lookup() Replyed to another name.\nWanted:%s\nGot: %+v\n", name, r)
			return
		}
		fmt.Fprintln(os.Stderr, "WARNING: Lookup() Replyed to another name. assuming i2pd c++ fluke")
	}

	addr = r.Pairs["VALUE"]
	return
}
