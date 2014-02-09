package goSam

import (
	"fmt"
)

type Result int

const (
	ResultOk             Result = iota //Operation completed successfully
	ResultCantReachPeer                //The peer exists, but cannot be reached
	ResultDuplicatedDest               //The specified Destination is already in use
	ResultI2PError                     //A generic I2P error (e.g. I2CP disconnection, etc.)
	ResultInvalidKey                   //The specified key is not valid (bad format, etc.)
	ResultKeyNotFound                  //The naming system can't resolve the given name
	ResultPeerNotFound                 //The peer cannot be found on the network
	ResultTimeout                      // Timeout while waiting for an event (e.g. peer answer)
)

type ReplyError struct {
	Result Result
	Reply  *Reply
}

func (r ReplyError) Error() string {
	return fmt.Sprintf("ReplyError: Result:%d - Reply:%+v", r.Reply)
}

func (c *Client) Lookup(name string) (addr string, err error) {
	msg := fmt.Sprintf("NAMING LOOKUP NAME=%s\n", name)
	if _, err = c.toSam.WriteString(msg); err != nil {
		return
	}

	if err = c.toSam.Flush(); err != nil {
		return
	}

	var (
		line string
		r    *Reply
	)
	for {
		line, err = c.fromSam.ReadString('\n')
		if err != nil {
			return
		}

		r, err = parseReply(line)
		if err != nil {
			break
		}

		if r.Topic != "NAMING" || r.Type != "REPLY" {
			err = fmt.Errorf("Unknown Reply: %+v\n", r)
			break
		}

		switch r.Pairs["RESULT"] {
		case "OK":
			addr = r.Pairs["VALUE"]
			return
		case "KEY_NOT_FOUND":
			err = ReplyError{ResultKeyNotFound, r}
		}

		if r.Pairs["NAME"] != name {
			err = fmt.Errorf("i2p Replyied with: %+v\n", r)
			break
		}

		break
	}

	return
}
