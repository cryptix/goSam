package goSam

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"net"

	"github.com/cryptix/go/debug"
)

// A Client represents a single Connection to the SAM bridge
type Client struct {
	host string
	port string

	SamConn net.Conn
	rd      *bufio.Reader
	//id      int32

	sigType     string
	destination string

	inLength   uint
	inVariance int
	inQuantity uint
	inBackups  uint

	outLength   uint
	outVariance int
	outQuantity uint
	outBackups  uint

	dontPublishLease bool
	encryptLease     bool

	reduceIdle         bool
	reduceIdleTime     uint
	reduceIdleQuantity uint

	closeIdle     bool
	closeIdleTime uint

	compression bool

	debug bool
	//NEVER, EVER modify lastaddr yourself.
	lastaddr string
	id       int32
}

var SAMsigTypes = []string{
	"SIGNATURE_TYPE=DSA_SHA1",
	"SIGNATURE_TYPE=ECDSA_SHA256_P256",
	"SIGNATURE_TYPE=ECDSA_SHA384_P384",
	"SIGNATURE_TYPE=ECDSA_SHA512_P521",
	"SIGNATURE_TYPE=EdDSA_SHA512_Ed25519",
}

// NewDefaultClient creates a new client, connecting to the default host:port at localhost:7656
func NewDefaultClient() (*Client, error) {
	return NewClient("localhost:7656")
}

// NewClient creates a new client, connecting to a specified port
func NewClient(addr string) (*Client, error) {
	return NewClientFromOptions(SetAddr(addr))
}

func (c *Client) NewID() int32 {
	return rand.Int31n(math.MaxInt32)
}

// NewClientFromOptions creates a new client, connecting to a specified port
func NewClientFromOptions(opts ...func(*Client) error) (*Client, error) {
	var c Client
	c.host = "127.0.0.1"
	c.port = "7656"
	c.inLength = 3
	c.inVariance = 0
	c.inQuantity = 1
	c.inBackups = 1
	c.outLength = 3
	c.outVariance = 0
	c.outQuantity = 1
	c.outBackups = 1
	c.dontPublishLease = true
	c.encryptLease = false
	c.reduceIdle = false
	c.reduceIdleTime = 300000
	c.reduceIdleQuantity = 1
	c.closeIdle = true
	c.closeIdleTime = 600000
	c.debug = false
	c.sigType = SAMsigTypes[4]
	c.id = 0
	for _, o := range opts {
		if err := o(&c); err != nil {
			return nil, err
		}
	}
	conn, err := net.Dial("tcp", c.samaddr())
	if err != nil {
		return nil, err
	}
	if c.debug {
		conn = debug.WrapConn(conn)
	}
	c.SamConn = conn
	c.rd = bufio.NewReader(conn)
	return &c, c.hello()
}

//return the combined host:port of the SAM bridge
func (c *Client) samaddr() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}

// send the initial handshake command and check that the reply is ok
func (c *Client) hello() error {
	r, err := c.sendCmd("HELLO VERSION MIN=3.0 MAX=3.1\n")
	if err != nil {
		return err
	}

	if r.Topic != "HELLO" {
		return fmt.Errorf("Unknown Reply: %+v\n", r)
	}

	if r.Pairs["RESULT"] != "OK" {
		return fmt.Errorf("Handshake did not succeed\nReply:%+v\n", r)
	}

	return nil
}

// helper to send one command and parse the reply by sam
func (c *Client) sendCmd(str string, args ...interface{}) (*Reply, error) {
	if _, err := fmt.Fprintf(c.SamConn, str, args...); err != nil {
		return nil, err
	}

	line, err := c.rd.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return parseReply(line)
}

// Close the underlying socket to SAM
func (c *Client) Close() error {
	c.rd = nil
	return c.SamConn.Close()
}

func (c *Client) NewClient() (*Client, error) {
	return NewClientFromOptions(
		SetHost(c.host),
		SetPort(c.port),
		SetDebug(c.debug),
		SetInLength(c.inLength),
		SetOutLength(c.outLength),
		SetInVariance(c.inVariance),
		SetOutVariance(c.outVariance),
		SetInQuantity(c.inQuantity),
		SetOutQuantity(c.outQuantity),
		SetInBackups(c.inBackups),
		SetOutBackups(c.outBackups),
		SetUnpublished(c.dontPublishLease),
		SetEncrypt(c.encryptLease),
		SetReduceIdle(c.reduceIdle),
		SetReduceIdleTime(c.reduceIdleTime),
		SetReduceIdleQuantity(c.reduceIdleQuantity),
		SetCloseIdle(c.closeIdle),
		SetCloseIdleTime(c.closeIdleTime),
		SetCompression(c.compression),
		setlastaddr(c.lastaddr),
		setid(c.id),
	)
}
