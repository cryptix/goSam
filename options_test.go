package goSam

import (
	"fmt"
	"math"
	"math/rand"
	//"net/http"
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
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionAddrStringLh(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("localhost:7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionAddrSlice(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("127.0.0.1", "7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionAddrMixedSlice(t *testing.T) {
	client, err := NewClientFromOptions(SetAddrMixed("127.0.0.1", 7656), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionHost(t *testing.T) {
	client, err := NewClientFromOptions(SetHost("127.0.0.1"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionPort(t *testing.T) {
	client, err := NewClientFromOptions(SetPort("7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionPortInt(t *testing.T) {
	client, err := NewClientFromOptions(SetPortInt(7656), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionDebug(t *testing.T) {
	client, err := NewClientFromOptions(SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionInLength(t *testing.T) {
	client, err := NewClientFromOptions(SetInLength(3), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	client.inlength()
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionOutLength(t *testing.T) {
	client, err := NewClientFromOptions(SetInLength(3), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	client.outlength()
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionInVariance(t *testing.T) {
	client, err := NewClientFromOptions(SetInVariance(1), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	client.invariance()
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionOutVariance(t *testing.T) {
	client, err := NewClientFromOptions(SetOutVariance(1), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	client.outvariance()
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionInQuantity(t *testing.T) {
	client, err := NewClientFromOptions(SetInQuantity(6), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	client.inquantity()
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionOutQuantity(t *testing.T) {
	client, err := NewClientFromOptions(SetOutQuantity(6), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	client.outquantity()
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionInBackups(t *testing.T) {
	client, err := NewClientFromOptions(SetInBackups(5), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	client.inbackups()
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionOutBackups(t *testing.T) {
	client, err := NewClientFromOptions(SetOutBackups(5), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	client.outbackups()
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionEncryptLease(t *testing.T) {
	client, err := NewClientFromOptions(SetEncrypt(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionUnpublishedLease(t *testing.T) {
	client, err := NewClientFromOptions(SetUnpublished(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionReduceIdle(t *testing.T) {
	client, err := NewClientFromOptions(SetReduceIdle(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionReduceIdleTime(t *testing.T) {
	client, err := NewClientFromOptions(SetReduceIdleTime(300001), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionReduceIdleCount(t *testing.T) {
	client, err := NewClientFromOptions(SetReduceIdleQuantity(4), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}

}

func TestOptionCloseIdle(t *testing.T) {
	client, err := NewClientFromOptions(SetCloseIdle(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionCloseIdleTime(t *testing.T) {
	client, err := NewClientFromOptions(SetCloseIdleTime(300001), SetDebug(true))
	if err != nil {
		t.Fatalf("NewClientFromOptions() Error: %q\n", err)
	}
	if result, err := client.validCreate(); err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Log(result)
	}
	client.CreateStreamSession("")
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}
