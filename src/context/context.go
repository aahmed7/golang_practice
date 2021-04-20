package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// a context.Context is created for each request by net/http.
	// accessible using Context() method
	ctx := req.Context()
	fmt.Println("server: hello handler started.")
	defer fmt.Println("server: hello handler ended.")

	select {
	// Wait for a few seconds before sending a reply to the client.
	// This could simulate some work the server is doing. While working,
	// keep an eye on the context’s Done() channel for a signal that we
	// should cancel the work and return as soon as possible.
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Hello.\n")

	case <-ctx.Done():
		// The context’s Err() method returns an error that explains why
		// the Done() channel was closed.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
