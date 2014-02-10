package goSam

import (
	"fmt"
)

const (
	ResultOk             = "OK"              //Operation completed successfully
	ResultCantReachPeer  = "CANT_REACH_PEER" //The peer exists, but cannot be reached
	ResultDuplicatedId   = "DUPLICATED_ID"   //If the nickname is already associated with a session :
	ResultDuplicatedDest = "DUPLICATED_DEST" //The specified Destination is already in use
	ResultI2PError       = "I2P_ERROR"       //A generic I2P error (e.g. I2CP disconnection, etc.)
	ResultInvalidKey     = "INVALID_KEY"     //The specified key is not valid (bad format, etc.)
	ResultKeyNotFound    = "KEY_NOT_FOUND"   //The naming system can't resolve the given name
	ResultPeerNotFound   = "PEER_NOT_FOUND"  //The peer cannot be found on the network
	ResultTimeout        = "TIMEOUT"         // Timeout while waiting for an event (e.g. peer answer)
)

type ReplyError struct {
	Result string
	Reply  *Reply
}

func (r ReplyError) Error() string {
	return fmt.Sprintf("ReplyError: Result:%s - Reply:%+v", r.Result, r.Reply)
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

	line, err = c.fromSam.ReadString('\n')
	if err != nil {
		return
	}

	r, err = parseReply(line)
	if err != nil {
		return
	}

	if r.Topic != "NAMING" || r.Type != "REPLY" {
		err = fmt.Errorf("Unknown Reply: %+v\n", r)
		return
	}

	switch r.Pairs["RESULT"] {
	case "OK":
		addr = r.Pairs["VALUE"]
		return
	case "KEY_NOT_FOUND":
		err = ReplyError{ResultKeyNotFound, r}
		return
	}

	if r.Pairs["NAME"] != name {
		err = fmt.Errorf("i2p Replyied with: %+v\n", r)
	}

	return
}
