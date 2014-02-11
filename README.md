goSam
=====

A go library for using the [I2P](https://geti2p.net/en/) Simple Anonymous Messaging ([SAM version 3.0](https://geti2p.net/en/docs/api/samv3)) bridge

This is in an **early development stage**. I would love to hear about any issues or ideas for improvement.

## Installation
```
go get github.com/cryptix/goSam
```

## Using it for HTTP Transport

I implemented `Client.Dial` like `net.Dial` so you can use go's library packages like http.

```go
package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/cryptix/goSam"
)

func main() {
	// create a default sam client
	sam, err := goSam.NewDefaultClient()
	checkErr(err)

	log.Println("Client Created")

	// create a transport that uses SAM to dial TCP Connections
	tr := &http.Transport{
		Dial: sam.Dial,
	}

	// create  a client using this transport
	client := &http.Client{Transport: tr}

	// send a get request
	resp, err := client.Get("http://stats.i2p/")
	checkErr(err)
	defer resp.Body.Close()

	log.Printf("Get returned %+v\n", resp)

	// create a file for the response
	file, err := os.Create("stats.html")
	checkErr(err)
	defer file.Close()

	// copy the response to the file
	_, err = io.Copy(file, resp.Body)
	checkErr(err)

	log.Println("Done.")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
```

### TODO

* Implement `STREAM ACCEPT` and `STREAM FORWARD`
* Implement datagrams (Repliable and Anon)