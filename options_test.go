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
	client, err := NewClientFromOptions(
		SetHost("127.0.0.1"),
		SetPort("7656"),
		SetInLength(3),
		SetOutLength(3),
		SetInVariance(1),
		SetOutVariance(1),
		SetInQuantity(6),
		SetOutQuantity(6),
		SetInBackups(2),
		SetOutBackups(2),
		SetEncrypt(true),
		SetDebug(true),
		SetUnpublished(true),
		SetReduceIdle(true),
		SetReduceIdleTime(300001),
		SetReduceIdleQuantity(4),
		SetCloseIdle(true),
		SetCloseIdleTime(300001),
	)
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
	client, err := NewClientFromOptions(
		SetHost("127.0.0.1"),
		SetPortInt(7656),
		SetInLength(3),
		SetOutLength(3),
		SetInVariance(1),
		SetOutVariance(1),
		SetInQuantity(6),
		SetOutQuantity(6),
		SetInBackups(2),
		SetOutBackups(2),
		SetEncrypt(true),
		SetDebug(true),
		SetUnpublished(true),
		SetReduceIdle(true),
		SetReduceIdleTime(300001),
		SetReduceIdleQuantity(4),
		SetCloseIdle(true),
		SetCloseIdleTime(300001),
	)
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
