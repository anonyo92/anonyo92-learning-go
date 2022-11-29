package concurrent

import "fmt"

/* When using channels as function parameters, we can
 * specify if a channel is meant to only send or receive values.
 * This specificity increases the type-safety of the program. */

/* This function accepts a channel for only sending values.
 * It would be a compile-time error to try to receive on this channel. */
func ping(sendingChan chan<- string, msg string) {
	sendingChan <- msg
}

/* This function accepts one channel for receives and a second for sends */
func pong(rcvingChan <-chan string, sendingChan chan<- string) {
	msg := <-rcvingChan
	sendingChan <- "received this:" + msg
}

func RunChannelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "pinging message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
