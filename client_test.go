package goSam

import "testing"

var (
	client *Client
)

func setup() {
	var err error

	// these tests expect a running SAM brige on this address
	client, err = NewDefaultClient()
	if err != nil {
		panic(err)
	}
}

func teardown() {
	client.Close()
}

func TestClientHello(t *testing.T) {
	var err error

	client, err = NewDefaultClient()
	if err != nil {
		t.Errorf("client.Hello() should not throw an error.\n%s\n", err)
	}

	client.Close()
}
