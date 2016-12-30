// This app shows the most basic form of web server app

package main

import (
	"fmt"
	"log"
	"net/http"
)

//defaultHandler handles any requests on root and simply outputs "Go is awesome!" as output
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is awesome!")
}

func main() {

	//add handler for requests on root
	//NOTE: If no handler is set then listener will return HTTP 404 when requested
	http.HandleFunc("/", defaultHandler)

	log.Println("Waiting for requests on port 4321")

	//This will block and listen for incoming requests
	// NOTE: localhost:4321 will onlt accept traffic from localhost, :4321 is open to any address
	log.Fatal(http.ListenAndServe(":4321", nil))
}
