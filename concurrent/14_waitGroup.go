package concurrent

/* To wait for multiple goroutines to finish, we can use a wait group. */

import (
    "fmt"
    "sync"
    "time"
)

// This is the function we’ll run in every goroutine.
func workerWg(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second) // Sleep to simulate an expensive task.
    fmt.Printf("Worker %d done\n", id)
}

func RunWaitGroup() {
	/* This WaitGroup is used to wait for all the goroutines launched here to finish. 
	 * Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer. */
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Avoid re-use of the same i value in each goroutine closure.
		i := i

		/* Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done. 
		 * This way the worker itself does not have to be aware of the concurrency primitives 
		 * involved in its execution. Contrast with the workerPool example, where the workers
		 * had to be aware of the shared channels for concurrent communication. */
		go func() {
            defer wg.Done()
            workerWg(i)
        }()
	}

	/* Block until the WaitGroup counter goes back to 0. 
	 * It will happen when all the workers notify they’re done. */
	wg.Wait()

	/* Note that this approach has no straightforward way to propagate errors from workers. 
	 * For more advanced use cases, consider using the errgroup package. */
}