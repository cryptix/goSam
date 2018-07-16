package goSam

import (
	"fmt"
	"testing"
)

func TestClientLookupInvalid(t *testing.T) {
	var err error

	setup(t)
	defer teardown(t)

	addr, err := client.Lookup("abc.i2p")
	if addr != "" || err == nil {
		t.Error("client.Lookup() should throw an error.")
	}

	repErr, ok := err.(ReplyError)
	if !ok {
		t.Fatalf("client.Lookup() should return a ReplyError")
	}
	if repErr.Result != ResultKeyNotFound {
		t.Errorf("client.Lookup() should throw an ResultKeyNotFound error.\nGot:%+v\n", repErr)
	}
}

func ExampleClient_Lookup() {
	client, err := NewDefaultClient()
	if err != nil {
		fmt.Printf("NewDefaultClient() should not throw an error.\n%s\n", err)
		return
	}

	_, err = client.Lookup("zzz.i2p")
	if err != nil {
		fmt.Printf("client.Lookup() should not throw an error.\n%s\n", err)
		return
	}

	fmt.Println("Address of zzz.i2p:")
	// Addresses change all the time
	// fmt.Println(addr)

	// Output:
	//Address of zzz.i2p:
}
