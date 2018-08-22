package goSam

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"
)

// helper to validate sendCmd inputs
func (c *Client) validCmd(str string, args ...interface{}) (string, error) {
	if s := fmt.Sprintf(str, args...); strings.Contains(s, "\n") {
		sl := strings.Split(s, "\n")
		if len(sl) == 2 {
			if sl[1] != "" {
				return sl[1], fmt.Errorf("Error, there should be no options after the newline")
			}
			for li, in := range sl {
				fmt.Println(li, in)
			}
			return s, nil
		}
		return "", fmt.Errorf("Error, invalid length: %d", len(sl))
	}
	return "", fmt.Errorf("Error, invalid input")
}

func (c *Client) validCreate() (string, error) {
	id := rand.Int31n(math.MaxInt32)
	result, err := c.validCmd("SESSION CREATE STYLE=STREAM ID=%d DESTINATION=%s %s\n", id, "abc.i2p", client.allOptions())
	return result, err
}

func TestOptionAddrString(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("127.0.0.1:7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionAddrStringLh(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("localhost:7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionAddrSlice(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("127.0.0.1", "7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionAddrMixedSlice(t *testing.T) {
	client, err := NewClientFromOptions(SetAddrMixed("127.0.0.1", 7656), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionHost(t *testing.T) {
	client, err := NewClientFromOptions(SetHost("127.0.0.1"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionPort(t *testing.T) {
	client, err := NewClientFromOptions(SetPort("7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionPortInt(t *testing.T) {
	client, err := NewClientFromOptions(SetPortInt(7656), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionDebug(t *testing.T) {
	client, err := NewClientFromOptions(SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionInLength(t *testing.T) {
	client, err := NewClientFromOptions(SetInLength(3), SetDebug(true))
	client.inlength()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionOutLength(t *testing.T) {
	client, err := NewClientFromOptions(SetInLength(3), SetDebug(true))
	client.outlength()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionInVariance(t *testing.T) {
	client, err := NewClientFromOptions(SetInVariance(1), SetDebug(true))
	client.invariance()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionOutVariance(t *testing.T) {
	client, err := NewClientFromOptions(SetOutVariance(1), SetDebug(true))
	client.outvariance()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionInQuantity(t *testing.T) {
	client, err := NewClientFromOptions(SetInQuantity(6), SetDebug(true))
	client.inquantity()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionOutQuantity(t *testing.T) {
	client, err := NewClientFromOptions(SetOutQuantity(6), SetDebug(true))
	client.outquantity()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionInBackups(t *testing.T) {
	client, err := NewClientFromOptions(SetInBackups(5), SetDebug(true))
	client.inbackups()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionOutBackups(t *testing.T) {
	client, err := NewClientFromOptions(SetOutBackups(5), SetDebug(true))
	client.outbackups()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionEncryptLease(t *testing.T) {
	client, err := NewClientFromOptions(SetEncrypt(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionUnpublishedLease(t *testing.T) {
	client, err := NewClientFromOptions(SetUnpublished(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionReduceIdle(t *testing.T) {
	client, err := NewClientFromOptions(SetReduceIdle(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionReduceIdleTime(t *testing.T) {
	client, err := NewClientFromOptions(SetReduceIdleTime(300001), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionReduceIdleCount(t *testing.T) {
	client, err := NewClientFromOptions(SetReduceIdleQuantity(4), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionCloseIdle(t *testing.T) {
	client, err := NewClientFromOptions(SetCloseIdle(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}

func TestOptionCloseIdleTime(t *testing.T) {
	client, err := NewClientFromOptions(SetCloseIdleTime(300001), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
}
