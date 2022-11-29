package webpkg

/* The Go standard library comes with excellent support for HTTP clients and servers in the net/http package.
In this example we’ll use it to issue simple HTTP requests. */

import (
	"bufio"
	"fmt"
	"net/http"
)

// Issue an HTTP GET request to a server
func runHttpGet(url string) {
	/* http.Get is a convenient shortcut around creating an http.Client object and calling its Get method;
	 * it uses the http.DefaultClient object which has useful default settings. */
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response status code:", resp.StatusCode)

	// Print the HTTP response headers.
	fmt.Println("\nResponse Headers:")
	for k, v := range resp.Header {
		fmt.Printf("%v : %v\n", k, v)
	}

	// Print the first 5 lines of the response body.
	fmt.Println("\nResponse Body (first 5 lines):")
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func RunHttpClient(method string, url string, payload string) {
	if url == "" {
		url = "http://gobyexample.com"
	}
	switch method {
	case "GET", "Get", "get":
		runHttpGet(url)
	// case "POST", "Post", "post": runHttpPost(url, payload)
	// case "PUT", "Put", "put": runHttpPut(url, payload)
	// case "DELETE", "Delete", "delete": runHttpDelete(url)
	default:
		fmt.Println("Unsupported HTTP method:", method)
	}
}
