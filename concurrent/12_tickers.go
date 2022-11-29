package concurrent

/* While timers are for when we want to do something once in the future,
 * tickers are for when we want to do something repeatedly at regular intervals. */

import (
    "fmt"
    "time"
)

func RunTickers() {
	/* Tickers use a similar mechanism to timers: a channel that is sent values. */
	ticker := time.NewTicker(2000 * time.Millisecond)
	done := make(chan bool)
	
	go func() {
        for { // Wait indefinitely till done signal is sent
            select {
            case <-done: // awaiting done signal
                return
            case t := <-ticker.C: // awaiting a 'tick' from the ticker
                fmt.Println("Tick at", t)
            }
        }
	}()
	
	time.Sleep(6600 * time.Millisecond)

	/* Tickers can be stopped like timers. 
	 * Once a ticker is stopped it won’t receive any more values on its channel. */
	ticker.Stop()
	time.Sleep(5 * time.Second) // waiting long enough to verify that ticker has indeed stopped.
    done <- true  // complete the other goroutine
    fmt.Println("Ticker stopped")
}