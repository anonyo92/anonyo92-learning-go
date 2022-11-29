package webpkg

/* A Context carries deadlines, cancellation signals,
 * and other request-scoped values across API boundaries and goroutines. */

import (
	"fmt"
	"net/http"
	"time"
)

/* HTTP servers are useful for demonstrating the usage of context.Context for controlling cancellation. */

func helloCtx(w http.ResponseWriter, req *http.Request) {
	/* A context.Context is created for each request by the net/http machinery,
	 * and is available with the Context() method of the request. */
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	// Wait for a few seconds before sending a reply to the client, simulating some work the server is doing.
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		/* While working, keep an eye on the context’s Done() channel for a signal
		 * that we should cancel the work and return as soon as possible. */

		// The context’s Err() method returns an error that explains why the Done() channel was closed.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func RunContexts(port int) {
	// Register our handler on the “/hello” route, and start serving.
	http.HandleFunc("/hello", helloCtx)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
