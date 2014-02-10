package goSam

import (
	"fmt"
	"math"
	"math/rand"
)

func (c *Client) createStreamSession(dest string) (id int32, newDest string, err error) {
	if dest == "" {
		dest = "TRANSIENT"
	}

	id = rand.Int31n(math.MaxInt32)
	createCmd := fmt.Sprintf("SESSION CREATE STYLE=STREAM ID=%d DESTINATION=%s\n", id, dest)
	_, err = c.toSam.WriteString(createCmd)
	if err != nil {
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
	fmt.Println("createStreamSession line:", line)

	r, err = parseReply(line)
	if err != nil {
		return
	}

	if r.Topic != "SESSION" || r.Type != "STATUS" {
		err = fmt.Errorf("Unknown Reply: %+v\n", r)
		return
	}

	switch r.Pairs["RESULT"] {
	case ResultOk:
		fmt.Println("createStreamSession created")
		newDest = r.Pairs["DESTINATION"]
		return
	case ResultDuplicatedId:
		err = ReplyError{ResultDuplicatedId, r}
		return
	case ResultDuplicatedDest:
		err = ReplyError{ResultDuplicatedDest, r}
		return
	case ResultInvalidKey:
		err = ReplyError{ResultInvalidKey, r}
		return
	case ResultI2PError:
		err = ReplyError{ResultKeyNotFound, r}
		return
	}

	return
}
