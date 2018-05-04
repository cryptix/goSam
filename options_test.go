package goSam

import "testing"

func TestOptionAddr(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("127.0.0.1"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionPort(t *testing.T) {
	client, err := NewClientFromOptions(SetPort("7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionDebug(t *testing.T) {
	client, err := NewClientFromOptions(SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}
