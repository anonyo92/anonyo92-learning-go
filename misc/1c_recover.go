package misc

/* Go makes it possible to recover from a panic, by using the recover built-in function. 
 * A recover can stop a panic from aborting the program and let it continue with execution instead. */

import "fmt"

// This function panics.
func mayPanic() {
    panic("a problem")
}

/* recover must be called within a deferred function. 
 * When the enclosing function panics, the defer will activate 
 * and a recover call within it will catch the panic. */
func RunRecover() {
	defer func() {
		/* The return value of recover is the error raised in the call to panic. */
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()
    mayPanic()

	/* This code will not run, because mayPanic panics. 
	 * The execution of main stops at the point of the panic and resumes in the deferred closure. */
    fmt.Println("After mayPanic()")
}