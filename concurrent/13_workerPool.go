package concurrent

import (
    "fmt"
    "time"
)

/* Here’s the worker, of which we’ll run several concurrent instances. 
 * These workers will receive work on the jobs channel and send the corresponding results on results. 
 * We’ll sleep a second per job to simulate an expensive task. */
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 100
    }
}

func RunWorkerPool() {
	/* In order to use our pool of workers we need to send them work and collect their results. 
	 * We make 2 channels for this. */
	const numWorkers = 3
	const numJobs = 5
    jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// This starts up 3 workers, all initially blocked because there are no jobs yet.
	for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results)
	}
	
	// Here we send 5 jobs and then close that channel to indicate that’s all the work we have.
	for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
	close(jobs) // done sending jobs
	
	/* Finally we collect all the results of the work. 
	 * This also ensures that the worker goroutines have finished. 
	 * An alternative way to wait for multiple goroutines is to use a WaitGroup. */
	for a := 1; a <= numJobs; a++ {
        fmt.Println("Received result:", <-results)
	}
	
	/* The program only takes about 2 seconds despite doing about 5 seconds of total work 
	 * because there are 3 workers operating concurrently. */
}