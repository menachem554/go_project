package main

import (
	"fmt"
)

func main() {
	// Long method (while loop)
	i := 1
	for i <= 10 {
	fmt.Println(i)
	i++
	}

	// Short method (standert for loop)
	for i := 1; i <= 10; i++ {
	fmt.Println("Number ", i)

	} 

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}