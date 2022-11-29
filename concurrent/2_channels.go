package concurrent

/* Channels are the pipes that connect concurrent goroutines.
 * You can send values into channels from one goroutine and
 * receive those values into another goroutine. */

import "fmt"

func RunChannels() {
	/* Create a new channel with make(chan val-type).
	 * Channels are typed by the values they convey. */
	msgChannel := make(chan string)

	/* Send a value into a channel using the `channel <- value` syntax. */
	go func() { 
		fmt.Println("Sending a message from a concurrent go-routine...")
		msgChannel <- "ping" 
	}()

	/* The `<-channel` syntax receives a value from the channel. */
	msg := <-msgChannel

	/* By default sends and receives block until both the sender and receiver are ready.
	 * This property allowed us to wait at the end of our program for the "ping" message
	 * without having to use any other synchronization. */
	fmt.Println("Received this message in main from a concurrent go-routine:", msg)
}
