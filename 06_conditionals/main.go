package main

import "fmt"

func main() {
	x := 10
	y := 10
	color := "red"

	// if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)

	} else if x == y {
		fmt.Printf("%d is equals %d\n", y, x)

	} else {
		fmt.Printf("%d is less than %d\n", y, x)

	}

	// Switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT blue or red")
	}
}
