package e

import (
	"fmt"
)

// NewPerson returns Person object.
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// Person is sumple struct.
type Person struct {
	Name string
	Age  int
}

// String returns string.
func (p *Person) String() string {
	return fmt.Sprintf("name %s, age %d", p.Name, p.Age)
}
