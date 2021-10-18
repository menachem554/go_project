package main

import "fmt"

var isCool = true


func main() {
	name, email, age := "menachem", "mr@gmail.com", 25
	fmt.Println(name, age, email)
	fmt.Printf("%T\n", isCool)
}