package misc

/* Use os.Exit to immediately exit with a given status. */

import (
    "fmt"
    "os"
)

func RunExit() {
	// defers will not be run when using os.Exit, so this fmt.Println will never be called.
	defer fmt.Println("!")

	// Exit with status 3.
	os.Exit(3)
	
	/* Unlike e.g. C, Go does not use an integer return value from main to indicate exit status. 
	 * If we need to exit with a non-zero status we should use os.Exit(). */
}