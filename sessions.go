package goSam

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// CreateStreamSession creates a new STREAM Session.
// Returns the Id for the new Client.
func (c *Client) CreateStreamSession(dest string) (int32, string, error) {
	if dest == "" {
		dest = "TRANSIENT"
	}

	id := rand.Int31n(math.MaxInt32)
	r, err := c.sendCmd("SESSION CREATE STYLE=STREAM ID=%d DESTINATION=%s\n", id, dest, c.allOptions())
	if err != nil {
		return -1, "", err
	}

	// TODO: move check into sendCmd()
	if r.Topic != "SESSION" || r.Type != "STATUS" {
		return -1, "", fmt.Errorf("Unknown Reply: %+v\n", r)
	}

	result := r.Pairs["RESULT"]
	if result != "OK" {
		return -1, "", ReplyError{ResultKeyNotFound, r}
	}

	return id, r.Pairs["DESTINATION"], nil
}
