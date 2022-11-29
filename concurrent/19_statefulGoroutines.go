package concurrent

/* This program illustrates implementing this same state management task as atomic-counters and mutexes
 * using only goroutines and channels. This channel-based approach aligns with Go’s ideas of sharing memory by 
 * communicating and having each piece of data owned by exactly 1 goroutine. */

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

/* In this example our state will be owned by a single goroutine.
 * This will guarantee that the data is never corrupted with concurrent access. 
 * In order to read or write that state, other goroutines will send messages 
 * 	to the owning goroutine and receive corresponding replies. 
 * These readOp and writeOp structs encapsulate those requests and a way for the owning goroutine to respond. */
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func RunStatefulGoroutine() {
	// Counters to count how many operations we perform.
	var readOps uint64
	var writeOps uint64
	
	/* The reads and writes channels will be used by other goroutines 
	 * 	to issue read and write requests, respectively. */
	reads := make(chan readOp)
	writes := make(chan writeOp)
	
	go func() {
		// This map represents a state that is owned privately by this stateful goroutine.
		var state = make(map[int]int)

		/* This goroutine repeatedly selects on the reads and writes channels, 
		 * responding to requests as they arrive. A response is executed by 
		 * first performing the requested operation (read or write),
		 * and then sending a value on the response channel resp to indicate success
		 * (and the desired value in the case of reads). */
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key] // read local state, and respond to channel with read value
            case write := <-writes:
                state[write.key] = write.val // write to local state, and respond with success
                write.resp <- true
            }
        }
	}()


	/* This starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel. */
	for r := 0; r < 100; r++ {
        go func() {
            for {
				// Each read requires constructing a readOp,
                read := readOp{
                    key:  rand.Intn(5),
					resp: make(chan int),
				}
                reads <- read // sending it over the reads channel
				<-read.resp // and receiving the result over the provided resp channel (blocking op).
				
                atomic.AddUint64(&readOps, 1) // Update read-op counter atomically - shared by all readers
                time.Sleep(time.Millisecond) // Stagger and also simulate read-delay
            }
        }()
	}

	// We start 10 writes as well, using a similar approach as reads.
	for w := 0; w < 10; w++ {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
					resp: make(chan bool),
				}
                writes <- write
				<-write.resp
				
                atomic.AddUint64(&writeOps, 1) // Update write-op counter atomically - shared by all writers
                time.Sleep(time.Millisecond) // Stagger and also simulate read-delay
            }
        }()
	}
	
	// Let the goroutines work for a while.
    time.Sleep(time.Second)

	// Finally, capture and report the op counts.
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}