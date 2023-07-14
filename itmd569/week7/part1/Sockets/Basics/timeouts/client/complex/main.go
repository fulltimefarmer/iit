package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second, // timeout when establishing a new TCP connection
				KeepAlive: 30 * time.Second, // timeout for how long to probe for the health of the connection
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second, // timeout when doing tls handshake
			ResponseHeaderTimeout: 10 * time.Second, // timeout when reading headers of the response
			ExpectContinueTimeout: 1 * time.Second,  // timeout between when client sends request headers and recieivng the signal to send the body of the request
		},
	}

	_, err := c.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {
		log.Fatal()
	}

}
