package basics

import "fmt"

type employee struct {
	name string
	age int
}

func createEmployee1(name string, age int) employee {
	return employee{name: name, age: age}
}

func createEmployee2(name string, age int) *employee {
	return &employee{name: name, age: age}
}

func RunTypeAssertions() {
	var a interface{} = "initial"

	// Printing Types
	fmt.Printf("%v is of type %T\n", a, a)

	// Asserting simple types
	if r, ok := a.(string); ok {
		fmt.Println(r, "is a string")
	} else {
		fmt.Println(a, "is not a string")
	}

	// Asserting struct types
	var emp1 interface{} = createEmployee1("John Doe", 42)
	fmt.Printf("%v is of type %T\n", emp1, emp1)
	if r, ok := emp1.(employee); ok {
		fmt.Println(r, "is an employee struct")
	} else {
		fmt.Println(r, "is not an employee struct")
	}

	var emp2 interface{} = createEmployee2("Martha Williams", 59)
	fmt.Printf("%v is of type %T\n", emp2, emp2)
	if r, ok := emp2.(*employee); ok {
		fmt.Println(r, "is an employee struct pointer")
	} else {
		fmt.Println(r, "is not an employee struct pointer")
	}
}

func RunTypeSwitch(i interface{}) {
	switch t := i.(type) {
	case string :
		fmt.Println(i, "is a string")
	case bool :
		fmt.Println(i, "is a boolean")
	case int :
		fmt.Println(i, "is an integer")
	case float64 :
		fmt.Println(i, "is a float")
	default:
		fmt.Printf("%v is of type %v\n", i, t)
	}
}