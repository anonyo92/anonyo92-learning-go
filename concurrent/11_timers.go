package concurrent

/* Timers represent a single event in the future. We tell the timer how long you want to wait, 
 * and it provides a channel that will be notified at that time. */

import (
    "fmt"
    "time"
)

func RunTimers() {
	start := time.Now()

	// This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
	val := <-timer1.C // Sends the current time (with elapsed seconds as the channel value)
	fmt.Printf("Timer 1 fired. Sent %v in channel. Elapsed since start:%v\n", val, time.Since(start))
	

	/* If we just wanted to wait, we could have used time.Sleep(). 
	 * One reason a timer may be useful is that we can cancel the timer before it fires.
	 * Here’s an example of that. */
	timer2 := time.NewTimer(time.Second)
    go func() {
        val := <-timer2.C
        fmt.Printf("Timer 2 fired. Sent %v in channel. Elapsed since start:%v\n", val, time.Since(start))
	}()
	
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped. Elapsed since start: ", time.Since(start))
	}

	// Giving the timer2 enough time to fire, if it ever was going to, to show that it is in fact stopped.
	time.Sleep(2 * time.Second)
	
	/* The first timer will fire ~2s after we start the program, 
	 * but the second should be stopped before it has a chance to fire. */
}