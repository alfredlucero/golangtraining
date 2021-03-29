package main

import "fmt"

// Structs are typed collections of fields and useful for grouping data together to form records
type person struct {
	name string
	age  int
}

// Safely return pointer to local variable as a local variable will survive the scope of the function
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	// Can also name the fields when initializing a struct
	fmt.Println(person{name: "Alfred", age: 26})

	// Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Regine", age: 24})

	// Idiomatic to encapsulate a new struct creation in constructor functions
	fmt.Println(newPerson("Juno"))

	// Access struct fields with a dot
	s := person{name: "Ally", age: 20}
	fmt.Println(s.name)

	// Can use dots with struct pointers; pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 25
	fmt.Println(sp.age)

}
