package goSam

import "testing"

func TestOptionAddrString(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("127.0.0.1:7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionAddrStringLh(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("localhost:7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionAddrSlice(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("127.0.0.1", "7656"), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionAddrMixedSlice(t *testing.T) {
	client, err := NewClientFromOptions(SetAddr("127.0.0.1", 7656), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionHost(t *testing.T) {
	client, err := NewClientFromOptions(SetHost("127.0.0.1"), SetDebug(true))
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

func TestOptionPortInt(t *testing.T) {
	client, err := NewClientFromOptions(SetPort(7656), SetDebug(true))
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

func TestOptionInLength(t *testing.T) {
	client, err := NewClientFromOptions(SetInLength(3), SetDebug(true))
	client.inlength()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionOutLength(t *testing.T) {
	client, err := NewClientFromOptions(SetInLength(3), SetDebug(true))
	client.outlength()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}
