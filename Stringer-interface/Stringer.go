package main

import (
	"fmt"
)

// Person include ID, Name
type Person struct {
	ID   string
	Name string
}

func (p *Person) String() string {
	return fmt.Sprintf("Id: %s, Name: %s", p.ID, p.Name)
}

func main() {
	a := Person{"#1", "Jack"}
	b := Person{"#2", "Tony"}
	fmt.Println(a)
	fmt.Println(b)
}
