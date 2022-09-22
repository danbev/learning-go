package main

import (
	"fmt"
)

type something struct {
	name string
	x int
}

func (s something) print() {
	fmt.Printf("s.name: %s, s.x: %d\n", s.name, s.x)
}

func main() {
	fmt.Printf("Struct example\n");
	var _ something 
	s := something {}
	s.print();

	s2 := something { "Fletch", 47}
	fmt.Printf("s2.name: %s, s2.x: %d\n", s2.name, s2.x)

	anon := struct {
		name string
		x int
	} {
		name: "Fletch",
		x: 47,
	}
	fmt.Printf("anon.name: %s, anon.x: %d\n", anon.name, anon.x)
}
