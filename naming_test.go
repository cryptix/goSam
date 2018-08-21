package goSam

import (
	"fmt"
	"testing"
)

func TestClientLookupInvalid(t *testing.T) {
	var err error

	setup(t)
	defer teardown(t)

	addr, err := client.Lookup(`!(@#)`)
	if addr != "" || err == nil {
		t.Error("client.Lookup() should throw an error.")
	}

	repErr, ok := err.(ReplyError)
	if !ok {
		t.Fatalf("client.Lookup() should return a ReplyError")
	}
	if repErr.Result != ResultInvalidKey {
		t.Errorf("client.Lookup() should throw an ResultKeyNotFound error.\nGot:%+v\n", repErr)
	}
}

func ExampleClient_Lookup() {
	client, err := NewDefaultClient()
	if err != nil {
		fmt.Printf("NewDefaultClient() should not throw an error.\n%s\n", err)
		return
	}

	addr, err := client.Lookup("zzz.i2p")
	if err != nil {
		fmt.Printf("client.Lookup() should not throw an error.\n%s\n", err)
		return
	}

	fmt.Println("Address of zzz.i2p:")
	// Addresses change all the time
	fmt.Println(addr)

	// Output:
	//Address of zzz.i2p:
	//GKapJ8koUcBj~jmQzHsTYxDg2tpfWj0xjQTzd8BhfC9c3OS5fwPBNajgF-eOD6eCjFTqTlorlh7Hnd8kXj1qblUGXT-tDoR9~YV8dmXl51cJn9MVTRrEqRWSJVXbUUz9t5Po6Xa247Vr0sJn27R4KoKP8QVj1GuH6dB3b6wTPbOamC3dkO18vkQkfZWUdRMDXk0d8AdjB0E0864nOT~J9Fpnd2pQE5uoFT6P0DqtQR2jsFvf9ME61aqLvKPPWpkgdn4z6Zkm-NJOcDz2Nv8Si7hli94E9SghMYRsdjU-knObKvxiagn84FIwcOpepxuG~kFXdD5NfsH0v6Uri3usE3uSzpWS0EHmrlfoLr5uGGd9ZHwwCIcgfOATaPRMUEQxiK9q48PS0V3EXXO4-YLT0vIfk4xO~XqZpn8~PW1kFe2mQMHd7oO89yCk-3yizRG3UyFtI7-mO~eCI6-m1spYoigStgoupnC3G85gJkqEjMm49gUjbhfWKWI-6NwTj0ZnAAAA
}
