package concurrent

/* For managing shared-state more complex than counters, 
 * we can use a mutex to safely access data across multiple goroutines. */

import (
    "fmt"
    "sync"
)

/* Container holds a map of counters. Since we want to update it concurrently from multiple goroutines,
 * we add a Mutex to synchronize access. Note that mutexes must not be copied, 
 * so if this struct is passed around, it should be done by pointer. */
type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {
    c.mu.Lock() // Lock the mutex before accessing counters
    defer c.mu.Unlock() // Unlock it at the end of the function using a defer statement.
    c.counters[name]++
}

func RunMutex() {
	var wg sync.WaitGroup

	c := Container{
		// Note that the zero value of a mutex is usable as-is, so no initialization is required here.
        counters: map[string]int{"a": 0, "b": 0},
    }

	// This function increments a named counter in a loop for `n` times.
	doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
	}
	
	// Run several goroutines concurrently.
	wg.Add(3)
    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	// Note that they all access the same Container, and two of them access the same counter.

    wg.Wait() // Wait a for the goroutines to finish
    fmt.Println(c.counters)
}