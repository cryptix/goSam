package goSam

import (
	"net"
	"testing"
)

var (
	client *Client
)

func setup() {
	var err error

	// these tests expect a running SAM brige on this address
	conn, err := net.Dial("tcp", "localhost:7656")
	if err != nil {
		panic(err)
	}

	client, err = NewClient(conn)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	client.Close()
}

func TestClientHello(t *testing.T) {
	var err error

	setup()
	defer teardown()

	err = client.Hello()
	if err != nil {
		t.Errorf("client.Hello() should not throw an error.\n%s\n", err)
	}
}
