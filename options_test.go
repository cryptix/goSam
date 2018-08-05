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
	client, err := NewClientFromOptions(SetAddrMixed("127.0.0.1", 7656), SetDebug(true))
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
	client, err := NewClientFromOptions(SetPortInt(7656), SetDebug(true))
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

func TestOptionInVariance(t *testing.T) {
	client, err := NewClientFromOptions(SetInVariance(1), SetDebug(true))
	client.invariance()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionOutVariance(t *testing.T) {
	client, err := NewClientFromOptions(SetOutVariance(1), SetDebug(true))
	client.outvariance()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionInQuantity(t *testing.T) {
	client, err := NewClientFromOptions(SetInQuantity(6), SetDebug(true))
	client.inquantity()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionOutQuantity(t *testing.T) {
	client, err := NewClientFromOptions(SetOutQuantity(6), SetDebug(true))
	client.outquantity()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionInBackups(t *testing.T) {
	client, err := NewClientFromOptions(SetInBackups(5), SetDebug(true))
	client.inbackups()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionOutBackups(t *testing.T) {
	client, err := NewClientFromOptions(SetOutBackups(5), SetDebug(true))
	client.outbackups()
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionEncryptLease(t *testing.T) {
	client, err := NewClientFromOptions(SetEncrypt(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionUnpublishedLease(t *testing.T) {
	client, err := NewClientFromOptions(SetUnpublished(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionReduceIdle(t *testing.T) {
	client, err := NewClientFromOptions(SetReduceIdle(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionReduceIdleTime(t *testing.T) {
	client, err := NewClientFromOptions(SetReduceIdleTime(300001), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionReduceIdleCount(t *testing.T) {
	client, err := NewClientFromOptions(SetReduceIdleQuantity(4), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionCloseIdle(t *testing.T) {
	client, err := NewClientFromOptions(SetCloseIdle(true), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestOptionCloseIdleTime(t *testing.T) {
	client, err := NewClientFromOptions(SetCloseIdleTime(300001), SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}
