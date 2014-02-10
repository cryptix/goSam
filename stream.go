package goSam

import (
	"fmt"
)

func (c *Client) StreamConnect(id int32, dest string) error {
	connectCmd := fmt.Sprintf("STREAM CONNECT ID=%d DESTINATION=%s\n", id, dest)
	_, err := c.toSam.WriteString(connectCmd)
	if err != nil {
		return err
	}

	if err = c.toSam.Flush(); err != nil {
		return err
	}

	var (
		line string
		r    *Reply
	)

	line, err = c.fromSam.ReadString('\n')
	if err != nil {
		return err
	}

	r, err = parseReply(line)
	if err != nil {
		return err
	}

	if r.Topic != "STREAM" || r.Type != "STATUS" {
		return fmt.Errorf("Unknown Reply: %+v\n", r)
	}

	switch r.Pairs["RESULT"] {
	case ResultOk:
		fmt.Println("StreamConnect OK")
		return nil
	case ResultDuplicatedId:
		return ReplyError{ResultDuplicatedId, r}
	case ResultDuplicatedDest:
		return ReplyError{ResultDuplicatedDest, r}
	case ResultInvalidKey:
		return ReplyError{ResultInvalidKey, r}
	case ResultI2PError:
		return ReplyError{ResultKeyNotFound, r}
	}

	return nil
}
