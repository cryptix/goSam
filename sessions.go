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

	var r *Reply

	id = rand.Int31n(math.MaxInt32)
	createCmd := fmt.Sprintf("SESSION CREATE STYLE=STREAM ID=%d DESTINATION=%s\n", id, dest)
	r, err = c.sendCmd(createCmd)
	if err != nil {
		return
	}

	if r.Topic != "SESSION" || r.Type != "STATUS" {
		err = fmt.Errorf("Unknown Reply: %+v\n", r)
		return
	}

	result := r.Pairs["RESULT"]
	if result != "OK" {
		err = ReplyError{ResultKeyNotFound, r}
		return
	}

	fmt.Println("createStreamSession created")
	newDest = r.Pairs["DESTINATION"]

	return
}
