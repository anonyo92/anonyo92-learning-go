package webpkg

import (
	"fmt"
	"net/http"
)

/* A fundamental concept in net/http servers is handlers.
 * A handler is an object implementing the http.Handler interface.
 * A common way to write a handler is by using the http.HandlerFunc adapter
 * on functions with the appropriate signature. */

/* Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments.
 * The response writer is used to fill in the HTTP response. */
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received request: ", req)
	// Here our simple response is just “hello\n”.
	fmt.Fprintf(w, "hello\n")
}

/* This handler does something a little more sophisticated */
func headers(w http.ResponseWriter, req *http.Request) {
	// reading all the HTTP request headers and,
	for name, headers := range req.Header {
		for _, h := range headers {
			// echoing them into the response body.
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func RunHttpServer(port int) {
	/* We register our handlers on server routes using the http.HandleFunc convenience function.
	 * It sets up the default router in the net/http package and takes a function as an argument. */
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	/* Finally, we call the ListenAndServe with the port and a handler.
	 * nil tells it to use the default router we’ve just set up. */
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
