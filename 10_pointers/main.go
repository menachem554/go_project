package main

import "fmt"

func main() {
	a := 5
	b := &a	
	
	fmt.Println(a, b) // prints 5 and memeroy address 0xc0000b4008 

	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // int*

	// read the value from memory address
	fmt.Println(*b) // prints 5
	fmt.Println(*&a) // prints the same

	// change value with the pointer
	*b = 10
	fmt.Println(a)

	// everything in go is passed by value
}