package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array example");
	var x [3]int  // this is just the declaration so now assignement operator ('=')
	fmt.Printf("x[0]: %d, x[1]: %d, x[2]: %d\n", x[0], x[1], x[2])
	var y = [3]int{1, 2, 3}
	fmt.Printf("y[0]: %d, y[1]: %d, y[2]: %d\n", y[0], y[1], y[2])
}
