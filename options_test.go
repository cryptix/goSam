package goSam

import "testing"

func optaddr(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func optport(t *testing.T) {
	client, err := NewClientFromOptions(SetPort("7656"))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionsHello(t *testing.T) {
	ConnDebug = true
	optaddr(t)
	optport(t)
}
