package concurrent

/* We can also use the range syntax to iterate
 * over values received from a channel */

import "fmt"

func RunChannelIteration() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	/* This range iterates over each element as it’s
	 * received from channel. Because we closed the channel above,
	 * the iteration terminates after receiving the 2 elements. 
	 * Had we not closed the channel, at the last iteration there'd be a deadlock,
	 	as this iteration would try to receive more messages from the channel and wait forever. */
	for elem := range queue {
		fmt.Println(elem)
	}
	/* This example also showed that it’s possible to close a
	 * non-empty channel but still have the remaining values be received. */
}
