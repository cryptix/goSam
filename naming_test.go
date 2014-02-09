package goSam

import (
	"fmt"
	"testing"
)

func TestClientLookupInvalid(t *testing.T) {
	var err error

	setup()
	defer teardown()

	client.Hello()

	addr, err := client.Lookup("abci2p")
	if addr != "" || err == nil {
		t.Error("client.Lookup() should throw an error.")
	}

	repErr, ok := err.(ReplyError)
	if ok && repErr.Result != ResultKeyNotFound {
		t.Error("client.Lookup() should throw an ResultKeyNotFound error. Got:%v\n", repErr)
	}
}

func ExampleClient_Lookup() {
	var err error

	setup()
	defer teardown()

	client.Hello()

	_, err = client.Lookup("zzz.i2p")
	if err != nil {
		fmt.Printf("client.Lookup() should not throw an error.\n%s\n", err)
	}

	fmt.Println("Address of zzz.i2p:")
	// Addresses change all the time
	// fmt.Println(addr)

	// Output:
	//Address of zzz.i2p:
}
