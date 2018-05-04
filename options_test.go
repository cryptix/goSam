package goSam

import "testing"

func TestOptionAddr(t *testing.T) {
    ConnDebug = true
	client, err := NewClientFromOptions(SetAddr("127.0.0.1"))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionPort(t *testing.T) {
    ConnDebug = true
	client, err := NewClientFromOptions(SetPort("7656"))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

