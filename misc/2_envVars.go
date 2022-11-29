package misc

import (
    "fmt"
    "os"
    "strings"
)

func RunEnvVars() {
	os.Setenv("FOO", "1") // Sets an environment variable "FOO" to "1"
    fmt.Println("FOO =", os.Getenv("FOO")) // Gets the value of an environment variable "FOO" 
	fmt.Println("BAR =", os.Getenv("BAR")) // Returns an empty string if the key isn’t present in the environment.
	
	fmt.Println()
	/* os.Environ() lists all key/value pairs in the environment. 
	 * This returns a slice of strings in the form KEY=value. */
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0], "=", pair[1])
    }
}