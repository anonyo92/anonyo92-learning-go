package concurrent

/* A goroutine is a lightweight thread of execution */

import (
	"fmt"
	"time"
)

func workerFunc1(caller string) {
	for i := 0; i < 3; i++ {
		fmt.Println(caller, ":", i)
	}
}
func RunGoroutines() {
	workerFunc1("direct")

	/* To invoke a function in a goroutine, use `go f(s)`
	 * This new goroutine will execute concurrently with the calling one */
	go workerFunc1("goroutine") // async non-blocking call

	/* We can also start a goroutine for an anonymous function call */
	go func(msg string) {
		fmt.Println(msg)
	}("going") // async non-blocking call

	time.Sleep(time.Second)
	fmt.Println("main done")
}
