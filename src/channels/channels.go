package main

import "fmt"

func main() {
	// Create a new channel with make(chan val-type).
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	// Here we send "ping" to the messages channel we made
	// above, from a new goroutine.
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
