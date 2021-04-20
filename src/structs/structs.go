package main

import "fmt"

// create a person type struct
type person struct {
	name string
	age  int
}

// newPerson constructs a new person
func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// Create a new person
	fmt.Println(person{"bob", 20})

	// Create a new person with struct fields specified
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero
	fmt.Println(person{name: "bob"})

	// Pointer to struct
	fmt.Println(&person{name: "bob"})

	// constructor style definition
	fmt.Println(newPerson("John"))

	// Access fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// pointer to struct
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

}
