package main

import (
	"io"
	"log"
	"net/http"
	"os"

	".."
)

func main() {
	goSam.ConnDebug = true

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
