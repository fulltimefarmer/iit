package main

import (
	"crypto/rand"
	"crypto/tls"
	"log"
)

func main() {
	// load server pem and key
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	// create a config to store cert
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader
	service := "0.0.0.0:8000"

	// start a TCP server secured with the server certificate that listens on 0.0.0.0:8000
	_, err = tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
}
