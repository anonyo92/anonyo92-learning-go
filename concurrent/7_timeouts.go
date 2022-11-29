package concurrent

import (
	"fmt"
	"time"
)

func RunTimeouts() {
	/* Using buffered channel, so that the send in the goroutine is
	 * nonblocking. This is a common pattern to prevent goroutine leaks
	 * in case the channel is never read. */
	start := time.Now()
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1: // this case awaits the result from c1
		fmt.Println(res)
	case <-time.After(1 * time.Second): // awaits a value to be sent after the timeout of 1s
		fmt.Println("timeout after 1s. Elapsed since start:", time.Since(start))
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	/* If we allow a longer timeout of 3s, then the receive from c2
	 * will succeed and we’ll print the result */
	select {
	case res := <-c2:
		fmt.Printf("Received result %v. Elapsed since start: %v\n", res, time.Since(start))
	case <-time.After(3 * time.Second):
		fmt.Println("timeout after 3s. Elapsed since start:", time.Since(start))
	}

	fmt.Printf("Ending now. Elapsed since start: %v\n", time.Since(start))
}
