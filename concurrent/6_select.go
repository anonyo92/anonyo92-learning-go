package concurrent

/* Go’s select lets you wait on multiple channel communication operations. 
	If one or more of the communications can proceed, 
		a single one that can proceed is chosen via a uniform pseudo-random selection.
	Otherwise, if there is a default case, that case is chosen. 
	If there is no default case, the "select" statement blocks until 
		at least one of the communications can proceed. */

import (
	"fmt"
	"time"
)

func RunChannelSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	/* Each channel will receive a value after some amount of time, to simulate
	 * e.g. blocking RPC operations executing in concurrent goroutines. */

	/* We’ll use select to await both of these values simultaneously,
	 * printing each one as it arrives. */
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	/* Timed Run this file to check total runtime.
	 * The total execution time will be only ~2 seconds since both
	 * the 1 and 2 second Sleeps execute concurrently */
}

/* Q. Why do we need the loop? It's running the select twice.
 *    What would've happened if we'd just done two channel receives
 *    like <-c1 and <-c2 back to back, instead of select-case?
 * A. In that case, both would have been sync. blocking calls */
