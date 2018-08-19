package e_test

import (
	"fmt"

	"github.com/budougumi0617/go-testing/e"
)

/**
 * $ godoc -http=:6060
 * http://localhost:6060/pkg/github.com/budougumi0617/go-testing/e/
 *
 * Examples
 *   Package
 *   NewPerson
 *   Person
 *   Person.String
 *   Person.String (Bob)
 */

func Example() {
	ps := map[int]*e.Person{
		0: &e.Person{Name: "Alice", Age: 12},
		1: &e.Person{Name: "Bob", Age: 10},
		2: &e.Person{Name: "Chris", Age: 15},
	}

	for _, p := range ps {
		fmt.Println(p)
	}

	// Unordered output:
	// name Alice, age 12
	// name Bob, age 10
	// name Chris, age 15
}

func ExampleNewPerson() {
	p := e.NewPerson("Bob", 10)
	fmt.Println(p.Name)
	fmt.Println(p.Age)

	// Output:
	// Bob
	// 10
}

func ExamplePerson() {
	p := &e.Person{Name: "Alice", Age: 12}
	fmt.Println(p.Name)
	fmt.Println(p.Age)

	// Output:
	// Alice
	// 12
}

func ExamplePerson_String() {
	p := &e.Person{Name: "Alice", Age: 12}
	fmt.Println(p)

	// Output:
	// name Alice, age 12
}

func ExamplePerson_String_bob() {
	p := &e.Person{Name: "Bob", Age: 10}
	fmt.Println(p)

	// Output:
	// name Bob, age 10
}
