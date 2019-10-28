// This sample code implement a simple web service.
package main

import (
	"log"
	"net/http"

	"01Go_In_Action/ch9Test_and_Performance/i4api/handlers"
)

// main is the entry point for the application.
func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
