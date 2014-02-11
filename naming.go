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
		err = fmt.Errorf("i2p Replyed to another name.\nWanted:%s\nGot: %+v\n", name, r)
		return
	}

	addr = r.Pairs["VALUE"]
	return
}
