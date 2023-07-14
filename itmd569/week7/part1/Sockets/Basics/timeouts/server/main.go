// timeouts are important so that slow or disappaering clients dont hog resources (file des)

package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())

}
