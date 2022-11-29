package concurrent

/* We can use channels to synchronize execution across goroutines.
 * Here’s an example of using a blocking receive to wait for
 * another goroutine to finish. When waiting for multiple goroutines
 * to finish, you may prefer to use a "WaitGroup". */

import (
	"fmt"
	"time"
)

/* This is the function we’ll run in a goroutine.
 * The done channel will be used to notify another goroutine that
 * this function’s work is done. */
func workerFunc2(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done.")

	// Send a value to notify that we’re done.
	done <- true
}

func RunChannelSync() {
	notifChannel := make(chan bool)

	// Start a worker goroutine, giving it the channel to notify on
	go workerFunc2(notifChannel)

	// Block until we receive a notification from the worker on the channel.
	<-notifChannel

	/* If we removed the <- notifChannel line from this program,
	 * the program would exit before the worker even started.
	 * This pattern of synchronization is like the wait-notify in Java. */
}
