package basics

import "fmt"

/* Structs are typed collections of fields */
type person struct {
	name string
	age  int
}

/* This function returns a struct-pointer */
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func runStructs() {
	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// But naming the fields is not necessary (auto-ordered)
	fmt.Println(person{"Bob", 20})

	// Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// Access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - they are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
