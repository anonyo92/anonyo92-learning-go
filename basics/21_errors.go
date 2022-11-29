package basics

import (
	"errors"
	"fmt"
)

/* In Go it’s idiomatic to communicate errors
 * via an explicit, separate return value */
func sampleFunc6(arg int) (int, error) {
	if arg == 42 {
		// errors.New() constructs a basic 'error' value with the given error message
		return -1, errors.New("can't accept arg 42")
	}
	// A nil value in the error position indicates that there was no error
	return arg + 3, nil
}

/* It’s possible to use custom types as errors by implementing the Error() method on them. */
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func sampleFunc7(arg int) (int, error) {
	if arg == 42 {
		// we use &argError syntax to build a new struct,
		// supplying values for the two fields arg and prob.
		return -1, &argError{arg, "can't work with 42"}
	}
	return arg + 3, nil
}

func runErrors() {
	for _, i := range []int{7, 42} {
		// an inline error check on the if line is a common idiom in Go code
		if r, e := sampleFunc6(i); e != nil {
			fmt.Println("sampleFunc6() failed", e)
		} else {
			fmt.Println("sampleFunc6() worked", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := sampleFunc7(i); e != nil {
			fmt.Println("sampleFunc7() failed", e)
		} else {
			fmt.Println("sampleFunc7() worked", r)
		}
	}

	/* If you want to programmatically use the data in a custom error,
	 * you’ll need to get the error as an instance of the custom error type
	 * via type assertion */
	_, e := sampleFunc7(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	} else {
		fmt.Println("Not an argError type")
	}
}
