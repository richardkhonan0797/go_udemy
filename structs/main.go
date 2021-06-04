package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	name := person{
		firstName: "First",
		lastName:  "Last",
		contactInfo: contactInfo{
			email:   "struct@email.co",
			zipCode: 12345,
		},
	}
	fmt.Println(name)
	name.firstName = "first"
	name.lastName = "last"
	fmt.Printf("%+v\n", name)
	name.print()
	namePointer := &name
	// shortcut
	// name.updateName("golang") => go will handle the pointer to person
	fmt.Printf("%p\n", namePointer)
	namePointer.updateName("golang")
	name.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}
