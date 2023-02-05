package main

import "fmt"

func main() {
	var x int = 1

	y := "hello" // implicit assignment, type inferred from value

	// y := 3 // error: no new variables on left side of :=
	// we can't change the type of a variable once it's been declared

	fmt.Println(x, y)

	var z int // zero value of int is 0
	// the above syntax is used to declare a variable without assigning a value

	fmt.Println(z)

	t := 1
	fmt.Printf("%T\n", t) // %T prints the type of the variable

	tt := uint(t) // type conversion
	fmt.Printf("%T\n", tt)
}
