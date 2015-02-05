package goSam

import "testing"

var client *Client

func setup(t *testing.T) {
	var err error

	// these tests expect a running SAM brige on this address
	client, err = NewClientFromOptions(SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}

}

func teardown(t *testing.T) {
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestClientHello(t *testing.T) {
	setup(t)
	teardown(t)
}
