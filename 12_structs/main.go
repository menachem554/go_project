package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// alternative syntax
	// firstName, lastName, city, gender string
	// age int
}

// Greeting method (value receiver)
func (person Person) greeet() string {
	return "Hello my name is " + person.firstName + " " + person.lastName +
		" and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spousedLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spousedLastName
	}
}

func main() {
	// init person
	person1 := Person{firstName: "Samantha", lastName: "Smith",
		city: "Boston", gender: "f", age: 25}
	fmt.Println(person1)

	// alternative syntax
	person2 := Person{"Bob", "Johnson", "NY", "m", 30}
	fmt.Println(person2)

	// Modify parts of the object
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	fmt.Println(person1.greeet())

	person1.getMarried("Williams")
	fmt.Println(person1.greeet())

	person2.getMarried("Thompson")
	fmt.Println(person2.greeet())
}

// we have two kind of methods
// value receivers and pointer receivers
