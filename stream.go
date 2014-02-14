package goSam

import (
	"fmt"
)

// StreamConnect asks SAM for a TCP-Like connection to dest, has to be called on a new Client
func (c *Client) StreamConnect(id int32, dest string) (err error) {
	var r *Reply

	r, err = c.sendCmd(fmt.Sprintf("STREAM CONNECT ID=%d DESTINATION=%s", id, dest))
	if err != nil {
		return err
	}

	if r.Topic != "STREAM" || r.Type != "STATUS" {
		return fmt.Errorf("Unknown Reply: %+v\n", r)
	}

	result := r.Pairs["RESULT"]
	if result != "OK" {
		return ReplyError{result, r}
	}

	return nil
}

// StreamAccept asks SAM to accept a TCP-Like connection
func (c *Client) StreamAccept(id int32) (r *Reply, err error) {

	r, err = c.sendCmd(fmt.Sprintf("STREAM ACCEPT ID=%d SILENT=false", id))
	if err != nil {
		return nil, err
	}

	if r.Topic != "STREAM" || r.Type != "STATUS" {
		return nil, fmt.Errorf("Unknown Reply: %+v\n", r)
	}

	result := r.Pairs["RESULT"]
	if result != "OK" {
		return nil, ReplyError{result, r}
	}

	return r, nil
}
